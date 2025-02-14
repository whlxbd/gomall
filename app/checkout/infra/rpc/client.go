package rpc

import (
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/whlxbd/gomall/app/checkout/conf"
	"github.com/whlxbd/gomall/common/clientsuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/product/productcatalogservice"
)

var (
	ProductClient productcatalogservice.Client
	AuthClient    authservice.Client
	OrderClient   orderservice.Client
	PaymentClient paymentservice.Client
	CartClient    cartservice.Client
	once          sync.Once
	err           error
)

func InitClient() {
	once.Do(func() {
		initProductClient()
		initAuthClient()
		initOrderClient()
		initPaymentClient()
		initCartClient()
	})
}

func initProductClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       os.Getenv("REGISTRY_ADDR"),
		}),
	}

	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	if err != nil {
		klog.Fatalf("failed to init client: %v", err)
	}
}

func initAuthClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       os.Getenv("REGISTRY_ADDR"),
		}),
	}

	AuthClient, err = authservice.NewClient("auth", opts...)
	if err != nil {
		klog.Errorf("failed to init client: %v", err)
	}
}

func initOrderClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       os.Getenv("REGISTRY_ADDR"),
		}),
	}

	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		klog.Errorf("failed to init client: %v", err)
	}
}

func initPaymentClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       os.Getenv("REGISTRY_ADDR"),
		}),
	}

	PaymentClient, err = paymentservice.NewClient("payment", opts...)
	if err != nil {
		klog.Errorf("failed to init client: %v", err)
	}
}

func initCartClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       os.Getenv("REGISTRY_ADDR"),
		}),
	}

	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		klog.Errorf("failed to init client: %v", err)
	}
}
