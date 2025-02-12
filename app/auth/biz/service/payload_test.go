package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	auth "github.com/whlxbd/gomall/rpc_gen/kitex_gen/auth"
)

func TestPayload_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPayloadService(ctx)
	// init req and assert value

	req := &auth.PayloadReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MTczOTQ1Mzg3MywidHlwZSI6ImFkbWluIiwidXNlcl9pZCI6MX0.8DNwRg6PI0DMTM4duv5csnf90_k98okKCLyuuyaoHK4",
	}

	resp, err := s.Run(req)
	assert.NoError(t, err)
	assert.Equal(t, int32(1), resp.UserId)
	assert.Equal(t, "admin", resp.Type)

}
