package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
import (
	"github.com/seanaye/geog483-final/server/pkg/message"
	"github.com/seanaye/geog483-final/server/pkg/session"
	"github.com/seanaye/geog483-final/server/pkg/user"
)

type Resolver struct{
	Session session.Session
	User user.User
	Message message.Message
}
