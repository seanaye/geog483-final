package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/seanaye/geog483-final/server/graph/generated"
	"github.com/seanaye/geog483-final/server/graph/model"
)

func (r *mutationResolver) CreateSession(ctx context.Context, input model.SessionInput) (*model.Session, error) {
	session, err := r.Session.Create(input.Name, input.X, input.Y)

	if err != nil {
		return nil, err
	}

	return &model.Session{
		Token: session.Token,
		User: &model.User{
			Name: session.User.Name,
			Radius: session.User.Radius,
			Coords: &model.Coords{
				X: session.User.Coords.X,
				Y: session.User.Coords.Y,
			},
		},
	}, nil
}

func (r *mutationResolver) ChangeRadius(ctx context.Context, token string, radius int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) _(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}
