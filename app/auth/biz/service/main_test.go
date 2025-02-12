package service

import (
	"os"
	"testing"

	"github.com/whlxbd/gomall/app/auth/biz/dal"
	"github.com/whlxbd/gomall/app/auth/infra/rpc"
)

func TestMain(m *testing.M) {
	// 设置工作目录
	if err := os.Chdir("/home/lry/workspace/go/gomall/app/user"); err != nil {
		panic(err)
	}
	dal.Init()
	os.Setenv("JWT_SECRET", "test")
	rpc.InitClient()
	os.Exit(m.Run())
}
