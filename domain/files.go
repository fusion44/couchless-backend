package domain

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	gcontext "github.com/fusion44/ll-backend/context"
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/fusion44/ll-backend/middleware"
	service "github.com/fusion44/ll-backend/services"
)

// HandleSingleFileUpload stores uploaded files to the disk
func (d *Domain) HandleSingleFileUpload(ctx context.Context, file graphql.Upload) (*model.FileDescriptor, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	cfg, err := gcontext.GetConfigFromContext(ctx)
	if err != nil {
		log.Printf("No config in context: %v", err)
		return nil, err
	}

	l, err := middleware.GetLoggerFromContext(ctx)
	if err != nil {
		log.Printf("No logger in context: %v", err)
		return nil, err
	}

	if strings.HasSuffix(file.Filename, ".FIT") {
		// Store the file descriptor to the DB
		fileDesc, err := d.FileRepository.AddFileDescriptor(&model.FileDescriptor{
			FileName:    file.Filename,
			UserID:      u.ID,
			ContentType: model.ContentTypeFIT,
		})

		if err != nil {
			l.Errorf("Unable to store file descriptor for %s: %s", file.Filename, err)
			return nil, fmt.Errorf("Could not upload file: %s", file.Filename)
		}

		fileService := service.NewFileService(cfg, l)
		filePath, persistErr := fileService.PersistFile(fileDesc, file.File)
		if persistErr != nil {
			l.Errorf("Unable to store file: %s\nRemoving file descriptor", filePath)
			// Delete the file descriptor as it's in the DB already
			_, err = d.FileRepository.DeleteFileDescriptor(fileDesc)

			if err != nil {
				l.Errorf("Unable to delete file descriptor: %s", err)
				return nil, fmt.Errorf("Internal server error while uploading: %s", file.Filename)
			}
			return nil, persistErr

		}

		return fileDesc, nil
	}

	// If we are here we didn't recognize the file type.
	l.Errorf("Unknown file uploaded: %s", file.Filename)
	return nil, fmt.Errorf("Unrecognized filetype: %s", file.Filename)
}
