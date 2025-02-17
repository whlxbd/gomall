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
	QueryOrder(ctx context.Context, Req *aiorder.QueryOrderReq, callOptions ...callopt.Option) (r *aiorder.QueryOrderResp, err error)
	SimulateOrder(ctx context.Context, Req *aiorder.SimulateOrderReq, callOptions ...callopt.Option) (r *aiorder.SimulateOrderResp, err error)
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

func (c *clientImpl) QueryOrder(ctx context.Context, Req *aiorder.QueryOrderReq, callOptions ...callopt.Option) (r *aiorder.QueryOrderResp, err error) {
	return c.kitexClient.QueryOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) SimulateOrder(ctx context.Context, Req *aiorder.SimulateOrderReq, callOptions ...callopt.Option) (r *aiorder.SimulateOrderResp, err error) {
	return c.kitexClient.SimulateOrder(ctx, Req, callOptions...)
}
