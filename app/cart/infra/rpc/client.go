package rpc

import (
	"sync"
	"os"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
		
	"github.com/whlxbd/gomall/common/clientsuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
	"github.com/whlxbd/gomall/app/cart/conf"
)

var (
	ProductClient productcatalogservice.Client
	AuthClient    authservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
		initAuthClient()
	})
}

func initProductClient() {
	opts := []client.Option {
		client.WithSuite(clientsuite.CommonGrpcClientSuite {
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr: os.Getenv("REGISTRY_ADDR"),
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		klog.Fatalf("failed to init client: %v", err)
	}
}

func initAuthClient() {
	opts := []client.Option {
		client.WithSuite(clientsuite.CommonGrpcClientSuite {
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr: os.Getenv("REGISTRY_ADDR"),
		}),
	}

	AuthClient, err = authservice.NewClient("cart", opts...)
	if err != nil {
		klog.Errorf("failed to init client: %v", err)
	}
}