package rule

import (
	"context"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Create(ctx context.Context, req *rule.CreateReq, callOptions ...callopt.Option) (resp *rule.CreateResp, err error) {
	resp, err = defaultClient.Create(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Create call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func List(ctx context.Context, req *rule.ListReq, callOptions ...callopt.Option) (resp *rule.ListResp, err error) {
	resp, err = defaultClient.List(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "List call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Delete(ctx context.Context, req *rule.DeleteReq, callOptions ...callopt.Option) (resp *rule.DeleteResp, err error) {
	resp, err = defaultClient.Delete(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Delete call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Get(ctx context.Context, req *rule.GetReq, callOptions ...callopt.Option) (resp *rule.GetResp, err error) {
	resp, err = defaultClient.Get(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Get call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Update(ctx context.Context, req *rule.UpdateReq, callOptions ...callopt.Option) (resp *rule.UpdateResp, err error) {
	resp, err = defaultClient.Update(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Update call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddWhiteRouter(ctx context.Context, req *rule.AddWhiteRouterReq, callOptions ...callopt.Option) (resp *rule.AddWhiteRouterResp, err error) {
	resp, err = defaultClient.AddWhiteRouter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddWhiteRouter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetWhiteList(ctx context.Context, req *rule.GetWhiteListReq, callOptions ...callopt.Option) (resp *rule.GetWhiteListResp, err error) {
	resp, err = defaultClient.GetWhiteList(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetWhiteList call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func DeleteWhiteRouter(ctx context.Context, req *rule.DeleteWhiteRouterReq, callOptions ...callopt.Option) (resp *rule.DeleteWhiteRouterResp, err error) {
	resp, err = defaultClient.DeleteWhiteRouter(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "DeleteWhiteRouter call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
