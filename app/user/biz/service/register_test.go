package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/whlxbd/gomall/app/user/biz/dal"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

func TestRegister_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{}

	// Test case: Passwords do not match
	req.Email = "admin@example.com"
	req.Password = "123456"
	req.ConfirmPassword = "1"
	resp, err := s.Run(req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	t.Logf("resp: %v", resp)

	// Test case: Successful registration
	req.Email = "admin@example.com"
	req.Password = "123456"
	req.ConfirmPassword = "123456"
	resp, err = s.Run(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("resp: %v", resp)

	// Test case: Email already exists
	req.Email = "admin@example.com"
	req.Password = "123456"
	req.ConfirmPassword = "123456"
	resp, err = s.Run(req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	t.Logf("resp: %v", resp)

	// Test case: Invalid email format
	req.Email = "invalid-email"
	req.Password = "123456"
	req.ConfirmPassword = "123456"
	resp, err = s.Run(req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	t.Logf("resp: %v", resp)

	// Test case: Password too short
	req.Email = "admin@example.com"
	req.Password = "123"
	req.ConfirmPassword = "123"
	resp, err = s.Run(req)
	assert.Error(t, err)
	assert.Nil(t, resp)
	t.Logf("resp: %v", resp)

	req.Email = "user@example.com"
	req.Password = "123456"
	req.ConfirmPassword = "123456"
	resp, err = s.Run(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	t.Logf("resp: %v", resp)
}
