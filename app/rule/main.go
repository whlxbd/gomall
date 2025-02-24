package main

import (
	"net"
	"os"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/server"
	"github.com/joho/godotenv"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/whlxbd/gomall/app/rule/biz/dal"
	"github.com/whlxbd/gomall/app/rule/conf"
	"github.com/whlxbd/gomall/app/rule/infra/rpc"
	"github.com/whlxbd/gomall/common/middleware/authenticator"
	"github.com/whlxbd/gomall/common/mtl"
	"github.com/whlxbd/gomall/common/serversuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule/ruleservice"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = conf.GetConf().Registry.RegistryAddress[0]
	MetricsPort  = conf.GetConf().Kitex.MetricsPort
)

func main() {
	_ = godotenv.Load()
	mtl.InitMetric(ServiceName, MetricsPort, RegistryAddr)
	dal.Init()
	rpc.InitClient()
	opts := kitexInit()

	svr := ruleservice.NewServer(new(RuleServiceImpl), opts...)

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

	opts = append(opts, server.WithSuite(serversuite.CommonServerSuite{
		CurrentServiceName: ServiceName,
		RegistryAddr:       RegistryAddr,
	}))

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

	opts = append(opts, server.WithMiddleware(authenticator.AuthenticatorMiddleware))
	return
}
