package service

import (
	"context"
	"testing"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

func TestUpdate_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateService(ctx)
	// init req and assert value

	req := &auth.UpdateReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
