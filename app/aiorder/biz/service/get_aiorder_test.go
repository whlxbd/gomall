package service

import (
	"context"
	"testing"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
)

func TestGetAIOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetAIOrderService(ctx)
	// init req and assert value

	req := &aiorder.GetAIOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
