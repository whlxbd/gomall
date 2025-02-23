package limiter

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/whlxbd/gomall/common/mtl"
)

type methodLimiter struct {
	currentLimit    int32 //当前限制
	currentRequests int32 //当前剩余
	qpsLimitGauge   prometheus.Gauge
	qpsCurrentGauge prometheus.Gauge
}

type DynamicMethodQPSLimiter struct {
	limiters     sync.Map
	defaultLimit int32
}

func NewDynamicMethodQPSLimiter(defaultLimit int32) *DynamicMethodQPSLimiter {
	limiter := &DynamicMethodQPSLimiter{
		defaultLimit: defaultLimit,
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
    // 获得方法对应限流器
    ri := rpcinfo.GetRPCInfo(ctx)
    if ri == nil {
        klog.Errorf("Failed to get RPC info from context")
        return false
    }
    
    method := ri.To().Method()  // 修改这里
    klog.Infof("Acquire called for method: %s", method)
    
    ml := l.GetMethodLimiter(method)
    
    current := atomic.LoadInt32(&ml.currentRequests)
    limit := atomic.LoadInt32(&ml.currentLimit)
    
    // 更新监控指标
    ml.qpsCurrentGauge.Set(float64(current))
    ml.qpsLimitGauge.Set(float64(limit))
    
    if current >= limit {
        klog.Warnf("method %s exceeded qps limit exceeded: current=%d, limit=%d ", method, current, limit)
        return false
    }
    
    atomic.AddInt32(&ml.currentRequests, 1)
    return true
}
func (l *DynamicMethodQPSLimiter) Status(ctx context.Context) (max, current int, interval time.Duration) {
	ri := rpcinfo.GetRPCInfo(ctx)
	method := ri.To().Method()
	fmt.Printf("Status Method: %s\n", method)

	ml := l.GetMethodLimiter(method)

	return int(ml.currentLimit), int(ml.currentRequests), time.Second
}

// 监控并调整每个方法的限流阈值
func (l *DynamicMethodQPSLimiter) monitor() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Printf("Starting monitor for DynamicMethodQPSLimiter\n")
	// 添加 Registry 检查
    if mtl.Registry == nil {
        klog.Errorf("Prometheus Registry is not initialized")
        panic("Prometheus Registry must be initialized before using the limiter")
    }
	for range ticker.C {
		l.limiters.Range(func(key, value interface{}) bool {
			method := key.(string)
			ml := value.(*methodLimiter)

			current := atomic.LoadInt32(&ml.currentRequests)
			limit := atomic.LoadInt32(&ml.currentLimit)

			usage := float64(current) / float64(limit)
			if usage > 0.8 {
				// 使用率较高
				atomic.AddInt32(&ml.currentLimit, limit/10)
				klog.Infof("Increase method %s QPS limit to %d", method, ml.currentLimit)
			} else if usage < 0.3 {
				// 使用率较低，减少限流阈值
				newLimit := limit - limit/10
				if newLimit >= l.defaultLimit {
					atomic.StoreInt32(&ml.currentLimit, newLimit)
					klog.Infof("Decrease method %s QPS limit to %d", method, ml.currentLimit)
				}
			}

			atomic.StoreInt32(&ml.currentRequests, 0)

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
		atomic.StoreInt32(&ml.(*methodLimiter).currentLimit, limit)
		atomic.StoreInt32(&ml.(*methodLimiter).currentRequests, 0)

		klog.Infof("Update method %s QPS limit to %d", method, limit)
	} else {
		l.initMethodLimiter(method)
	}
	return nil
}
