// Code generated by Kitex v0.9.1. DO NOT EDIT.

package ruleservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Create(ctx context.Context, Req *rule.CreateReq, callOptions ...callopt.Option) (r *rule.CreateResp, err error)
	List(ctx context.Context, Req *rule.ListReq, callOptions ...callopt.Option) (r *rule.ListResp, err error)
	Delete(ctx context.Context, Req *rule.DeleteReq, callOptions ...callopt.Option) (r *rule.DeleteResp, err error)
	Get(ctx context.Context, Req *rule.GetReq, callOptions ...callopt.Option) (r *rule.GetResp, err error)
	Update(ctx context.Context, Req *rule.UpdateReq, callOptions ...callopt.Option) (r *rule.UpdateResp, err error)
	AddWhiteRouter(ctx context.Context, Req *rule.AddWhiteRouterReq, callOptions ...callopt.Option) (r *rule.AddWhiteRouterResp, err error)
	GetWhiteList(ctx context.Context, Req *rule.GetWhiteListReq, callOptions ...callopt.Option) (r *rule.GetWhiteListResp, err error)
	DeleteWhiteRouter(ctx context.Context, Req *rule.DeleteWhiteRouterReq, callOptions ...callopt.Option) (r *rule.DeleteWhiteRouterResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRuleServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRuleServiceClient struct {
	*kClient
}

func (p *kRuleServiceClient) Create(ctx context.Context, Req *rule.CreateReq, callOptions ...callopt.Option) (r *rule.CreateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Create(ctx, Req)
}

func (p *kRuleServiceClient) List(ctx context.Context, Req *rule.ListReq, callOptions ...callopt.Option) (r *rule.ListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.List(ctx, Req)
}

func (p *kRuleServiceClient) Delete(ctx context.Context, Req *rule.DeleteReq, callOptions ...callopt.Option) (r *rule.DeleteResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Delete(ctx, Req)
}

func (p *kRuleServiceClient) Get(ctx context.Context, Req *rule.GetReq, callOptions ...callopt.Option) (r *rule.GetResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Get(ctx, Req)
}

func (p *kRuleServiceClient) Update(ctx context.Context, Req *rule.UpdateReq, callOptions ...callopt.Option) (r *rule.UpdateResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Update(ctx, Req)
}

func (p *kRuleServiceClient) AddWhiteRouter(ctx context.Context, Req *rule.AddWhiteRouterReq, callOptions ...callopt.Option) (r *rule.AddWhiteRouterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AddWhiteRouter(ctx, Req)
}

func (p *kRuleServiceClient) GetWhiteList(ctx context.Context, Req *rule.GetWhiteListReq, callOptions ...callopt.Option) (r *rule.GetWhiteListResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetWhiteList(ctx, Req)
}

func (p *kRuleServiceClient) DeleteWhiteRouter(ctx context.Context, Req *rule.DeleteWhiteRouterReq, callOptions ...callopt.Option) (r *rule.DeleteWhiteRouterResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteWhiteRouter(ctx, Req)
}
