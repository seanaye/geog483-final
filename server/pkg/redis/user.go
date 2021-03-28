package redis

import (
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/seanaye/geog483-final/server/pkg/user"
)


func (t *RedisService) UpdateRadius(id string, radius int) (*user.User, error) {
	client := t.getConnection()
	defer client.Close()

	radiusErr := client.Set(ctx, fmt.Sprintf("%s_radius", id), radius, 0).Err()
	if radiusErr != nil {
		return nil, radiusErr
	}

	return t.GetUser(id)
}

func (t *RedisService) UpdateLocation(id string, x float64, y float64) (*user.User, error) {
	client := t.getConnection()
	defer client.Close()

	location_name := fmt.Sprintf("%s_loc", id)
	loc := &redis.GeoLocation{
		Name: location_name,
		Longitude: x,
		Latitude: y,
	}

	client.Publish(ctx, "clients", fmt.Sprintf("%s_update", id))

	locErr := client.GeoAdd(ctx, location_name, loc).Err()
	if locErr != nil {
		return nil, locErr
	}

	return t.GetUser(id)
}


func (t *RedisService) GetUser(id string) (*user.User, error) {
	users, err := t.GetUsers(id)

	if err != nil {
		return nil, err
	}

	return users[0], nil
}

func (t *RedisService) GetUsers(ids ...string) ([]*user.User, error) {
	client := t.getConnection()
	defer client.Close()

	pipe := client.TxPipeline()
	var radiusCmds []*redis.StringCmd
	var posCmds []*redis.GeoPosCmd
	var nameCmds []*redis.StringCmd
	for _, id := range ids {
		radiusCmds = append(radiusCmds, pipe.Get(ctx, fmt.Sprintf("%s_radius", id)))

		posCmds = append(posCmds, pipe.GeoPos(ctx, fmt.Sprintf("%s_loc")))

		nameCmds = append(nameCmds, pipe.Get(ctx, fmt.Sprintf("%s_name")))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, err
	}

	var output []*user.User

	for i, id := range ids {
		pos := posCmds[i].Val()
		radius, radiusErr := radiusCmds[i].Int()
		name := nameCmds[i].Val()

		if radiusErr != nil {
			return nil, radiusErr
		}

		user := &user.User{
			Id: id,
			Name: name,
			Radius: radius,
			Coords: user.Coords{
				X: pos[0].Longitude,
				Y: pos[0].Latitude,
			},
		}

		output = append(output, user)
	}

	return output, nil
}

func (t *RedisService) GetAllUsers() ([]*user.User, error) {
	client := t.getConnection()
	defer client.Close()

	ids, err := client.LRange(ctx, "users", 0, -1).Result()

	if err != nil {
		return nil, err
	}

	return t.GetUsers(ids...)
}
