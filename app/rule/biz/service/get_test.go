package service

import (
	"context"
	"testing"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

func TestGet_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetService(ctx)
	// init req and assert value

	req := &rule.GetReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
