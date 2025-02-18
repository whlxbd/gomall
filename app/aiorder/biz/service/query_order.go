package service

import (
	"context"
	"fmt"

	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
	"github.com/whlxbd/gomall/app/aiorder/agent"
	"github.com/cloudwego/kitex/pkg/klog"
)

type QueryOrderService struct {
	ctx context.Context
} // NewQueryOrderService new QueryOrderService
func NewQueryOrderService(ctx context.Context) *QueryOrderService {
	return &QueryOrderService{ctx: ctx}
}

// Run create note info
func (s *QueryOrderService) Run(req *aiorder.QueryOrderReq) (resp *aiorder.QueryOrderResp, err error) {
	// Finish your business logic.
	usrMsg := &agent.UserMessage {
		UserId: req.UserId,
		Content: fmt.Sprintf("当前用户ID: %d，用户给出的订单查询文本: %s", req.UserId, req.Content),
	}
	fmt.Printf("usrMsg: %+v\n", usrMsg)

	sr, err := agent.GetFlow().Invoke(s.ctx, usrMsg)
	if err != nil {
		klog.Errorf("Invoke flow failed: %v", err)
		fmt.Printf("Invoke flow failed: %+v\n", err)
		return &aiorder.QueryOrderResp{
            Result: "抱歉，查询失败，请稍后重试",
        }, err
	}
	resp = &aiorder.QueryOrderResp{
		Result: sr.Content,
	}
	return
}
