package service

import (
	"context"
	"testing"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

func TestCheckWhite_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckWhiteService(ctx)
	// init req and assert value

	req := &auth.CheckWhiteReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
