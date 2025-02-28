package limiter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/whlxbd/gomall/common/mtl"
)

type cpuMonitor struct {
	client   v1.API
	hostname string
}

type methodLimiter struct {
	currentLimit    int32   // 最大令牌数/容量
	availableTokens int32   // 当前可用令牌数
	lastRefillTime  int64   // 上次填充令牌的时间戳
	refillRate      float64 // 每秒填充令牌数
	qpsLimitGauge   prometheus.Gauge
	qpsCurrentGauge prometheus.Gauge
}

type DynamicMethodQPSLimiter struct {
	limiters         sync.Map
	defaultLimit     int32
	cpuMonitor       *cpuMonitor
	highCPUThreshold float64 // 高负载阈值
	lowCPUThreshold  float64 // 低负载阈值
}

func newCPUMonitor(promAddr string) (*cpuMonitor, error) {
	client, err := api.NewClient(api.Config{
		Address: promAddr,
	})
	if err != nil {
		return nil, err
	}

	v1api := v1.NewAPI(client)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	return &cpuMonitor{
		client:   v1api,
		hostname: hostname,
	}, nil
}

func (c *cpuMonitor) getCPUUsage(ctx context.Context) (float64, error) {
	query := `rate(process_cpu_seconds_total[30s]) * 100`

	result, _, err := c.client.Query(ctx, query, time.Now())
	if err != nil {
		return 0, err
	}

	if result.Type() != model.ValVector {
		return 0, errors.New("unexpected result type")
	}

	vector := result.(model.Vector)
	if len(vector) == 0 {
		return 0, errors.New("no CPU metrics found")
	}

	return float64(vector[0].Value), nil
}

func NewDynamicMethodQPSLimiter(defaultLimit int32) *DynamicMethodQPSLimiter {
	promAddr := os.Getenv("METRICS_PORT")
	if promAddr == "" {
		klog.Errorf("METRICS_PORT is not set")
		panic("METRICS_PORT must be set to use the limiter")
	}

	if !strings.HasPrefix(promAddr, "http://") && !strings.HasPrefix(promAddr, "https://") {
		promAddr = "http://" + promAddr
	}

	cpuMonitor, err := newCPUMonitor(promAddr)
	if err != nil {
		klog.Errorf("Failed to create CPU monitor: %v", err)
		panic("Failed to create CPU monitor" + err.Error())
	}

	limiter := &DynamicMethodQPSLimiter{
		defaultLimit:     defaultLimit,
		cpuMonitor:       cpuMonitor,
		highCPUThreshold: 80, // 默认80
		lowCPUThreshold:  20, // 默认20
	}

	go limiter.monitor()
	return limiter
}

func (l *DynamicMethodQPSLimiter) initMethodLimiter(method string) {
	fmt.Printf("Initializing method limiter for %s\n", method)

	qpsLimitGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "kitex_method_qps_limit",
		Help: "Current QPS limit per method",
		ConstLabels: prometheus.Labels{
			"method": method,
		},
	})

	qpcCurrentGauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "kitex_method_qps_current",
		Help: "Current QPS per method",
		ConstLabels: prometheus.Labels{
			"method": method,
		},
	})

	mtl.Registry.MustRegister(qpsLimitGauge)
	mtl.Registry.MustRegister(qpcCurrentGauge)

	ml := &methodLimiter{
		currentLimit:    l.defaultLimit,
		availableTokens: l.defaultLimit, // 初始填满令牌
		lastRefillTime:  time.Now().UnixNano(),
		refillRate:      float64(l.defaultLimit), // 每秒填充满桶
		qpsLimitGauge:   qpsLimitGauge,
		qpsCurrentGauge: qpcCurrentGauge,
	}

	l.limiters.Store(method, ml)
}

func (l *DynamicMethodQPSLimiter) GetMethodLimiter(method string) *methodLimiter {
	// 获取方法对应限流器，若不存在则创建新的
	if ml, ok := l.limiters.Load(method); ok {
		return ml.(*methodLimiter)
	}

	l.initMethodLimiter(method)
	ml, _ := l.limiters.Load(method)
	return ml.(*methodLimiter)
}

func (l *DynamicMethodQPSLimiter) Acquire(ctx context.Context) bool {
	ri := rpcinfo.GetRPCInfo(ctx)
	if ri == nil {
		klog.Errorf("Failed to get RPC info from context")
		return false
	}

	method := ri.To().Method()
	ml := l.GetMethodLimiter(method)

	// 令牌桶算法实现
	now := time.Now().UnixNano()
	lastRefill := atomic.LoadInt64(&ml.lastRefillTime)

	// 计算应该添加的令牌数
	elapsedSeconds := float64(now-lastRefill) / float64(time.Second.Nanoseconds())
	tokensToAdd := int32(elapsedSeconds * ml.refillRate)

	if tokensToAdd > 0 {
		// 尝试更新lastRefillTime，使用CAS操作确保原子性
		if atomic.CompareAndSwapInt64(&ml.lastRefillTime, lastRefill, now) {
			// 添加新令牌，但不超过桶容量
			limit := atomic.LoadInt32(&ml.currentLimit)
			availableTokens := atomic.LoadInt32(&ml.availableTokens)
			newTokens := availableTokens + tokensToAdd
			newTokens = min(newTokens, limit)
			atomic.StoreInt32(&ml.availableTokens, newTokens)
		}
	}

	// 尝试获取令牌
	for {
		available := atomic.LoadInt32(&ml.availableTokens)
		limit := atomic.LoadInt32(&ml.currentLimit)

		// 更新监控指标
		ml.qpsCurrentGauge.Set(float64(limit - available)) // 使用中的令牌数
		ml.qpsLimitGauge.Set(float64(limit))

		if available <= 0 {
			klog.Warnf("method %s exceeded qps limit: available=%d, limit=%d", method, available, limit)
			return false
		}

		// 原子方式减少令牌
		if atomic.CompareAndSwapInt32(&ml.availableTokens, available, available-1) {
			return true
		}

		// CAS失败，说明有竞争，重试
	}
}

func (l *DynamicMethodQPSLimiter) Status(ctx context.Context) (max, current int, interval time.Duration) {
	ri := rpcinfo.GetRPCInfo(ctx)
	method := ri.To().Method()

	ml := l.GetMethodLimiter(method)

	return int(ml.currentLimit), int(ml.availableTokens), time.Second
}

// 监控并调整每个方法的限流阈值
func (l *DynamicMethodQPSLimiter) monitor() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	fmt.Printf("Starting monitor for DynamicMethodQPSLimiter with CPU monitoring")

	// 添加 Registry 检查
	if mtl.Registry == nil {
		klog.Errorf("Prometheus Registry is not initialized")
		panic("Prometheus Registry must be initialized before using the limiter")
	}

	for range ticker.C {
		// 优先基于CPU使用率调整
		if l.cpuMonitor != nil {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			cpuUsage, err := l.cpuMonitor.getCPUUsage(ctx)
			cancel()

			if err == nil {

				// 根据CPU使用率全局调整所有服务限流
				var adjustFactor float64
				if cpuUsage > l.highCPUThreshold {
					// CPU高负载，降低所有限流阈值
					adjustFactor = (l.highCPUThreshold / cpuUsage) * 0.9
					klog.Warnf("High CPU usage (%.2f%%), reducing limits by factor %.2f\n", cpuUsage, adjustFactor)
				} else if cpuUsage < l.lowCPUThreshold {
					// CPU低负载，增加所有限流阈值
					adjustFactor = 1.0 + ((l.lowCPUThreshold-cpuUsage)/l.lowCPUThreshold)*0.5
					adjustFactor = min(adjustFactor, 1.5)
					klog.Infof("Low CPU usage (%.2f%%), increasing limits by factor %.2f\n", cpuUsage, adjustFactor)
				} else {
					// CPU使用率正常，微调
					adjustFactor = 1.0
				}

				if adjustFactor != 1.0 {
					l.limiters.Range(func(key, value interface{}) bool {
						method := key.(string)
						ml := value.(*methodLimiter)

						limit := atomic.LoadInt32(&ml.currentLimit)
						newLimit := int32(float64(limit) * adjustFactor)

						// 确保不超过合理范围
						if newLimit < l.defaultLimit/5 {
							newLimit = l.defaultLimit / 5
						} else if newLimit > l.defaultLimit*3 {
							newLimit = l.defaultLimit * 3
						}

						atomic.StoreInt32(&ml.currentLimit, newLimit)
						klog.Infof("Adjusting method %s limit: %d -> %d based on CPU usage", method, limit, newLimit)

						// 更新监控指标
						ml.qpsLimitGauge.Set(float64(newLimit))
						return true
					})
				}
			} else {
				klog.Errorf("Failed to get CPU usage: %v", err)
			}
		}

		// 针对monitor函数中的令牌桶处理部分
		l.limiters.Range(func(key, value interface{}) bool {
			method := key.(string)
			ml := value.(*methodLimiter)

			available := atomic.LoadInt32(&ml.availableTokens)
			limit := atomic.LoadInt32(&ml.currentLimit)

			usage := float64(limit-available) / float64(limit)
			klog.Infof("Method: %s, Available: %d, Limit: %d, Usage: %.2f\n", method, available, limit, usage)

			// 不需要重置令牌，令牌桶会自行填充
			// 只需更新监控指标
			ml.qpsCurrentGauge.Set(float64(limit - available))
			ml.qpsLimitGauge.Set(float64(limit))

			return true
		})
	}
}

// 手动调整限流阈值
func (l *DynamicMethodQPSLimiter) UpdateMethodLimit(method string, limit int32) error {
	if limit <= 0 {
		return errors.New("limit must be greater than 0")
	}

	if ml, ok := l.limiters.Load(method); ok {
		limiter := ml.(*methodLimiter)
		oldLimit := atomic.LoadInt32(&limiter.currentLimit)
		atomic.StoreInt32(&limiter.currentLimit, limit)

		// 更新令牌填充速率
		limiter.refillRate = float64(limit)

		// 如果新限制大于旧限制，立即添加额外的令牌
		if limit > oldLimit {
			available := atomic.LoadInt32(&limiter.availableTokens)
			newAvailable := available + (limit - oldLimit)
			if newAvailable > limit {
				newAvailable = limit
			}
			atomic.StoreInt32(&limiter.availableTokens, newAvailable)
		} else if available := atomic.LoadInt32(&limiter.availableTokens); available > limit {
			// 如果新限制小于旧限制且当前可用令牌超过新限制，裁剪令牌数
			atomic.StoreInt32(&limiter.availableTokens, limit)
		}

		klog.Infof("Update method %s QPS limit to %d", method, limit)
	} else {
		l.initMethodLimiter(method)
	}
	return nil
}
