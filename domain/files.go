package domain

import (
	"context"
	"log"
	"strings"

	"github.com/99designs/gqlgen/graphql"
	gcontext "github.com/fusion44/ll-backend/context"
	"github.com/fusion44/ll-backend/middleware"
	service "github.com/fusion44/ll-backend/services"
)

// HandleSingleFileUpload stores uploaded files to the disk
func (d *Domain) HandleSingleFileUpload(ctx context.Context, file graphql.Upload) (bool, error) {
	cfg, err := gcontext.GetConfigFromContext(ctx)
	if err != nil {
		log.Printf("No config in context: %v", err)
		return false, err
	}

	l, err := middleware.GetLoggerFromContext(ctx)
	if err != nil {
		log.Printf("No logger in context: %v", err)
		return false, err
	}

	if strings.HasSuffix(file.Filename, ".FIT") {
		fileService := service.NewFileService(cfg, l)
		filePath, err := fileService.PersistFile(file.Filename, file.File)
		if err != nil {
			l.Errorf("Unable to store file: %s", filePath)
			return false, err
		}
	}

	return true, nil
}
