package message

import (
	"github.com/go-redis/redis/v8"
	"github.com/seanaye/geog483-final/server/pkg/user"
)

type MessageItem struct {
	Content string
	User user.UserItem
}

type Message interface {
	ListenMessages(user *user.UserItem) (<-chan *MessageItem, *redis.PubSub)
	CreateMessage(user *user.UserItem, message string) (bool, error)
}
