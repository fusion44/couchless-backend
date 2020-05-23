package repositories

import (
	"fmt"

	"github.com/fusion44/ll-backend/graph/model"
	"github.com/go-pg/pg/v9"
)

// ActivitiesRepository contains all functions regarding activities
type ActivitiesRepository struct {
	DB *pg.DB
}

// GetActivitiesByField returns the user for the given field
func (r *ActivitiesRepository) GetActivitiesByField(field, value string) ([]*model.Activity, error) {
	var activity []*model.Activity
	err := r.DB.Model(&activity).Where(fmt.Sprintf("%v = ?", field), value).Select()
	return activity, err
}

// GetActivityByID returns the activity for the given ID
func (r *ActivitiesRepository) GetActivityByID(id string) (*model.Activity, error) {
	var activity model.Activity
	err := r.DB.Model(&activity).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &activity, nil
}

// GetActivitiesByFileID gets an activity for the given file ID
func (r *ActivitiesRepository) GetActivitiesByFileID(fileID string) ([]*model.Activity, error) {
	var activities []*model.Activity
	err := r.DB.Model(&activities).Where("file_id = ?", fileID).Select()

	if err != nil {
		return nil, err
	}

	return activities, nil
}

// DeleteActivity deletes the given activity
func (r *ActivitiesRepository) DeleteActivity(activity *model.Activity) error {
	_, err := r.DB.Model(activity).Where("id = ?", activity.ID).Delete()
	return err
}

// GetActivities returns all activities in the database
func (r *ActivitiesRepository) GetActivities(userID string, filter *model.ActivityFilter, limit, offset *int) ([]*model.Activity, error) {
	var activities []*model.Activity

	query := r.DB.Model(&activities).Where("user_id = ?", userID).Order("id")

	if filter != nil {
		if filter.StartTime != nil {
			query.Where("start_time >= ?", filter.StartTime)
		}

		if filter.EndTime != nil {
			query.Where("end_time <= ?", filter.EndTime)
		}

		if filter.Comment != nil && *filter.Comment != "" {
			query.Where("comment ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Comment))
		}

		if filter.SportType != nil && *filter.SportType != "" {
			query.Where("sport_type = ?", *filter.SportType)
		}
	}

	if limit != nil {
		query.Limit(*limit)
	}

	if offset != nil {
		query.Offset(*offset)
	}
	err := query.Select()

	if err != nil {
		return nil, err
	}

	return activities, nil
}

// AddActivity inserts an activity into the database
func (r *ActivitiesRepository) AddActivity(activity *model.Activity) (*model.Activity, error) {
	_, err := r.DB.Model(activity).Returning("*").Insert()
	return activity, err
}

// UpdateActivity updates an activity in the database
func (r *ActivitiesRepository) UpdateActivity(activity *model.Activity) (*model.Activity, error) {
	_, err := r.DB.Model(activity).Where("id = ?", activity.ID).Returning("*").Update()
	return activity, err
}
