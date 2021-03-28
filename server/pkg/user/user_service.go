package user

type Coords struct {
	X float64
	Y float64
}

type User struct {
	Id string
	Radius int
	Name string
	Coords Coords
}

type SessionItem struct {
	Token string
	User User
}

type Session interface {
	Create(name string, x float64, y float64) (*SessionItem, error)
}

