package aiorder

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func QueryOrder(ctx context.Context, req *aiorder.QueryOrderReq, callOptions ...callopt.Option) (resp *aiorder.QueryOrderResp, err error) {
	resp, err = defaultClient.QueryOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "QueryOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SimulateOrder(ctx context.Context, req *aiorder.SimulateOrderReq, callOptions ...callopt.Option) (resp *aiorder.SimulateOrderResp, err error) {
	resp, err = defaultClient.SimulateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SimulateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
