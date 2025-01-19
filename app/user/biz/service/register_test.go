package service

import (
	"context"
	"github.com/whlxbd/gomall/app/user/biz/dal"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestRegister_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterReq{}

	req.Email = "admin@example.com"
	req.Password = "123456"
	req.ConfirmPassword = "1"
	resp, err := s.Run(req)
	if err == nil {
		t.Errorf("Password not match, expect error")
	} else {
		t.Logf("err: %v", err)
	}
	t.Logf("resp: %v", resp)

	req.Email = "admin@example.com"
	req.Password = "123456"
	req.ConfirmPassword = "123456"
	resp, err = s.Run(req)
	if err != nil {
		t.Errorf("Create user failed, err: %v", err)
	}
	t.Logf("resp: %v", resp)

	req.Email = "admin@example.com"
	req.Password = "123456"
	req.ConfirmPassword = "123456"
	resp, err = s.Run(req)
	if err == nil {
		t.Errorf("Email confirm: %v", err)
	} else {
		t.Logf("")
	}
	t.Logf("resp: %v", resp)
}
