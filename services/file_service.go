package service

import (
	"fmt"
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
	p := filepath.Join(config.FileStoragePath, DefaultFITFileDir)
	_ = os.MkdirAll(p, os.ModePerm)
	return &FileService{storagePath: p, log: log}
}

// PersistFile saves a given file file to the disk
func (fs *FileService) PersistFile(fileName string, file io.Reader) (*string, error) {
	filePath := filepath.Join(fs.storagePath, fileName)
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer f.Close()
	io.Copy(f, file)

	return &filePath, nil
}
