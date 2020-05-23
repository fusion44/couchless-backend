package model

import (
	"fmt"
	"path/filepath"
	"time"

	gcontext "github.com/fusion44/ll-backend/context"
)

const (
	// ContentTypeFIT is for binary files from the ANT SDK
	ContentTypeFIT string = "fit"
	// ContentTypeImage are all image file types (JPEG, PNG)
	ContentTypeImage = "image"
)

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
	// `FileType` is the type of the file (FIT, IMAGE, ...)
	ContentType string `json:"contentType"`
}

// IsOwner returns whether the given user is the owner of this file
func (a *FileDescriptor) IsOwner(user *User) bool {
	return a.UserID == user.ID
}

// GetStoragePath constructs the full storage path for the filetype and current user
func (a *FileDescriptor) GetStoragePath(basePath string) (string, error) {
	var p string
	switch a.ContentType {
	case ContentTypeFIT:
		p = filepath.Join(basePath, gcontext.DefaultFITFileDir, a.UserID)
	case ContentTypeImage:
		p = filepath.Join(basePath, gcontext.DefaultImageFileDir, a.UserID)
	default:
		return "", fmt.Errorf("Unsupported file content type: %s", a.ContentType)
	}

	return p, nil
}

// GetFilePath constructs the file path for the filetype and current user
func (a *FileDescriptor) GetFilePath(basePath string) (string, error) {
	storagePath, err := a.GetStoragePath(basePath)
	return filepath.Join(storagePath, a.FileName), err
}
