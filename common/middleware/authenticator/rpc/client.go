package rpc

import (
	"os"
	"sync"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"

	"github.com/whlxbd/gomall/app/checkout/conf"
	"github.com/whlxbd/gomall/common/clientsuite"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule/ruleservice"
)

var (
	AuthClient authservice.Client
	RuleClient ruleservice.Client
	once       sync.Once
	err        error
)

func InitClient() {
	once.Do(func() {
		initAuthClient()
		initRuleClient()
	})
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

func initRuleClient() {
	opts := []client.Option{
		client.WithSuite(clientsuite.CommonGrpcClientSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       os.Getenv("REGISTRY_ADDR"),
		}),
	}

	RuleClient, err = ruleservice.NewClient("rule", opts...)
	if err != nil {
		klog.Errorf("failed to init client: %v", err)
	}
}
