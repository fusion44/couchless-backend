package model

import (
	"time"
)

// InputType describes the mode with which the activity was added
type InputType string

const (
	// Manual means the Activity was added manually from the user
	Manual InputType = "manual"
	// Recorded means the Activity was recorded by a native app
	Recorded = "recorded"
	// Imported means the Activity was imported from another app
	Imported = "imported"
)

// Activity represents all activity data
type Activity struct {
	ID                   string    `json:"id"`
	FileID               string    `json:"fileId"`
	InputType            InputType `json:"inputType"`
	StartTime            time.Time `json:"startTime"`
	EndTime              time.Time `json:"endTime"`
	Comment              string    `json:"comment"`
	SportType            string    `json:"sport"`
	UserID               string    `json:"userId"`
	BoundaryNorth        float64   `json:"boundaryNorth"`
	BoundarySouth        float64   `json:"boundarySouth"`
	BoundaryEast         float64   `json:"boundaryEast"`
	BoundaryWest         float64   `json:"boundaryWest"`
	TimePaused           int64     `json:"timePaused"`
	AvgPace              float64   `json:"avgPace"`
	AvgSpeed             float64   `json:"avgSpeed"`
	AvgCadence           int       `json:"avgCadence"`
	AvgFractionalCadence int       `json:"AvgFractionalCadence"`
	MaxCadence           int       `json:"maxCadence"`
	MaxSpeed             float64   `json:"maxSpeed"`
	TotalDistance        float64   `json:"totalDistance"`
	TotalAscent          int       `json:"totalAscent"`
	TotalDescent         int       `json:"totalDescent"`
	MaxAltitude          float64   `json:"maxAltitude"`
	AvgHeartRate         int       `json:"avgHeartRate"`
	MaxHeartRate         int       `json:"maxHeartRate"`
	TotalTrainingEffect  float64   `json:"totalTrainingEffect"`
	CreatedAt            time.Time `json:"createdAt"`
}

// NewActivityWithMaxBoundaries create new instance of Activity with
// non-nil boundary values
func NewActivityWithMaxBoundaries() *Activity {
	a := Activity{}
	a.BoundaryNorth = -180.0
	a.BoundarySouth = 180.0
	a.BoundaryEast = -180.0
	a.BoundaryWest = 180.0
	return &a
}

// IsOwner returns whether the given user is the owner of this activity
func (a *Activity) IsOwner(user *User) bool {
	return a.UserID == user.ID
}
