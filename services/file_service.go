package service

import (
	"io"
	"os"
	"path/filepath"

	"github.com/fusion44/ll-backend/context"
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/op/go-logging"
)

// FileService handles all file operations
type FileService struct {
	storagePath string
	log         *logging.Logger
}

// NewFileService creates a new upload service instance
func NewFileService(config *context.Config, log *logging.Logger) *FileService {
	return &FileService{storagePath: config.FileStoragePath, log: log}
}

// PersistFile saves a given file to the disk
func (fs *FileService) PersistFile(fileDesc *model.FileDescriptor, file io.Reader) (*string, error) {
	filePath, err := fileDesc.GetStoragePath(fs.storagePath)
	if err != nil {
		fs.log.Errorf("Unable to get full file path: %s", err)
		return nil, err
	}
	_ = os.MkdirAll(filePath, os.ModePerm)

	filePath = filepath.Join(filePath, fileDesc.FileName)

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fs.log.Errorf("Unable to open file for writing: %s", err)
		return nil, err
	}

	defer f.Close()
	io.Copy(f, file)

	return &filePath, nil
}

// FileExists checks whether the given file exists
func (fs *FileService) FileExists(fileDesc *model.FileDescriptor) (bool, error) {
	filePath, err := fileDesc.GetStoragePath(fs.storagePath)

	if err != nil {
		return false, err
	}

	filePath = filepath.Join(filePath, fileDesc.FileName)

	if _, err := os.Stat(filePath); err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
