package translate

import (
	"github.com/seanaye/geog483-final/server/pkg/user"
	"github.com/seanaye/geog483-final/server/pkg/session"
	"github.com/seanaye/geog483-final/server/graph/model"
)

func MakeUser(u *user.UserItem) *model.User {
	return &model.User{
		ID: u.Id,
		Name: u.Name,
		Radius: u.Radius,
		Coords: &model.Coords{
			X: u.Coords.X,
			Y: u.Coords.Y,
		},
	}
}

func MakeSession(s *session.SessionItem) *model.Session {
	return &model.Session{
		Token: s.Token,
		User: MakeUser(&s.User),
	}
}

