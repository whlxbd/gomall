package main

import (
	"net"
	"os"
	"time"
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/whlxbd/gomall/app/order/biz/dal"
	"github.com/whlxbd/gomall/app/order/biz/dal/mq"
	"github.com/whlxbd/gomall/app/order/conf"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/order/orderservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"github.com/whlxbd/gomall/common/utils/pool"
)

func main() {
	_ = godotenv.Load()
	opts := kitexInit()
	pool.Init()
	dal.Init()
	defer pool.Release()

    // 初始化MQ
    if err := mq.Init(os.Getenv("RMQENDPOINT")); err != nil {
        klog.Fatalf("init mq failed: %v", err)
    }

    // 启动消费者，使用context控制生命周期
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()


	_ = pool.Submit(func() {
		mq.StartOrderConsumer(ctx)
	})

	svr := orderservice.NewServer(new(OrderServiceImpl), opts...)

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
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Kitex.LogFileName,
			MaxSize:    conf.GetConf().Kitex.LogMaxSize,
			MaxBackups: conf.GetConf().Kitex.LogMaxBackups,
			MaxAge:     conf.GetConf().Kitex.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})
	return
}
