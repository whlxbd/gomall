package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whlxbd/gomall/app/user/biz/dal"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

func TestInfo_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewInfoService(ctx)
	// init req and assert value

	req := &user.InfoReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	assert.Equal(t, resp.UserId, int32(1))
	assert.Equal(t, nil, err)
	assert.Equal(t, "admin@test.com", resp.Email)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
