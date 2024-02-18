package infrastructure

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisConn struct {
	ctx    context.Context
	Client *redis.Client
}

type RedisOptions struct {
	redis.Options
}

func (config *RedisOptions) ApplyDefault() *RedisOptions {
	if len(config.Addr) == 0 {
		config.Addr = "localhost:6379"
	}
	return config
}

func NewRedisConn(opts RedisOptions) *RedisConn {
	opts.ApplyDefault()

	rdb := redis.NewClient(&opts.Options)
	return &RedisConn{ctx: context.Background(), Client: rdb}
}
