package service

import (
	"context"
	order "github.com/whlxbd/gomall/rpc_gen/kitex_gen/order"
)

type EditOrderService struct {
	ctx context.Context
} // NewEditOrderService new EditOrderService
func NewEditOrderService(ctx context.Context) *EditOrderService {
	return &EditOrderService{ctx: ctx}
}

// Run create note info
func (s *EditOrderService) Run(req *order.EditOrderReq) (resp *order.EditOrderResp, err error) {
	// Finish your business logic.
	
	return
}
