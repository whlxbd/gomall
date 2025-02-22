package service

import (
	"context"
	"testing"
	rule "github.com/whlxbd/gomall/rpc_gen/kitex_gen/rule"
)

func TestCheckPermission_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckPermissionService(ctx)
	// init req and assert value

	req := &rule.CheckPermissionReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
