package service

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
)

type CreateAIOrderService struct {
	ctx context.Context
} // NewCreateAIOrderService new CreateAIOrderService
func NewCreateAIOrderService(ctx context.Context) *CreateAIOrderService {
	return &CreateAIOrderService{ctx: ctx}
}

// Run create note info
func (s *CreateAIOrderService) Run(req *aiorder.CreateAIOrderReq) (resp *aiorder.CreateAIOrderResp, err error) {
	// Finish your business logic.

	return
}
