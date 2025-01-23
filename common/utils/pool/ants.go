package pool

import (
    "github.com/panjf2000/ants/v2"
)

var (
	defaultPool *ants.Pool
)

func Init() {
	var err error
	defaultPool, err = ants.NewPool(100000)
	if err != nil {
		panic(err)
	}
}

func Submit(task func()) error {
	return defaultPool.Submit(task)
}

func Release() {
	defaultPool.Release()
}
