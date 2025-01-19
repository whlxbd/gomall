package service

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
)

type CancelAIOrderService struct {
	ctx context.Context
} // NewCancelAIOrderService new CancelAIOrderService
func NewCancelAIOrderService(ctx context.Context) *CancelAIOrderService {
	return &CancelAIOrderService{ctx: ctx}
}

// Run create note info
func (s *CancelAIOrderService) Run(req *aiorder.CancelAIOrderReq) (resp *aiorder.CancelAIOrderResp, err error) {
	// Finish your business logic.

	return
}
