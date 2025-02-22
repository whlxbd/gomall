package service

import (
	"context"
	"testing"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

func TestGet_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetService(ctx)
	// init req and assert value

	req := &auth.GetReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
