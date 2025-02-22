package auth

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Create(ctx context.Context, req *auth.CreateReq, callOptions ...callopt.Option) (resp *auth.CreateResp, err error) {
	resp, err = defaultClient.Create(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Create call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func List(ctx context.Context, req *auth.ListReq, callOptions ...callopt.Option) (resp *auth.ListResp, err error) {
	resp, err = defaultClient.List(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "List call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Delete(ctx context.Context, req *auth.DeleteReq, callOptions ...callopt.Option) (resp *auth.DeleteResp, err error) {
	resp, err = defaultClient.Delete(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Get(ctx context.Context, req *auth.GetReq, callOptions ...callopt.Option) (resp *auth.GetResp, err error) {
	resp, err = defaultClient.Get(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Get call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Update(ctx context.Context, req *auth.UpdateReq, callOptions ...callopt.Option) (resp *auth.UpdateResp, err error) {
	resp, err = defaultClient.Update(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Update call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
