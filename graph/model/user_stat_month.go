package model

import "time"

// UserStatMonth describes accumulated monthly statistics for activities by sport type
type UserStatMonth struct {
	// First day of the month of this stat
	Period time.Time `json:"period"`
	// Total time spent with this sport in seconds
	Total int `json:"total"`
	// The sport type
	SportType string `json:"sportType"`
	// The owning user id
	UserID string `json:"userId"`
}
