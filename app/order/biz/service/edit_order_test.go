package service

import (
	"context"
	"testing"
	order "github.com/whlxbd/gomall/rpc_gen/kitex_gen/order"
)

func TestEditOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewEditOrderService(ctx)
	// init req and assert value

	req := &order.EditOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
