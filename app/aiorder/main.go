package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"context"

	"github.com/joho/godotenv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/whlxbd/gomall/app/aiorder/agent"
	"github.com/whlxbd/gomall/app/aiorder/biz/dal"
	"github.com/whlxbd/gomall/app/aiorder/conf"
	"github.com/whlxbd/gomall/app/aiorder/infra/rpc"
	"github.com/whlxbd/gomall/common/limiter"
	"github.com/whlxbd/gomall/common/middleware/authenticator"
	"github.com/whlxbd/gomall/common/utils/pool"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder/aiorderservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/cloudwego/kitex/pkg/endpoint"
)

func main() {
	enverr := godotenv.Load()
	if enverr != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", enverr))
	}
	pool.Init()
	dal.Init()
	agent.Init()
	rpc.InitClient()
	defer pool.Release()

	opts := kitexInit()

	svr := aiorderservice.NewServer(new(AIOrderServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// consul
	r, err := consul.NewConsulRegister(os.Getenv("REGISTRY_ADDR")) // 使用配置中的 Consul 地址
	if err != nil {
		klog.Fatal(err)
	}
	opts = append(opts, server.WithRegistry(r))

	// klog
	var flushInterval time.Duration
	if os.Getenv("GO_ENV") == "online" {
		flushInterval = time.Minute
	} else {
		flushInterval = time.Second
	}
	logger := kitexlogrus.NewLogger(kitexlogrus.WithLogger(kitexlogrus.NewLogger().Logger()))
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: flushInterval,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})

	// 创建限流器
	qpsLimiter := limiter.NewDynamicMethodQPSLimiter(100)
	klog.Infof("Limiter initialized: %+v", qpsLimiter)

	// 显式注册中间件
	opts = append(opts, server.WithMiddleware(func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req, resp interface{}) (err error) {
			if !qpsLimiter.Acquire(ctx) {
				klog.Warnf("Request limited by QPS limiter")
				panic("Request limited by QPS limiter")
			}
			return next(ctx, req, resp)
		}
	}))

	opts = append(opts, server.WithMiddleware(authenticator.AuthenticatorMiddleware))
	return
}
