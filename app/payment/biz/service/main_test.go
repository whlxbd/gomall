package service

import (
	"os"
	"testing"

	"github.com/whlxbd/gomall/app/payment/biz/dal"
)

func TestMain(m *testing.M) {
	// 设置工作目录
	if err := os.Chdir("/home/lry/workspace/go/gomall/app/user"); err != nil {
		panic(err)
	}
	dal.Init()
	os.Exit(m.Run())
}
