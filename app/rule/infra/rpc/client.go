package rpc

import (
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/whlxbd/gomall/app/rule/conf"
	"github.com/whlxbd/gomall/common/clientsuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth/authservice"
)

var (
	AuthClient   authservice.Client
	once         sync.Once
	err          error
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = os.Getenv("REGISTRY_ADDR")
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
	})
}

func initAuthClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		klog.Fatalf("failed to init client: %v", err)
	}
}
