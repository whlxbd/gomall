package rpc

import (
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/whlxbd/gomall/app/cart/conf"
	"github.com/whlxbd/gomall/app/user/kitex_gen/user/userservice"
	"github.com/whlxbd/gomall/common/clientsuite"
)

var (
	UserClient userservice.Client
	once       sync.Once
	err        error
)

func InitClient() {
	once.Do(func() {
		initUserClient()
	})
}

func initUserClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       os.Getenv("REGISTRY_ADDR"),
		}),
	}

	UserClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		klog.Fatalf("failed to init client: %v", err)
	}
}
