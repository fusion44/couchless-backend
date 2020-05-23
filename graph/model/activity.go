package model

import "time"

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
	ID        string    `json:"id"`
	FileID    string    `json:"fileId"`
	InputType InputType `json:"inputType"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Comment   string    `json:"comment"`
	SportType string    `json:"sport"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
}

// IsOwner returns whether the given user is the owner of this activity
func (a *Activity) IsOwner(user *User) bool {
	return a.UserID == user.ID
}
