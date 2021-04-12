package session

import (
	"github.com/go-redis/redis/v8"
	"github.com/seanaye/geog483-final/server/pkg/user"
)

type SessionItem struct {
	Token string
	User user.UserItem
}

type Session interface {
	CreateSession(name string, x float64, y float64) (*SessionItem, error)
	EndSession(id string) error
	ListenEndedSession() (<-chan string, *redis.PubSub)
}
