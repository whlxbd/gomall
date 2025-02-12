package rpc

import (
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/whlxbd/gomall/app/checkout/conf"
	"github.com/whlxbd/gomall/common/clientsuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth/authservice"
)

var (
	AuthClient authservice.Client
	once       sync.Once
	err        error
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
	})
}

func initAuthClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}),
	}

	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		klog.Errorf("failed to init client: %v", err)
	}
}
