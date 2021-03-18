package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/seanaye/geog483-final/server/graph/generated"
	"github.com/seanaye/geog483-final/server/graph/model"
)

func (r *mutationResolver) CreateSession(ctx context.Context, name string) (*model.Session, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ChangeRadius(ctx context.Context, token string, radius int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) _(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
