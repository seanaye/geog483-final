package redis 

import (
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/seanaye/geog483-final/server/pkg/jwt"
	"github.com/seanaye/geog483-final/server/pkg/random"
	"github.com/seanaye/geog483-final/server/pkg/user"
)

func (t *RedisService) Create(name string, x float64, y float64) (*user.SessionItem, error) {
	client := t.getConnection()
	defer client.Close()
	id := random.RandString(8)
	token, err := jwt.CreateToken(id, name)
	if err != nil {
		return nil, err
	}

	nameErr := client.Set(ctx, fmt.Sprintf("%s_name", id), name, 0).Err()
	if nameErr != nil {
		return nil, nameErr
	}

	location_name := fmt.Sprintf("%s_loc", id)
	loc := &redis.GeoLocation{
		Name: location_name,
		Latitude: y,
		Longitude: x,
	}
	locRes, locErr := client.GeoAdd(ctx, location_name, loc).Result()
	if locRes < 1 || locErr != nil {
		return nil, errors.New("Error adding location to database")
	}

	defaultRadius := 1000
	radiusErr := client.Set(ctx, fmt.Sprintf("%s_radius", id), defaultRadius, 0).Err()
	if radiusErr != nil {
		return nil, radiusErr
	}

	// add to list of users
	client.RPush(ctx, "users", id)

	// publish to channel
	client.Publish(ctx, "clients", fmt.Sprintf("%s_enter", id))

	return &user.SessionItem{
		Token: token,
		User: user.User{
			Id: id,
			Name: name,
			Coords: user.Coords{
				X: x,
				Y: y,
			},
			Radius: defaultRadius,
		},
	}, nil
}

func (t *RedisService) End(id string) error {
	client := t.getConnection()
	defer client.Close()

	keyStubs := []string{"_name", "_radius", "_loc"}

	var removeKeys []string
	for _, x := range keyStubs {
		removeKeys = append(removeKeys, fmt.Sprintf("%s%s", id, x))
	}

	client.LRem(ctx, "users", 0, id)
	client.Publish(ctx, "clients", fmt.Sprintf("%s_exit", id))

	return client.Del(ctx, removeKeys...).Err()
}

