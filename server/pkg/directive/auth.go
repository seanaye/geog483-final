package directive

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/seanaye/geog483-final/server/pkg/middleware"
)

func Auth(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	user := middleware.ForContext(ctx)

	if user == nil {
		return nil, errors.New("User not found in context")
	}

	return next(ctx)
}
