package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type RedisService struct {
	Host string
}

// flush database on app startup
func (t *RedisService) Clear () {
	client := t.getConnection()
	client.FlushDB(ctx)
}

func (t *RedisService) getConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: t.Host,
	})
}

