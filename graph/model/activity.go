package model

// Activity represents all activity data
type Activity struct {
	ID      string `json:"id"`
	Comment string `json:"comment"`
	UserID  string `json:"user"`
}
