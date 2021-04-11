package user

import "context"

type Coords struct {
	X float64
	Y float64
}

type UserItem struct {
	Id string
	Radius int
	Name string
	Coords Coords
}

type User interface {
	UpdateUserRadius(id string, radius int) (*UserItem, error)
	UpdateUserLocation(user *UserItem, x float64, y float64) (*UserItem, error)
	UpdateUserName(id string, name string) (*UserItem, error)
	CreateUser(name string, radius int, x float64, y float64) (*UserItem, error)
	DeleteUser(id string) error
	GetUser(id string) (*UserItem, error)
	GetUsers(ids ...string) ([]*UserItem, error)
	GetAllUsers() ([]*UserItem, error)
	ListenUsers() (chan *UserItem, error, context.CancelFunc)
}


