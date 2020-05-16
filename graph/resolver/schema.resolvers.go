package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	loader "github.com/fusion44/ll-backend/db/loaders"
	"github.com/fusion44/ll-backend/domain"
	"github.com/fusion44/ll-backend/graph/generated"
	"github.com/fusion44/ll-backend/graph/model"
)

func (r *activityResolver) User(ctx context.Context, obj *model.Activity) (*model.User, error) {
	return loader.GetUserLoader(ctx).Load(obj.UserID)
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, domain.ErrInvalidInput
	}

	return r.Domain.Register(ctx, input)
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)

	if !isValid {
		return nil, domain.ErrInvalidInput
	}

	return r.Domain.Login(ctx, input)
}

func (r *mutationResolver) AddActivity(ctx context.Context, input model.NewActivity) (*model.Activity, error) {
	return r.Domain.AddActivity(ctx, input)
}

func (r *mutationResolver) UpdateActivity(ctx context.Context, input model.UpdateActivity) (*model.Activity, error) {
	return r.Domain.UpdateActivity(ctx, input)
}

func (r *mutationResolver) DeleteActivity(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteActivity(ctx, id)
}

func (r *queryResolver) Activity(ctx context.Context, id string) (*model.Activity, error) {
	return r.Domain.GetActivityByID(ctx, id)
}

func (r *queryResolver) Activities(ctx context.Context, filter *model.ActivityFilter, limit *int, offset *int) ([]*model.Activity, error) {
	return r.Domain.GetActivities(ctx, filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.Domain.GetUserByID(ctx, id)
}

// Activity returns generated.ActivityResolver implementation.
func (r *Resolver) Activity() generated.ActivityResolver { return &activityResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type activityResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
