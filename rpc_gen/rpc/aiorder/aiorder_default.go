package aiorder

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateAIOrder(ctx context.Context, req *aiorder.CreateAIOrderReq, callOptions ...callopt.Option) (resp *aiorder.CreateAIOrderResp, err error) {
	resp, err = defaultClient.CreateAIOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateAIOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetAIOrder(ctx context.Context, req *aiorder.GetAIOrderReq, callOptions ...callopt.Option) (resp *aiorder.GetAIOrderResp, err error) {
	resp, err = defaultClient.GetAIOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetAIOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CancelAIOrder(ctx context.Context, req *aiorder.CancelAIOrderReq, callOptions ...callopt.Option) (resp *aiorder.CancelAIOrderResp, err error) {
	resp, err = defaultClient.CancelAIOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CancelAIOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
