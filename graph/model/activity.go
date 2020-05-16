package model

// Activity represents all activity data
type Activity struct {
	ID        string `json:"id"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Comment   string `json:"comment"`
	SportType string `json:"sport"`
	UserID    string `json:"userId"`
	CreatedAt string `json:"createdAt"`
}

// IsOwner returns whether the given user is the owner of this activity
func (a *Activity) IsOwner(user *User) bool {
	return a.UserID == user.ID
}
