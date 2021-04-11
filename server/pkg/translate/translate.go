package translate

import (
	"github.com/seanaye/geog483-final/server/graph/model"
	"github.com/seanaye/geog483-final/server/pkg/message"
	"github.com/seanaye/geog483-final/server/pkg/session"
	"github.com/seanaye/geog483-final/server/pkg/user"
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

func MakeMessage(s *message.MessageItem) *model.Message {
	return &model.Message{
		Content: s.Content,
		User: MakeUser(&s.User),
	}
}

