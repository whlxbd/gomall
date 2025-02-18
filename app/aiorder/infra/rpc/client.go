package rpc

import (
	"sync"
	"os"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
		
	"github.com/whlxbd/gomall/common/clientsuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/whlxbd/gomall/app/aiorder/conf"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart/cartservice"
)

var (
	OrderClient   orderservice.Client
	CartClient	  cartservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initorderClient()
		initcartClient()
	})
}

func initorderClient() {
	opts := []client.Option {
		client.WithSuite(clientsuite.CommonGrpcClientSuite {
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr: os.Getenv("REGISTRY_ADDR"),
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		klog.Fatalf("failed to init client: %v", err)
	}
}

func initcartClient() {
	opts := []client.Option {
		client.WithSuite(clientsuite.CommonGrpcClientSuite {
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr: os.Getenv("REGISTRY_ADDR"),
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		klog.Fatalf("failed to init cart client: %v", err)
	}
}