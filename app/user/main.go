package main

import (
	"net"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/whlxbd/gomall/app/user/biz/dal"
	"github.com/whlxbd/gomall/app/user/infra/rpc"
	"github.com/whlxbd/gomall/common/middleware/authenticator"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/whlxbd/gomall/app/user/conf"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/user/userservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	_ "github.com/cloudwego/kitex/pkg/remote/codec/protobuf/encoding/gzip"
)

func main() {
	_ = godotenv.Load()
	opts := kitexInit()

	dal.Init()
	rpc.InitClient()

	svr := userservice.NewServer(new(UserServiceImpl), opts...)

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
		FlushInterval: time.Second,
	}
	klog.SetOutput(asyncWriter)
	server.RegisterShutdownHook(func() {
		asyncWriter.Sync()
	})

	opts = append(opts, server.WithMiddleware(authenticator.AuthenticatorMiddleware))
	return
}
