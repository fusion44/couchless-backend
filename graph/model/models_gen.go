// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/99designs/gqlgen/graphql"
)

type ActivityFilter struct {
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Comment   *string    `json:"comment"`
	SportType *string    `json:"sportType"`
}

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
	User      *User      `json:"user"`
}

type AuthToken struct {
	AccessToken string    `json:"accessToken"`
	ExpiredAt   time.Time `json:"expiredAt"`
}

// The `ImportActivity` input represents a to imported activity
type ImportActivity struct {
	// The `fileID` is the ID of a `FileDescriptor`
	FileID string `json:"fileID"`
	// The `comment` is an optional comment to be added to the activity
	Comment *string `json:"comment"`
}

// The `LoginInput` type represents the required login input
type LoginInput struct {
	// The `username` can either be an email or the actual username
	Username string `json:"username"`
	// The `password` length must be 8 characters minimum
	Password string `json:"password"`
}

type NewActivity struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
	Comment   *string   `json:"comment"`
	SportType string    `json:"sportType"`
}

type RegisterInput struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirmPassword"`
}

type UpdateActivity struct {
	ID        string     `json:"id"`
	StartTime *time.Time `json:"startTime"`
	EndTime   *time.Time `json:"endTime"`
	Comment   *string    `json:"comment"`
	SportType *string    `json:"sportType"`
}

type UploadFile struct {
	ID   int            `json:"id"`
	File graphql.Upload `json:"file"`
}
