package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/fusion44/ll-backend/graph/generated"
	"github.com/fusion44/ll-backend/graph/model"
)

func (r *activityResolver) User(ctx context.Context, obj *model.Activity) (*model.User, error) {
	return r.UsersRepo.GetUserByID(obj.UserID)
}

func (r *mutationResolver) AddActivity(ctx context.Context, input model.NewActivity) (*model.Activity, error) {
	// TODO: Add checks
	activity := model.Activity{
		Comment:   *input.Comment,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		SportType: input.SportType,
		UserID:    "1",
	}
	return r.ActivityRepo.AddActivity(&activity)
}

func (r *queryResolver) Activity(ctx context.Context, id string) (*model.Activity, error) {
	return r.ActivityRepo.GetActivityByID(id)
}

func (r *queryResolver) Activities(ctx context.Context) ([]*model.Activity, error) {
	return r.ActivityRepo.GetActivities()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UsersRepo.GetUserByID(id)
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.UsersRepo.GetUsers()
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
