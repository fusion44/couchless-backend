package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"

	loader "github.com/fusion44/ll-backend/db/loaders"
	"github.com/fusion44/ll-backend/graph/generated"
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/fusion44/ll-backend/middleware"

	gcontext "github.com/fusion44/ll-backend/context"
)

// Authentication errors
var (
	ErrBadCredentials  = errors.New("Login credentials not valid")
	ErrUnauthenticated = errors.New("Unauthenticated")
)

func (r *activityResolver) User(ctx context.Context, obj *model.Activity) (*model.User, error) {
	return loader.GetUserLoader(ctx).Load(obj.UserID)
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := r.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("Email is already in use")
	}

	_, err = r.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("Username is already in use")
	}

	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
	}

	errMsg := errors.New("Something went wrong")

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("Password hashing failed")
		return nil, errMsg
	}

	tx, err := r.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		return nil, errMsg
	}

	defer tx.Rollback()

	if _, err := r.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err // tx.Rollback is called
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error while commiting: %v", err)
		return nil, err // tx.Rollback is called
	}

	cfg, err := gcontext.GetConfigFromContext(ctx)
	if err != nil {
		log.Printf("No config in context: %v", err)
		return nil, err
	}

	token, err := user.GenToken(cfg)
	if err != nil {
		log.Printf("Error while generating token: %v", err)
		return nil, err
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil

	// Rollback is called, but on an empty transaction,
	// as it was committed above
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	if input.Email == nil && input.Username == nil {
		return nil, errors.New("Login failed, no email or username provided")
	}

	var user *model.User
	var err error

	if input.Email != nil && len(*input.Email) > 3 {
		user, err = r.UsersRepo.GetUserByEmail(*input.Email)
	}
	if user == nil && input.Username != nil && *input.Username != "" {
		user, err = r.UsersRepo.GetUserByUsername(*input.Username)
	}

	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	cfg, err := gcontext.GetConfigFromContext(ctx)
	if err != nil {
		log.Printf("No config in context: %v", err)
		return nil, err
	}

	token, err := user.GenToken(cfg)
	if err != nil {
		return nil, ErrBadCredentials
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) AddActivity(ctx context.Context, input model.NewActivity) (*model.Activity, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	// TODO: Add checks
	activity := model.Activity{
		Comment:   *input.Comment,
		StartTime: input.StartTime,
		EndTime:   input.EndTime,
		SportType: input.SportType,
		UserID:    u.ID,
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

func (r *mutationResolver) DeleteActivity(ctx context.Context, id string) (bool, error) {
	activity, err := r.ActivityRepo.GetActivityByID(id)
	if err != nil || activity == nil {
		return false, errors.New("Activity does not exists")
	}

	err = r.ActivityRepo.DeleteActivity(activity)

	if err != nil {
		return false, fmt.Errorf("Error deleting activity: %v", err)
	}

	return true, nil
}

func (r *queryResolver) Activity(ctx context.Context, id string) (*model.Activity, error) {
	return r.ActivityRepo.GetActivityByID(id)
}

func (r *queryResolver) Activities(ctx context.Context, filter *model.ActivityFilter, limit *int, offset *int) ([]*model.Activity, error) {
	return r.ActivityRepo.GetActivities(filter, limit, offset)
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
