package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/seanaye/geog483-final/server/pkg/random"
	"github.com/seanaye/geog483-final/server/pkg/jwt"
	"github.com/seanaye/geog483-final/server/pkg/user"
)

var ctx = context.Background()

type RedisService struct {
	Host string
}

func (t *RedisService) Create(name string, x float64, y float64) (*user.SessionItem, error) {
	client := t.getConnection()
	id := random.RandString(8)
	token, err := jwt.CreateToken(id, name)
	if err != nil {
		return nil, err
	}

	nameErr := client.Set(ctx, fmt.Sprintf("%s_name", id), name, 0).Err()
	if nameErr != nil {
		return nil, nameErr
	}
	xErr := client.Set(ctx, fmt.Sprintf("%s_X", id), x, 0).Err()
	if xErr != nil {
		return nil, xErr
	}

	yErr := client.Set(ctx, fmt.Sprintf("%s_Y", id), y, 0).Err()
	if err != nil {
		return nil, yErr
	}

	defaultRadius := 1000
	radiusErr := client.Set(ctx, fmt.Sprintf("%s_radius", id), defaultRadius, 0).Err()
	if radiusErr != nil {
		return nil, radiusErr
	}

	return &user.SessionItem{
		Token: token,
		User: user.User{
			Name: name,
			Coords: user.Coords{
				X: x,
				Y: y,
			},
			Radius: defaultRadius,
		},
	}, nil
}

func (t *RedisService) getConnection() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: t.Host,
	})
}
