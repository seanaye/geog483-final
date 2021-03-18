package redis

import (
	"github.com/go-redis/redis/v8"
)

type UserService struct {
	Host string
}

func (t *UserService) Initialize() error {
	rdb := redis.NewClient(&redis.Options{
		Addr: redis_addr,
	})
}

func (t *UserService) Connect() {

}

func (t *UserService) getConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: t.Host
	})
}
