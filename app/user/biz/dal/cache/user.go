package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ServiceName  = "user"
	FunctionName = "info"
)

func GetByID(rdb *redis.Client, ctx context.Context, id int32) (string, error) {
	key := fmt.Sprintf("%s_%s_%d", ServiceName, FunctionName, id)
	return rdb.Get(ctx, key).Result()
}

func SetByID(rdb *redis.Client, ctx context.Context, id int32, value string) error {
	key := fmt.Sprintf("%s_%s_%d", ServiceName, FunctionName, id)
	return rdb.Set(ctx, key, value, time.Minute).Err()
}
