package cache

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/whlxbd/gomall/app/rule/biz/dal/model/whitelist"
	"github.com/whlxbd/gomall/app/rule/biz/dal/redis"
)

func GetWhiteRouter(ctx context.Context, router string) (resp *whitelist.WhiteRouter, err error) {
	ok := redis.RedisClient.Get(ctx, router)
	if ok.Err() != nil {
		return nil, ok.Err()
	}
	if ok.Val() == "" {
		return nil, kerrors.NewBizStatusError(403, "router not in white list")
	}
	resp = &whitelist.WhiteRouter{
		Router: router,
	}
	return
}

func DeleteWhiteRouter(ctx context.Context, router string) (err error) {
	ok := redis.RedisClient.Del(ctx, router)
	if ok.Err() != nil {
		return ok.Err()
	}
	return
}

func SetWhiteRouter(ctx context.Context, router string) (err error) {
	ok := redis.RedisClient.Set(ctx, router, "1", 0)
	if ok.Err() != nil {
		return ok.Err()
	}
	return
}
