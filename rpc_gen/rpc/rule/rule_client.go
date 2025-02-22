package rule

import (
	"context"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"

	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule/ruleservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() ruleservice.Client
	Service() string
	Create(ctx context.Context, Req *rule.CreateReq, callOptions ...callopt.Option) (r *rule.CreateResp, err error)
	List(ctx context.Context, Req *rule.ListReq, callOptions ...callopt.Option) (r *rule.ListResp, err error)
	Delete(ctx context.Context, Req *rule.DeleteReq, callOptions ...callopt.Option) (r *rule.DeleteResp, err error)
	Get(ctx context.Context, Req *rule.GetReq, callOptions ...callopt.Option) (r *rule.GetResp, err error)
	Update(ctx context.Context, Req *rule.UpdateReq, callOptions ...callopt.Option) (r *rule.UpdateResp, err error)
	AddWhiteRouter(ctx context.Context, Req *rule.AddWhiteRouterReq, callOptions ...callopt.Option) (r *rule.AddWhiteRouterResp, err error)
	GetWhiteList(ctx context.Context, Req *rule.GetWhiteListReq, callOptions ...callopt.Option) (r *rule.GetWhiteListResp, err error)
	DeleteWhiteRouter(ctx context.Context, Req *rule.DeleteWhiteRouterReq, callOptions ...callopt.Option) (r *rule.DeleteWhiteRouterResp, err error)
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

func (c *clientImpl) Create(ctx context.Context, Req *rule.CreateReq, callOptions ...callopt.Option) (r *rule.CreateResp, err error) {
	return c.kitexClient.Create(ctx, Req, callOptions...)
}

func (c *clientImpl) List(ctx context.Context, Req *rule.ListReq, callOptions ...callopt.Option) (r *rule.ListResp, err error) {
	return c.kitexClient.List(ctx, Req, callOptions...)
}

func (c *clientImpl) Delete(ctx context.Context, Req *rule.DeleteReq, callOptions ...callopt.Option) (r *rule.DeleteResp, err error) {
	return c.kitexClient.Delete(ctx, Req, callOptions...)
}

func (c *clientImpl) Get(ctx context.Context, Req *rule.GetReq, callOptions ...callopt.Option) (r *rule.GetResp, err error) {
	return c.kitexClient.Get(ctx, Req, callOptions...)
}

func (c *clientImpl) Update(ctx context.Context, Req *rule.UpdateReq, callOptions ...callopt.Option) (r *rule.UpdateResp, err error) {
	return c.kitexClient.Update(ctx, Req, callOptions...)
}

func (c *clientImpl) AddWhiteRouter(ctx context.Context, Req *rule.AddWhiteRouterReq, callOptions ...callopt.Option) (r *rule.AddWhiteRouterResp, err error) {
	return c.kitexClient.AddWhiteRouter(ctx, Req, callOptions...)
}

func (c *clientImpl) GetWhiteList(ctx context.Context, Req *rule.GetWhiteListReq, callOptions ...callopt.Option) (r *rule.GetWhiteListResp, err error) {
	return c.kitexClient.GetWhiteList(ctx, Req, callOptions...)
}

func (c *clientImpl) DeleteWhiteRouter(ctx context.Context, Req *rule.DeleteWhiteRouterReq, callOptions ...callopt.Option) (r *rule.DeleteWhiteRouterResp, err error) {
	return c.kitexClient.DeleteWhiteRouter(ctx, Req, callOptions...)
}
