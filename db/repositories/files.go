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

// GetFileDescriptorByID returns the file descriptor for the given ID
func (r *FileDescRepository) GetFileDescriptorByID(id string) (*model.FileDescriptor, error) {
	var descr model.FileDescriptor
	err := r.DB.Model(&descr).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &descr, nil
}

// GetFileDescriptorByFileName returns the file descriptor for the given ID
func (r *FileDescRepository) GetFileDescriptorByFileName(fn string) (*model.FileDescriptor, error) {
	var descr model.FileDescriptor
	err := r.DB.Model(&descr).Where("file_name = ?", fn).First()

	if err != nil {
		return nil, err
	}

	return &descr, nil
}

// DeleteFileDescriptor deletes a file descriptor from the database
func (r *FileDescRepository) DeleteFileDescriptor(file *model.FileDescriptor) (*model.FileDescriptor, error) {
	_, err := r.DB.Model(file).Where("id = ?", file.ID).Delete()
	return file, err
}
