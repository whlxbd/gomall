package service

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
)

type GetAIOrderService struct {
	ctx context.Context
} // NewGetAIOrderService new GetAIOrderService
func NewGetAIOrderService(ctx context.Context) *GetAIOrderService {
	return &GetAIOrderService{ctx: ctx}
}

// Run create note info
func (s *GetAIOrderService) Run(req *aiorder.GetAIOrderReq) (resp *aiorder.GetAIOrderResp, err error) {
	// Finish your business logic.

	return
}
