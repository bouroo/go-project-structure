package infrastructure

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConn struct {
	Context context.Context
	Client  *redis.Client
}

func NewRedisConn(opts redis.Options) *RedisConn {
	if len(opts.Addr) == 0 {
		opts.Addr = "localhost:6379"
	}

	rdb := redis.NewClient(&opts)
	return &RedisConn{Context: context.Background(), Client: rdb}
}
