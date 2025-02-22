package auth

import (
	"context"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func DeliverTokenByRPC(ctx context.Context, req *auth.DeliverTokenReq, callOptions ...callopt.Option) (resp *auth.DeliveryResp, err error) {
	resp, err = defaultClient.DeliverTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeliverTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func VerifyTokenByRPC(ctx context.Context, req *auth.VerifyTokenReq, callOptions ...callopt.Option) (resp *auth.VerifyResp, err error) {
	resp, err = defaultClient.VerifyTokenByRPC(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "VerifyTokenByRPC call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetPayload(ctx context.Context, req *auth.GetPayloadReq, callOptions ...callopt.Option) (resp *auth.GetPayloadResp, err error) {
	resp, err = defaultClient.GetPayload(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetPayload call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckPermission(ctx context.Context, req *auth.CheckPermissionReq, callOptions ...callopt.Option) (resp *auth.CheckPermissionResp, err error) {
	resp, err = defaultClient.CheckPermission(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckPermission call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckWhite(ctx context.Context, req *auth.CheckWhiteReq, callOptions ...callopt.Option) (resp *auth.CheckWhiteResp, err error) {
	resp, err = defaultClient.CheckWhite(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckWhite call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func LoadPolicy(ctx context.Context, req *auth.LoadPolicyReq, callOptions ...callopt.Option) (resp *auth.LoadPolicyResp, err error) {
	resp, err = defaultClient.LoadPolicy(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "LoadPolicy call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RemovePolicy(ctx context.Context, req *auth.RemovePolicyReq, callOptions ...callopt.Option) (resp *auth.RemovePolicyResp, err error) {
	resp, err = defaultClient.RemovePolicy(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RemovePolicy call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
