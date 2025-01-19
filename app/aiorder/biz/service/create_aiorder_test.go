package service

import (
	"context"
	"testing"
	aiorder "github.com/whlxbd/gomall/rpc_gen/kitex_gen/aiorder"
)

func TestCreateAIOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateAIOrderService(ctx)
	// init req and assert value

	req := &aiorder.CreateAIOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
