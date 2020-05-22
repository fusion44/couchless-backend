package model

import "time"

// FileDescriptor type represents a file in the system
type FileDescriptor struct {
	// The `id` is the file id in the database
	ID string `json:"id"`
	// The `fileName` is the original name of the file
	FileName string `json:"fileName"`
	// `createdAt` is the time when the file was uploaded
	CreatedAt time.Time `json:"createdAt"`
	// `userID` is the user ID of the owner
	UserID string `json:"userId"`
}

// IsOwner returns whether the given user is the owner of this file
func (a *FileDescriptor) IsOwner(user *User) bool {
	return a.UserID == user.ID
}
