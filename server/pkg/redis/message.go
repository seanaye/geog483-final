package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/seanaye/geog483-final/server/pkg/message"
	"github.com/seanaye/geog483-final/server/pkg/user"
)

type msg struct {
	Content string
	User *user.UserItem
}

func (t *RedisService) CreateMessage(user *user.UserItem, message string) (bool, error) {

	client := t.getConnection()

	req := redis.GeoRadiusQuery{
		Radius: float64(user.Radius),
	}

	res, err := client.GeoRadius(ctx, "user_locations", user.Coords.X, user.Coords.Y, &req).Result()

	if err != nil {
		return false, err
	}

	// publish to available channels
	for _, loc := range res {
		id := strings.Replace(loc.Name, "_loc", "", -1)
		//publish to the users own receive channel
		m := msg{
			Content: message,
			User: user,
		}

		b, err := json.Marshal(m)
		if err != nil {
			return false, err
		}

		client.Publish(ctx, fmt.Sprintf("%s_rcv_msg", id), b)
	}

	return true, nil
}


func (t *RedisService) ListenMessages(user *user.UserItem) (<-chan *message.MessageItem, error, context.CancelFunc) {
	client := t.getConnection()

	sub := client.Subscribe(ctx, fmt.Sprintf("%s_rcv_msg", user.Id))

	channel := sub.Channel()

	out := make(chan *message.MessageItem)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <-ctx.Done():
				client.Close()
				return
			case msg := <- channel:
				var item message.MessageItem
				json.Unmarshal([]byte(msg.Payload), &item)
				out <- &item
			}
		}
	}()

	return out, nil, cancel
}
