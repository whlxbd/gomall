package service

import (
	"context"
	"github.com/whlxbd/gomall/app/user/biz/dal"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
	"testing"
)

func TestLogin_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	//req := &user.LoginReq{}
	//resp, err := s.Run(req)
	//t.Logf("err: %v", err)
	//t.Logf("resp: %v", resp)

	req := &user.LoginReq{
		Email:    "admin@example.com",
		Password: "123456",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)
	if err != nil {
		t.Error(err)
	}

}
