package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	loader "github.com/fusion44/ll-backend/db/loaders"
	"github.com/fusion44/ll-backend/graph/generated"
	"github.com/fusion44/ll-backend/graph/model"
)

func (r *activityResolver) User(ctx context.Context, obj *model.Activity) (*model.User, error) {
	return loader.GetUserLoader(ctx).Load(obj.UserID)
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

func (r *mutationResolver) UpdateActivity(ctx context.Context, input model.UpdateActivity) (*model.Activity, error) {
	activity, err := r.ActivityRepo.GetActivityByID(input.ID)
	if err != nil || activity == nil {
		return nil, errors.New("Activity does not exists")
	}

	didUpdate := false

	if input.Comment != nil {
		if len(*input.Comment) < 2 {
			return nil, errors.New("Comment should be at least two characters")
		}
		activity.Comment = *input.Comment
		didUpdate = true
	}

	if input.StartTime != nil {
		activity.StartTime = *input.StartTime
		didUpdate = true
	}

	if input.EndTime != nil {
		activity.EndTime = *input.EndTime
		didUpdate = true
	}
	if input.SportType != nil {
		activity.SportType = *input.SportType
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("Nothing to update")
	}

	activity, err = r.ActivityRepo.UpdateActivity(activity)

	if err != nil {
		return nil, fmt.Errorf("Error while updating activity: %v", err)
	}

	return activity, nil
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
