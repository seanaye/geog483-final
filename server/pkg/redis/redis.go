package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisService struct {
	Host string
}


func (t *RedisService) getConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: t.Host,
	})
}
