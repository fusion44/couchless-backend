package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

// FitUTCTimestamp file timestamps start at 631065600 (UTC 00:00 Dec 31 1989)
const FitUTCTimestamp = 631065600

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

		startTime := time.Unix(int64(f.Session.StartTime+FitUTCTimestamp), 0)
		endTime := time.Unix(int64(f.Session.StartTime+f.Activity.TotalTimerTime+FitUTCTimestamp), 0)

		recs := make([]*model.ActivityRecord, len(f.Records))

		// It can happen that PositionLat and PositionLong is 0
		// I suspect that, when an activity is started on the watch
		// and the GPS is not ready it will lead to these NOK values
		// find the first non 0 values and use these as fallback
		fallbackLat, fallbackLong := findFirstNonNilPos(&f)

		var a = &model.Activity{
			SportType:            f.Sport.Sport,
			StartTime:            startTime,
			EndTime:              endTime,
			TimePaused:           0, // TODO: calculate
			AvgPace:              0, // TODO: calculate
			AvgSpeed:             f.Session.AvgSpeed,
			AvgCadence:           f.Session.AvgCadence,
			AvgFractionalCadence: f.Session.AvgFractionalCadence,
			MaxCadence:           f.Session.MaxCadence,
			MaxSpeed:             f.Session.MaxSpeed,
			TotalDistance:        f.Session.TotalDistance,
			TotalAscent:          f.Session.TotalAscent,
			TotalDescent:         f.Session.TotalDescent,
			MaxAltitude:          f.Session.MaxAltitude,
			AvgHeartRate:         f.Session.AvgHeartRate,
			MaxHeartRate:         f.Session.MaxHeartRate,
			TotalTrainingEffect:  f.Session.TotalTrainingEffect,
		}

		scale := 180.0 / math.Pow(2, 31)

		if len(f.Records) > 0 {
			a.BoundaryNorth = fallbackLat * scale
			a.BoundarySouth = fallbackLat * scale
			a.BoundaryEast = fallbackLong * scale
			a.BoundaryWest = fallbackLong * scale
		}

		for i, rec := range f.Records {
			ts := time.Unix(int64(rec.Timestamp+FitUTCTimestamp), 0)

			// Fill nil position values with the first lat and long pos
			latVal := rec.PositionLat
			if rec.PositionLat == 0 {
				latVal = fallbackLat
			}

			longVal := rec.PositionLong
			if rec.PositionLong == 0 {
				longVal = fallbackLong
			}

			ar := &model.ActivityRecord{
				Timestamp:               ts,
				PositionLat:             latVal * scale,
				PositionLong:            longVal * scale,
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

			if ar.Altitude > a.MaxAltitude {
				a.MaxAltitude = ar.Altitude
			}

			// Find latitude boundaries
			if ar.PositionLat > a.BoundaryNorth {
				a.BoundaryNorth = ar.PositionLat
			} else if ar.PositionLat < a.BoundarySouth {
				a.BoundarySouth = ar.PositionLat
			}

			// Find longitude boundaries
			if ar.PositionLong > a.BoundaryEast {
				a.BoundaryEast = ar.PositionLong
			} else if ar.PositionLong < a.BoundaryWest {
				a.BoundaryWest = ar.PositionLong
			}

			recs[i] = ar
		}

		return &ImportResult{Activity: a, Records: recs}, nil
	}
	return nil, fmt.Errorf("Unable to import file type: %s", f.FileID.Type)
}

func findFirstNonNilPos(pff *model.PrettyFitFile) (float64, float64) {
	var lat, long float64

	for _, rec := range pff.Records {
		if rec.PositionLat != 0 && rec.PositionLong != 0 {
			return rec.PositionLat, rec.PositionLong
		}
	}

	// nothing found. Return nil
	return lat, long
}
