package redis

import (
	"context"

	"github.com/seanaye/geog483-final/server/pkg/jwt"
	"github.com/seanaye/geog483-final/server/pkg/session"
)

func (t *RedisService) CreateSession(name string, x float64, y float64) (*session.SessionItem, error) {
	client := t.getConnection()
	defer client.Close()

	const defaultRadius = 1000
	user, err := t.CreateUser(name, defaultRadius, x, y)

	if err != nil {
		return nil, err
	}

	token, err := jwt.CreateToken(user.Id, name)
	if err != nil {
		return nil, err
	}


	// publish to channel
	client.Publish(ctx, "client_id", user.Id)

	return &session.SessionItem{
		Token: token,
		User: *user,
	}, nil
}

func (t *RedisService) EndSession(id string) error {
	client := t.getConnection()
	defer client.Close()

	client.Publish(ctx, "deleted_id", id)

	return t.DeleteUser(id)
}

func (t *RedisService) ListenEndedSession() (<-chan string, error, context.CancelFunc) {
	client := t.getConnection()

	sub := client.Subscribe(ctx, "deleted_id")

	channel := sub.Channel()

	out := make(chan string)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			select {
			case <- ctx.Done():
				client.Close()
				return
			case id := <- channel:
				out <- id.Payload
			}
		}
	}()

	return out, nil, cancel
}

