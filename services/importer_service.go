package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/fusion44/ll-backend/graph/model"
	"github.com/op/go-logging"
)

// ImportResult contains the result of the loading process
type ImportResult struct {
	Activity model.Activity
}

// ImporterService handles conversions
type ImporterService struct {
	storagePath string
	log         *logging.Logger
}

// NewImporterService imports data to the database
func NewImporterService(log *logging.Logger) *ImporterService {
	return &ImporterService{log: log}
}

// ImportFITJSON a fit2json file
func (cs *ImporterService) ImportFITJSON(jsonFilePath string) (*ImportResult, error) {
	file, err := os.Open(jsonFilePath)
	if err != nil {
		log.Fatal(err)
	}

	js, err := ioutil.ReadAll(file)
	var f model.PrettyFitFile
	json.Unmarshal(js, &f)

	if f.FileID.Type == "activity" {
		startTime := time.Unix(int64(f.Session.StartTime), 0)
		endTime := time.Unix(int64(f.Session.StartTime+f.Activity.TotalTimerTime), 0)

		var a = model.Activity{
			SportType: f.Sport.Sport,
			StartTime: startTime,
			EndTime:   endTime,
		}

		// TODO: read location data

		return &ImportResult{Activity: a}, nil
	}
	return nil, fmt.Errorf("Unable to import file type: %s", f.FileID.Type)
}
