package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"micro_tiktok/pkg/constants"
	"time"
)

var RDB *redis.Client

func Init() {
	RDB = redis.NewClient(&redis.Options{
		Addr: constants.RedisAddress,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
}
