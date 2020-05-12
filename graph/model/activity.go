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
