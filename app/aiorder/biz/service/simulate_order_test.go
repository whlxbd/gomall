package service

import (
	"context"
	"testing"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
)

func TestSimulateOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSimulateOrderService(ctx)
	// init req and assert value

	req := &aiorder.SimulateOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
