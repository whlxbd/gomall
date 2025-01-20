package redis

import (
	"fmt"
	"os"
	"strconv"

	"context"

	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)

func Init() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s", os.Getenv("REDIS_ADDR")),
		Username: fmt.Sprintf("%s", os.Getenv("REDIS_USER")),
		Password: fmt.Sprintf("%s", os.Getenv("REDIS_PSWD")),
		DB:       db,
	})
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}
