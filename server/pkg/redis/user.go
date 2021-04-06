package redis

import (
	"fmt"
	"log"
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/seanaye/geog483-final/server/pkg/user"
	"github.com/seanaye/geog483-final/server/pkg/random"
)

func (t *RedisService) CreateUser(name string, radius int, x float64, y float64) (*user.UserItem, error) {
	client := t.getConnection()
	defer client.Close()

	id := random.RandString(8)
	nameErr := client.Set(ctx, fmt.Sprintf("%s_name", id), name, 0).Err()
	if nameErr != nil {
		return nil, nameErr
	}

	loc := &redis.GeoLocation{
		Name: fmt.Sprintf("%s_loc", id),
		Latitude: y,
		Longitude: x,
	}
	locRes, locErr := client.GeoAdd(ctx, "user_locations", loc).Result()
	log.Printf("geoadd: %d", locRes)

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

	return &user.UserItem{
		Id: id,
		Name: name,
		Radius: radius,
		Coords: user.Coords{
			X: x,
			Y: y,
		},
	}, nil
}

func (t *RedisService) DeleteUser(id string) error {
	client := t.getConnection()

	keyStubs := []string{"_name", "_radius", "_loc"}

	var removeKeys []string
	for _, x := range keyStubs {
		removeKeys = append(removeKeys, fmt.Sprintf("%s%s", id, x))
	}

	client.LRem(ctx, "users", 0, id)
	client.ZRem(ctx, "user_locations", removeKeys[2])
	return client.Del(ctx, removeKeys[0], removeKeys[1]).Err()
}



func (t *RedisService) UpdateUserRadius(id string, radius int) (*user.UserItem, error) {
	client := t.getConnection()
	defer client.Close()

	radiusErr := client.Set(ctx, fmt.Sprintf("%s_radius", id), radius, 0).Err()
	if radiusErr != nil {
		return nil, radiusErr
	}

	return t.GetUser(id)
}

func (t *RedisService) UpdateUserName(id string, name string) (*user.UserItem, error) {
	client := t.getConnection()
	defer client.Close()

	nameErr := client.Set(ctx, fmt.Sprintf("%s_name", id), name, 0).Err()

	if nameErr != nil {
		return nil, nameErr
	}

	return t.GetUser(id)
}

func (t *RedisService) UpdateUserLocation(id string, x float64, y float64) (*user.UserItem, error) {
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


func (t *RedisService) GetUser(id string) (*user.UserItem, error) {
	users, err := t.GetUsers(id)

	if err != nil {
		return nil, err
	}

	return users[0], nil
}

func (t *RedisService) GetUsers(ids ...string) ([]*user.UserItem, error) {
	log.Printf("getting users: %s", ids)
	client := t.getConnection()
	defer client.Close()

	pipe := client.TxPipeline()
	var radiusCmds []*redis.StringCmd
	var posCmds []*redis.GeoPosCmd
	var nameCmds []*redis.StringCmd
	for _, id := range ids {
		radiusCmds = append(radiusCmds, pipe.Get(ctx, fmt.Sprintf("%s_radius", id)))

		posCmds = append(posCmds, pipe.GeoPos(ctx, "user_locations", fmt.Sprintf("%s_loc", id)))

		nameCmds = append(nameCmds, pipe.Get(ctx, fmt.Sprintf("%s_name", id)))
	}

	_, err := pipe.Exec(ctx)

	if err != nil {
		return nil, err
	}

	var output []*user.UserItem

	log.Printf("pos cmds: %s", posCmds)
	for i, id := range ids {
		pos := posCmds[i].Val()
		radius, radiusErr := radiusCmds[i].Int()
		name := nameCmds[i].Val()

		if radiusErr != nil {
			return nil, radiusErr
		}

		user := &user.UserItem{
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

func (t *RedisService) GetAllUsers() ([]*user.UserItem, error) {
	client := t.getConnection()
	defer client.Close()

	ids, err := client.LRange(ctx, "users", 0, -1).Result()

	if err != nil {
		return nil, err
	}

	return t.GetUsers(ids...)
}


func (t *RedisService) mapChanUser(sub <-chan *redis.Message, userChan chan<- *user.UserItem, client *redis.Client) {
	for message := range sub {
		user, _ := t.GetUser(message.String())
		if user != nil {
			userChan <- user
		}
	}
	defer client.Close()
}

func (t *RedisService) ListenUsers() (chan *user.UserItem, error) {
	client := t.getConnection()

	sub := client.Subscribe(ctx, "clients")
	channel := sub.Channel()

	out := make(chan *user.UserItem)

	go t.mapChanUser(channel, out, client)

	return out, nil
}

