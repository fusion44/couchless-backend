package repositories

import (
	"fmt"

	"github.com/fusion44/couchless-backend/graph/model"
	"github.com/go-pg/pg/v9"
)

// StatsRepository contains all functions regarding statistics
type StatsRepository struct {
	DB *pg.DB
}

// CalculateStatsForUser calculates the current statistics with a sql query
// This can be computational intensive and results should be cached to the DB
func (r *StatsRepository) CalculateStatsForUser(userID string) ([]*model.UserStatMonth, error) {
	q := `
SELECT to_char(date_trunc('month', start_time), 'YYYY-MM-DD') AS period, sum(duration) AS total, sport_type
FROM activities
WHERE user_id = ?
GROUP BY period, sport_type
ORDER BY period
	`

	var stats []*model.UserStatMonth
	_, err := r.DB.Model(&stats).Query(&stats, q, userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return stats, nil
}

// InsertOrUpdateStatsForUser inserts or updates statistical activity data for a given user id
func (r *StatsRepository) InsertOrUpdateStatsForUser(stats []*model.UserStatMonth, userID string) ([]*model.UserStatMonth, error) {
	err := r.DB.RunInTransaction(func(Tx *pg.Tx) error {
		for _, stat := range stats {
			stat.UserID = userID
			_, err := r.DB.Model(stat).OnConflict("(period, sport_type) DO UPDATE").Insert()
			if err != nil {
				return err
			}
		}

		return nil
	})

	return stats, err
}

// GetStatsForUserFromDB returns statistical activity data for user ID
// These are cached results. Use CalculateStatsForUser id activities where added, updated or deleted
func (r *StatsRepository) GetStatsForUserFromDB(userID string) ([]*model.UserStatMonth, error) {
	var stats []*model.UserStatMonth
	err := r.DB.Model(&stats).Returning("*").Where("user_id = ?", userID).Order("period DESC").Select()

	if err != nil {
		return nil, err
	}

	return stats, err
}
