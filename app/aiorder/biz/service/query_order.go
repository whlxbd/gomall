package service

import (
	"context"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/aiorder/agent"
	"github.com/whlxbd/gomall/app/aiorder/biz/dal/mysql"
	"github.com/whlxbd/gomall/app/aiorder/biz/model"
	"github.com/whlxbd/gomall/common/utils/pool"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
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
	_ = pool.Submit(func(){
		err_db := model.CreateMessage(s.ctx, mysql.DB, req.UserId, req.Content, "query")
		if err_db != nil {
			klog.Errorf("CreateMessage failed: %v", err_db)
			fmt.Printf("CreateMessage failed: %+v\n", err_db)
		}
		for i := 1; i <= 2; i++ {
			err_db = model.CreateMessage(s.ctx, mysql.DB, req.UserId, req.Content, "query")
			if err_db != nil {
				klog.Errorf("%d try CreateMessage failed: %v", i, err_db)
				fmt.Printf("%d try CreateMessage failed: %+v\n", i, err_db)
			}
		}
	})

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
