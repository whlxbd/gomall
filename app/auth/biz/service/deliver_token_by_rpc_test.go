package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.DeliverTokenReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.Token)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	t.Logf("token: %v", resp.Token)
}
