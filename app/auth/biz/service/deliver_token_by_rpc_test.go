package service

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

func TestDeliverTokenByRPC_Run(t *testing.T) {
	os.Setenv("JWT_SECRET", "test")
	ctx := context.Background()
	s := NewDeliverTokenByRPCService(ctx)
	// init req and assert value


	req := &auth.DeliverTokenReq{
		UserId: 1,
	}
	expectedToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjF9.U6p6P6we6Zez-Ak7MtpXSSeAnITU0XTeD6FNBeUWdZQ"
	resp, err := s.Run(req)
	assert.NoError(t, err)
	assert.Equal(t, expectedToken, resp.Token)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
