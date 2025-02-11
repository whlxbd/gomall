package service

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 设置工作目录
	if err := os.Chdir("/home/lry/workspace/go/gomall/app/user"); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}
