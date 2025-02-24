package rpc

import (
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/whlxbd/gomall/app/cart/conf"
	"github.com/whlxbd/gomall/common/clientsuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient   userservice.Client
	once         sync.Once
	err          error
	ServiceName  = conf.GetConf().Kitex.Service
	RegistryAddr = os.Getenv("REGISTRY_ADDR")
)

func InitClient() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: ServiceName,
			RegistryAddr:       RegistryAddr,
		}),
	}

	UserClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		klog.Fatalf("failed to init client: %v", err)
	}
}
