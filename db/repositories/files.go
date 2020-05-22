package repositories

import (
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/go-pg/pg/v9"
)

// FileDescRepository contains all functions regarding file management
type FileDescRepository struct {
	DB *pg.DB
}

// AddFileDescriptor inserts a file descriptor into the database
func (r *FileDescRepository) AddFileDescriptor(file *model.FileDescriptor) (*model.FileDescriptor, error) {
	_, err := r.DB.Model(file).Returning("*").Insert()
	return file, err
}

// DeleteFileDescriptor deletes a file descriptor from the database
func (r *FileDescRepository) DeleteFileDescriptor(file *model.FileDescriptor) (*model.FileDescriptor, error) {
	_, err := r.DB.Model(file).Where("id = ?", file.ID).Delete()
	return file, err
}
