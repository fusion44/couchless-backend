package service

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fusion44/ll-backend/context"
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/op/go-logging"
)

const (
	// DefaultJSONFileDir the directory where json files will be stored
	DefaultJSONFileDir = "fit_json"
)

// ConverterService handles conversions
type ConverterService struct {
	fit2JSONPath string
	basePath     string
	log          *logging.Logger
}

// NewConverterService converts files and data from one format to another
func NewConverterService(config *context.Config, log *logging.Logger) *ConverterService {
	p := filepath.Join(config.FileStoragePath, DefaultJSONFileDir)
	return &ConverterService{fit2JSONPath: config.Fit2JSONPath, basePath: p, log: log}
}

// ConvertFITtoJSON converts a .FIT file to JSON
// fitFilePath is the path to the .FIT file
// Returns the file path of the JSON file or an error
func (cs *ConverterService) ConvertFITtoJSON(user *model.User, fitFilePath string) (string, error) {
	p := filepath.Join(cs.basePath, user.ID)
	_ = os.MkdirAll(p, os.ModePerm)

	jsonFilePath := filepath.Join(p, filepath.Base(fitFilePath)+".json")

	cmd := &exec.Cmd{Path: cs.fit2JSONPath, Args: []string{
		cs.fit2JSONPath,
		"-p",
		"-i",
		fitFilePath,
		"-o",
		jsonFilePath,
	}, Stdout: os.Stdout, Stderr: os.Stdout}

	if err := cmd.Run(); err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}

	return jsonFilePath, nil
}
