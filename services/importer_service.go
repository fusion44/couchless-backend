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
	Activity *model.Activity
	Records  []*model.ActivityRecord
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

		var a = &model.Activity{
			SportType: f.Sport.Sport,
			StartTime: startTime,
			EndTime:   endTime,
		}
		recs := make([]*model.ActivityRecord, len(f.Records))

		// It can happen that PositionLat and PositionLong is 0
		// I suspect that, when an activity is started on the watch
		// and the GPS is not ready it will lead to these NOK values
		// find the first non 0 values and use these as fallback
		fallbackLat, fallbackLong := findFirstNonNilPos(&f)

		for i, rec := range f.Records {
			ts := time.Unix(int64(rec.Timestamp), 0)

			// Fill nil position values with the first lat and long pos
			latVal := int32(rec.PositionLat)
			if rec.PositionLat == 0 {
				latVal = fallbackLat
			}

			longVal := int32(rec.PositionLong)
			if rec.PositionLong == 0 {
				longVal = fallbackLong
			}

			ar := &model.ActivityRecord{
				Timestamp:               ts,
				PositionLat:             int(latVal),
				PositionLong:            int(longVal),
				Distance:                rec.Distance,
				TimeFromCourse:          int(rec.TimeFromCourse),
				CompressedSpeedDistance: rec.CompressedSpeedDistance,
				HeartRate:               int(rec.HeartRate),
				Altitude:                rec.Altitude,
				Speed:                   rec.Speed,
				Power:                   int(rec.Power),
				Grade:                   int(rec.Grade),
				Cadence:                 int(rec.Cadence),
				FractionalCadence:       int(rec.FractionalCadence),
				Resistance:              int(rec.Registance),
				CycleLength:             int(rec.CycleLength),
				Temperature:             int(rec.Temperature),
				AccumulatedPower:        int(rec.AccumulatedPower),
			}
			recs[i] = ar
		}

		return &ImportResult{Activity: a, Records: recs}, nil
	}
	return nil, fmt.Errorf("Unable to import file type: %s", f.FileID.Type)
}

func findFirstNonNilPos(pff *model.PrettyFitFile) (int32, int32) {
	var lat, long int32

	for _, rec := range pff.Records {
		if rec.PositionLat != 0 && rec.PositionLong != 0 {
			return int32(rec.PositionLat), int32(rec.PositionLong)
		}
	}

	// nothing found. Return nil
	return lat, long
}
