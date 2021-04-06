package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/seanaye/geog483-final/server/graph/generated"
	"github.com/seanaye/geog483-final/server/graph/model"
	"github.com/seanaye/geog483-final/server/pkg/middleware"
	"github.com/seanaye/geog483-final/server/pkg/translate"
)

func (r *mutationResolver) CreateSession(ctx context.Context, input model.SessionInput) (*model.Session, error) {
	session, err := r.Session.CreateSession(input.Name, input.X, input.Y)

	if err != nil {
		return nil, err
	}

	return translate.MakeSession(session), nil
}

func (r *mutationResolver) UpdateRadius(ctx context.Context, radius int) (*model.User, error) {
	user := middleware.ForContext(ctx)

	updated, err := r.User.UpdateUserRadius(user.Id, radius)

	if err != nil {
		return nil, err
	}

	return translate.MakeUser(updated), nil
}

func (r *mutationResolver) UpdateName(ctx context.Context, name string) (*model.User, error) {
	user := middleware.ForContext(ctx)

	updated, err := r.User.UpdateUserName(user.Id, name)

	if err != nil {
		return nil, err
	}

	return translate.MakeUser(updated), nil
}

func (r *mutationResolver) Connect(ctx context.Context, id string) (bool, error) {
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.User.GetAllUsers()

	if err != nil {
		return nil, err
	}

	var output []*model.User

	for _, user := range users {
		output = append(output, translate.MakeUser(user))
	}

	return output, nil
}

func (r *subscriptionResolver) NewUsers(ctx context.Context) (<-chan *model.User, error) {
	userChan, err := r.User.ListenUsers()

	if err != nil {
		return nil, err
	}

	out := make(chan *model.User)

	go func() {
		for user := range userChan {
			out <- translate.MakeUser(user)
		}
	}()

	return out, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
