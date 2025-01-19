package aiorder

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"

	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder/aiorderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() aiorderservice.Client
	Service() string
	CreateAIOrder(ctx context.Context, Req *aiorder.CreateAIOrderReq, callOptions ...callopt.Option) (r *aiorder.CreateAIOrderResp, err error)
	GetAIOrder(ctx context.Context, Req *aiorder.GetAIOrderReq, callOptions ...callopt.Option) (r *aiorder.GetAIOrderResp, err error)
	CancelAIOrder(ctx context.Context, Req *aiorder.CancelAIOrderReq, callOptions ...callopt.Option) (r *aiorder.CancelAIOrderResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := aiorderservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient aiorderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() aiorderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateAIOrder(ctx context.Context, Req *aiorder.CreateAIOrderReq, callOptions ...callopt.Option) (r *aiorder.CreateAIOrderResp, err error) {
	return c.kitexClient.CreateAIOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) GetAIOrder(ctx context.Context, Req *aiorder.GetAIOrderReq, callOptions ...callopt.Option) (r *aiorder.GetAIOrderResp, err error) {
	return c.kitexClient.GetAIOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) CancelAIOrder(ctx context.Context, Req *aiorder.CancelAIOrderReq, callOptions ...callopt.Option) (r *aiorder.CancelAIOrderResp, err error) {
	return c.kitexClient.CancelAIOrder(ctx, Req, callOptions...)
}
