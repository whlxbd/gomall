package service

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
)

type SimulateOrderService struct {
	ctx context.Context
} // NewSimulateOrderService new SimulateOrderService
func NewSimulateOrderService(ctx context.Context) *SimulateOrderService {
	return &SimulateOrderService{ctx: ctx}
}

// Run create note info
func (s *SimulateOrderService) Run(req *aiorder.SimulateOrderReq) (resp *aiorder.SimulateOrderResp, err error) {
	// Finish your business logic.

	return
}
