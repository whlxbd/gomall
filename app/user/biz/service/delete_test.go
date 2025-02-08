package service

import (
	"context"
	"testing"

	"github.com/whlxbd/gomall/app/user/biz/dal"
	"github.com/stretchr/testify/assert"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

func TestDelete_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewDeleteService(ctx)
	
	req := &user.DeleteReq{
		UserId: 1,
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
