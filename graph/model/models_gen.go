// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewActivity struct {
	StartTime string  `json:"startTime"`
	EndTime   string  `json:"endTime"`
	Comment   *string `json:"comment"`
	SportType string  `json:"sportType"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
