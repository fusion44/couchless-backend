package service

import (
	"io"
	"os"
	"path/filepath"

	"github.com/fusion44/ll-backend/context"
	"github.com/op/go-logging"
)

const (
	// DefaultFITFileDir the directory where fit files will be stored
	DefaultFITFileDir = "fit"
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
func (fs *FileService) PersistFile(fileName, userID string, file io.Reader) (*string, error) {
	filePath := filepath.Join(fs.storagePath, DefaultFITFileDir, userID)
	_ = os.MkdirAll(filePath, os.ModePerm)

	filePath = filepath.Join(filePath, fileName)

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fs.log.Errorf("Unable to open file for writing: %s", err)
		return nil, err
	}

	defer f.Close()
	io.Copy(f, file)

	return &filePath, nil
}
