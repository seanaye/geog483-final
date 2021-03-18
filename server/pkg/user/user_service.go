package user

import "github.com/seanaye/geog483-final/server/graph/model"

type User interface {
	CreateSession(name string)
}
