package auth

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"

	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth/ruleservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() ruleservice.Client
	Service() string
	Create(ctx context.Context, Req *auth.CreateReq, callOptions ...callopt.Option) (r *auth.CreateResp, err error)
	List(ctx context.Context, Req *auth.ListReq, callOptions ...callopt.Option) (r *auth.ListResp, err error)
	Delete(ctx context.Context, Req *auth.DeleteReq, callOptions ...callopt.Option) (r *auth.DeleteResp, err error)
	Get(ctx context.Context, Req *auth.GetReq, callOptions ...callopt.Option) (r *auth.GetResp, err error)
	Update(ctx context.Context, Req *auth.UpdateReq, callOptions ...callopt.Option) (r *auth.UpdateResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := ruleservice.NewClient(dstService, opts...)
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
	kitexClient ruleservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() ruleservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Create(ctx context.Context, Req *auth.CreateReq, callOptions ...callopt.Option) (r *auth.CreateResp, err error) {
	return c.kitexClient.Create(ctx, Req, callOptions...)
}

func (c *clientImpl) List(ctx context.Context, Req *auth.ListReq, callOptions ...callopt.Option) (r *auth.ListResp, err error) {
	return c.kitexClient.List(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *auth.DeleteReq, callOptions ...callopt.Option) (r *auth.DeleteResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}

func (c *clientImpl) Get(ctx context.Context, Req *auth.GetReq, callOptions ...callopt.Option) (r *auth.GetResp, err error) {
	return c.kitexClient.Get(ctx, Req, callOptions...)
}

func (c *clientImpl) Update(ctx context.Context, Req *auth.UpdateReq, callOptions ...callopt.Option) (r *auth.UpdateResp, err error) {
	return c.kitexClient.Update(ctx, Req, callOptions...)
}
