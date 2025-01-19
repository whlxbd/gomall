package service

import (
	"context"
	"os"
	"testing"

	"github.com/whlxbd/gomall/app/user/biz/dal"
	user "github.com/whlxbd/gomall/rpc_gen/kitex_gen/user"
)

func TestMain(m *testing.M) {
	// 设置工作目录
	if err := os.Chdir("/home/lry/code/go/gomall/app/user"); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestLogin_Run(t *testing.T) {
	dal.Init()
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	// req := &user.LoginReq{}
	// resp, err := s.Run(req)
	// t.Logf("err: %v", err)
	// t.Logf("resp: %v", resp)

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
