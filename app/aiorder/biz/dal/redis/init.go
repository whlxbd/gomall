package redis

import (
	"context"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
	"github.com/whlxbd/gomall/common/utils/pool"
	"time"
	"github.com/cloudwego/kitex/pkg/klog"

)

var RedisClient *redis.Client

func Init() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USER"),
		Password: os.Getenv("REDIS_PSWD"),
		DB:       db,
		Protocol: 2,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	// 设置内存策略为allkeys-lru
    if err := RedisClient.ConfigSet(ctx, "maxmemory-policy", "allkeys-lru").Err(); err != nil {
        panic(err)
    }

	// 配置RDB持久化
    if err := RedisClient.ConfigSet(ctx, "save", "900 1 300 10 60 10000").Err(); err != nil {
        panic(err)
    }

    // 启用AOF持久化
    if err := RedisClient.ConfigSet(ctx, "appendonly", "yes").Err(); err != nil {
        panic(err)
    }
    
    // 设置AOF持久化策略
    if err := RedisClient.ConfigSet(ctx, "appendfsync", "everysec").Err(); err != nil {
        panic(err)
    }

	_ = pool.Submit(func() {
		ticker := time.NewTicker(15 * time.Minute)
        for range ticker.C {
            if err := RedisClient.Save(ctx).Err(); err != nil {
                // 记录错误但不中断服务
                klog.Errorf("Redis save failed: %v", err)
            }
        }
	})
}
