package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

func TestVerifyTokenByRPC_Run(t *testing.T) {
	ctx := context.Background()
	s := NewVerifyTokenByRPCService(ctx)
	// init req and assert value

	req := &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjF9.U6p6P6we6Zez-Ak7MtpXSSeAnITU0XTeD6FNBeUWdZQ",
	}
	resp, err := s.Run(req)
	assert.NoError(t, err)
	assert.Equal(t, true, resp.Res)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	req = &auth.VerifyTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsIiR5cCI6IkpXVCJ9.eyJ1c2VyaWQiOjF9.U6p6P6we6Zez-Ak7MtpXSSeAnITU0XTeD6FNBeUWdZQ",
	}
	resp, err = s.Run(req)
	assert.Error(t, err)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
}
