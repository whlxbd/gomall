package main

import (
	"net"
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/whlxbd/gomall/app/cart/biz/dal"

	// "github.com/whlxbd/gomall/app/cart/biz/dal/mq"
	"github.com/whlxbd/gomall/app/cart/conf"
	"github.com/whlxbd/gomall/app/cart/infra/rpc"
	"github.com/whlxbd/gomall/common/middleware/authenticator"
	"github.com/whlxbd/gomall/common/mtl"
	"github.com/whlxbd/gomall/common/utils/pool"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	_ = godotenv.Load()
	pool.Init()
	dal.Init()
	rpc.InitClient()
	// defer mq.GetCartMQ().Close()
	defer pool.Release()

	mtl.InitMetric(conf.GetConf().Kitex.Service, conf.GetConf().Kitex.MetricsPort, os.Getenv("REGISTRY_ADDR"))
	opts := kitexInit()

	svr := cartservice.NewServer(new(CartServiceImpl), opts...)

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

	opts = append(opts, server.WithMetaHandler(transmeta.ServerHTTP2Handler))
	opts = append(opts, server.WithMetaHandler(transmeta.ServerTTHeaderHandler))
	opts = append(opts, server.WithMiddleware(authenticator.AuthenticatorMiddleware))
	return
}
