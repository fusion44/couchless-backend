package domain

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/fusion44/ll-backend/graph/model"
	"github.com/fusion44/ll-backend/middleware"
)

// GetActivities returns the filtered activities for the current user
func (d *Domain) GetActivities(ctx context.Context, filter *model.ActivityFilter, limit, offset *int) ([]*model.Activity, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	activities, err := d.ActivityRepo.GetActivities(u.ID, filter, limit, offset)

	if err != nil {
		log.Printf("Error fetching activities: %v", err)
		return nil, ErrInternalServer
	}

	// no further authentication since we only fetch activities for the current user

	return activities, nil
}

// GetActivityByID returns an existing activity
func (d *Domain) GetActivityByID(ctx context.Context, id string) (*model.Activity, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	a, err := d.ActivityRepo.GetActivityByID(id)

	if err != nil {
		return nil, err
	}

	if !a.IsOwner(u) {
		return nil, ErrUnauthorized
	}

	return a, nil
}

// AddActivity adds a new Activity to the database for current user
func (d *Domain) AddActivity(ctx context.Context, input model.NewActivity) (*model.Activity, error) {
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
	return d.ActivityRepo.AddActivity(&activity)
}

// UpdateActivity updates an Activity if the user is logged is
func (d *Domain) UpdateActivity(ctx context.Context, input model.UpdateActivity) (*model.Activity, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	activity, err := d.ActivityRepo.GetActivityByID(input.ID)
	if err != nil || activity == nil {
		return nil, errors.New("Activity does not exists")
	}

	if !activity.IsOwner(u) {
		return nil, ErrUnauthorized
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

	activity, err = d.ActivityRepo.UpdateActivity(activity)

	if err != nil {
		return nil, fmt.Errorf("Error while updating activity: %v", err)
	}

	return activity, nil
}

// DeleteActivity deletes an Activity if the user is logged is
func (d *Domain) DeleteActivity(ctx context.Context, id string) (bool, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return false, ErrUnauthenticated
	}

	activity, err := d.ActivityRepo.GetActivityByID(id)
	if err != nil || activity == nil {
		return false, errors.New("Activity does not exist")
	}

	if !activity.IsOwner(u) {
		return false, ErrUnauthorized
	}

	err = d.ActivityRepo.DeleteActivity(activity)

	if err != nil {
		return false, fmt.Errorf("Error deleting activity: %v", err)
	}

	return true, nil
}