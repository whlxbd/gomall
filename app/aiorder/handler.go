package main

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
	"github.com/whlxbd/gomall/rpc_gen/biz/service"
)

// AIOrderServiceImpl implements the last service interface defined in the IDL.
type AIOrderServiceImpl struct{}

// CreateAIOrder implements the AIOrderServiceImpl interface.
func (s *AIOrderServiceImpl) CreateAIOrder(ctx context.Context, req *aiorder.CreateAIOrderReq) (resp *aiorder.CreateAIOrderResp, err error) {
	resp, err = service.NewCreateAIOrderService(ctx).Run(req)

	return resp, err
}

// GetAIOrder implements the AIOrderServiceImpl interface.
func (s *AIOrderServiceImpl) GetAIOrder(ctx context.Context, req *aiorder.GetAIOrderReq) (resp *aiorder.GetAIOrderResp, err error) {
	resp, err = service.NewGetAIOrderService(ctx).Run(req)

	return resp, err
}

// CancelAIOrder implements the AIOrderServiceImpl interface.
func (s *AIOrderServiceImpl) CancelAIOrder(ctx context.Context, req *aiorder.CancelAIOrderReq) (resp *aiorder.CancelAIOrderResp, err error) {
	resp, err = service.NewCancelAIOrderService(ctx).Run(req)

	return resp, err
}
