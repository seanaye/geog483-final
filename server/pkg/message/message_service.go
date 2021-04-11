package message

import (
	"context"

	"github.com/seanaye/geog483-final/server/pkg/user"
)

type MessageItem struct {
	Content string
	User user.UserItem
}

type Message interface {
	ListenMessages(user *user.UserItem) (<-chan *MessageItem, error, context.CancelFunc)
	CreateMessage(user *user.UserItem, message string) (bool, error)
}
