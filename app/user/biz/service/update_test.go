package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whlxbd/gomall/app/user/biz/dal"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

func TestUpdate_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewUpdateService(ctx)

	// Test case: Successful update
	req := &user.UpdateReq{
		UserId:    1,
		Email:     "newemail@example.com",
		Username:  "newusername",
		AvatarUrl: "newavatarurl",
		Type:      "admin",
	}
	resp, err := s.Run(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.True(t, resp.Success)

	// Test case: User not found
	req = &user.UpdateReq{
		UserId:    999, // Assuming this user ID does not exist
		Email:     "nonexistent@example.com",
		Username:  "nonexistent",
		AvatarUrl: "nonexistenturl",
		Type:      "user",
	}
	resp, err = s.Run(req)
	assert.Error(t, err)
	assert.NotNil(t, resp)
	assert.False(t, resp.Success)

}
