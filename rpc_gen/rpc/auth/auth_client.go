package auth

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"

	"github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth/authservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() authservice.Client
	Service() string
	DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error)
	VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error)
	GetPayload(ctx context.Context, Req *auth.GetPayloadReq, callOptions ...callopt.Option) (r *auth.GetPayloadResp, err error)
	CheckPermission(ctx context.Context, Req *auth.CheckPermissionReq, callOptions ...callopt.Option) (r *auth.CheckPermissionResp, err error)
	CheckWhite(ctx context.Context, Req *auth.CheckWhiteReq, callOptions ...callopt.Option) (r *auth.CheckWhiteResp, err error)
	LoadPolicy(ctx context.Context, Req *auth.LoadPolicyReq, callOptions ...callopt.Option) (r *auth.LoadPolicyResp, err error)
	RemovePolicy(ctx context.Context, Req *auth.RemovePolicyReq, callOptions ...callopt.Option) (r *auth.RemovePolicyResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := authservice.NewClient(dstService, opts...)
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
	kitexClient authservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() authservice.Client {
	return c.kitexClient
}

func (c *clientImpl) DeliverTokenByRPC(ctx context.Context, Req *auth.DeliverTokenReq, callOptions ...callopt.Option) (r *auth.DeliveryResp, err error) {
	return c.kitexClient.DeliverTokenByRPC(ctx, Req, callOptions...)
}

func (c *clientImpl) VerifyTokenByRPC(ctx context.Context, Req *auth.VerifyTokenReq, callOptions ...callopt.Option) (r *auth.VerifyResp, err error) {
	return c.kitexClient.VerifyTokenByRPC(ctx, Req, callOptions...)
}

func (c *clientImpl) GetPayload(ctx context.Context, Req *auth.GetPayloadReq, callOptions ...callopt.Option) (r *auth.GetPayloadResp, err error) {
	return c.kitexClient.GetPayload(ctx, Req, callOptions...)
}

func (c *clientImpl) CheckPermission(ctx context.Context, Req *auth.CheckPermissionReq, callOptions ...callopt.Option) (r *auth.CheckPermissionResp, err error) {
	return c.kitexClient.CheckPermission(ctx, Req, callOptions...)
}

func (c *clientImpl) CheckWhite(ctx context.Context, Req *auth.CheckWhiteReq, callOptions ...callopt.Option) (r *auth.CheckWhiteResp, err error) {
	return c.kitexClient.CheckWhite(ctx, Req, callOptions...)
}

func (c *clientImpl) LoadPolicy(ctx context.Context, Req *auth.LoadPolicyReq, callOptions ...callopt.Option) (r *auth.LoadPolicyResp, err error) {
	return c.kitexClient.LoadPolicy(ctx, Req, callOptions...)
}

func (c *clientImpl) RemovePolicy(ctx context.Context, Req *auth.RemovePolicyReq, callOptions ...callopt.Option) (r *auth.RemovePolicyResp, err error) {
	return c.kitexClient.RemovePolicy(ctx, Req, callOptions...)
}
