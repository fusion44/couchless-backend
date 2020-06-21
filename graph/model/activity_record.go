package model

import "time"

// ActivityRecord provides timestamped and geo located data for an activity
type ActivityRecord struct {
	ID                      string    `json:"id"`
	UserID                  string    `json:"user_id"`
	ActivityID              string    `json:"activity_id"`
	Timestamp               time.Time `json:"timestamp"`
	PositionLat             float64   `json:"position_lat"`
	PositionLong            float64   `json:"position_long"`
	Distance                float64   `json:"distance"`
	TimeFromCourse          int       `json:"time_from_course"`
	CompressedSpeedDistance float64   `json:"compressed_speed_distance"`
	HeartRate               int       `json:"heart_rate"`
	Altitude                float64   `json:"altitude"`
	Speed                   float64   `json:"speed"`
	Power                   int       `json:"power"`
	Grade                   int       `json:"grade"`
	Cadence                 int       `json:"cadence"`
	FractionalCadence       int       `json:"fractional_cadence"`
	Resistance              int       `json:"resistance"`
	CycleLength             int       `json:"cycle_length"`
	Temperature             int       `json:"temperature"`
	AccumulatedPower        int       `json:"accumulated_power"`
}
