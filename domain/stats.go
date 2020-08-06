package domain

import (
	"context"

	"github.com/fusion44/ll-backend/graph/model"
	"github.com/fusion44/ll-backend/middleware"
)

// UpdateStatsForCurrentUser updates all statistics for the current user
func (d *Domain) UpdateStatsForCurrentUser(ctx context.Context) ([]*model.UserStatMonth, error) {
	logger, _ := middleware.GetLoggerFromContext(ctx)
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	// Get updated statistics for the current user
	stats, err := d.StatsRepo.CalculateStatsForUser(u.ID)
	if err != nil {
		logger.Errorf("Error calculating stats for user %s: %s", u.ID, err)
		return nil, err
	}

	// Cache the updated statistics to the DB
	stats, err = d.StatsRepo.InsertOrUpdateStatsForUser(stats, u.ID)
	if err != nil {
		logger.Errorf("Error inserting stats for user %s: %s", u.ID, err)
		return nil, err
	}

	return stats, nil
}

// GetStatsForCurrentUser retrieves statistical activity data for a user
func (d *Domain) GetStatsForCurrentUser(ctx context.Context) ([]*model.UserStatMonth, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	stats, err := d.StatsRepo.GetStatsForUserFromDB(u.ID)

	if err != nil {
		return nil, err
	}

	return stats, nil
}
