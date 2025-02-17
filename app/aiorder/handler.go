package main

import (
	"context"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
	"github.com/whlxbd/gomall/app/aiorder/biz/service"
)

// AIOrderServiceImpl implements the last service interface defined in the IDL.
type AIOrderServiceImpl struct{}

// QueryOrder implements the AIOrderServiceImpl interface.
func (s *AIOrderServiceImpl) QueryOrder(ctx context.Context, req *aiorder.QueryOrderReq) (resp *aiorder.QueryOrderResp, err error) {
	resp, err = service.NewQueryOrderService(ctx).Run(req)

	return resp, err
}

// SimulateOrder implements the AIOrderServiceImpl interface.
func (s *AIOrderServiceImpl) SimulateOrder(ctx context.Context, req *aiorder.SimulateOrderReq) (resp *aiorder.SimulateOrderResp, err error) {
	resp, err = service.NewSimulateOrderService(ctx).Run(req)

	return resp, err
}
