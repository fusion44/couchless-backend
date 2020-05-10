package repositories

import (
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/go-pg/pg/v9"
)

// ActivitiesRepository contains all functions regarding activities
type ActivitiesRepository struct {
	DB *pg.DB
}

// GetActivityByID returns the activity for the given ID
func (r *ActivitiesRepository) GetActivityByID(id string) (*model.Activity, error) {
	var activity *model.Activity
	err := r.DB.Model(&activity).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return activity, nil
}

// GetActivities returns all activities in the database
func (r *ActivitiesRepository) GetActivities() ([]*model.Activity, error) {
	var activities []*model.Activity
	err := r.DB.Model(&activities).Select()

	if err != nil {
		return nil, err
	}

	return activities, nil
}
