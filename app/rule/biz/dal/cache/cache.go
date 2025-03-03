package cache

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/redis/go-redis/v9"
	"github.com/whlxbd/gomall/app/rule/biz/dal/model/whitelist"
)

func GetWhiteRouter(ctx context.Context, rdb *redis.Client, router string) (resp *whitelist.WhiteRouter, err error) {
	ok := rdb.Get(ctx, router)
	if ok.Err() != nil {
		return &whitelist.WhiteRouter{
			Router: "",
		}, nil
	}
	if ok.Val() == "" {
		return nil, kerrors.NewBizStatusError(403, "router not in white list")
	}
	resp = &whitelist.WhiteRouter{
		Router: router,
	}
	return
}

func DeleteWhiteRouter(ctx context.Context, rdb *redis.Client, router string) (err error) {
	ok := rdb.Del(ctx, router)
	if ok.Err() != nil {
		return ok.Err()
	}
	return
}

func SetWhiteRouter(ctx context.Context, rdb *redis.Client, router string) (err error) {
	ok := rdb.Set(ctx, router, "1", 0)
	if ok.Err() != nil {
		return ok.Err()
	}
	return
}
