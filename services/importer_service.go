package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"time"

	"github.com/fusion44/couchless-backend/graph/model"
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
		var startTime, endTime time.Time
		var finalStopEvent, numPauses int
		numEvent := len(f.Events)

		// The last stop event is not necessarily the very last event => search for it
		// We need for the pause time calculations below
		for i := numEvent - 1; i >= 0; i-- {
			if f.Events[i].Event == "timer" && f.Events[i].EventType == "stop_all" {
				finalStopEvent = i
				break
			}
		}

		firstStartFound := false
		totalPauseTime := 0
		lastStopTime := 0
		for i, e := range f.Events {
			if e.Event == "timer" && e.EventType == "start" && !firstStartFound {
				// first start
				firstStartFound = true
				startTime = time.Unix(int64(e.Timestamp+FitUTCTimestamp), 0)
			} else if e.Event == "timer" && e.EventType == "start" && firstStartFound {
				// subsequent start, we continue the activity here
				totalPauseTime += e.Timestamp - lastStopTime
				lastStopTime = 0
				numPauses++
			} else if e.Event == "timer" && e.EventType == "stop_all" && i != finalStopEvent {
				lastStopTime = e.Timestamp
			} else if i == finalStopEvent {
				endTime = time.Unix(int64(e.Timestamp+FitUTCTimestamp), 0)
			} else {
				// ignore
			}
		}

		recs := make([]*model.ActivityRecord, len(f.Records))

		// It can happen that PositionLat and PositionLong is 0
		// I suspect that, when an activity is started on the watch
		// and the GPS is not ready it will lead to these NOK values
		// find the first non 0 values and use these as fallback
		fallbackLat, fallbackLong := findFirstNonNilPosLatLong(&f)

		duration := int(endTime.Sub(startTime).Seconds())

		var a = &model.Activity{
			SportType:            f.Sport.Sport,
			StartTime:            startTime,
			EndTime:              endTime,
			Duration:             duration,
			TimePaused:           totalPauseTime,
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
			TotalTrainingEffect:  float64(f.Session.TotalTrainingEffect) / 10,
		}

		scale := 180.0 / math.Pow(2, 31)

		if len(f.Records) > 0 {
			a.BoundaryNorth = fallbackLat * scale
			a.BoundarySouth = fallbackLat * scale
			a.BoundaryEast = fallbackLong * scale
			a.BoundaryWest = fallbackLong * scale
		}

		var lastLatVal, lastLongVal float64

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

			altitude := rec.Altitude
			if rec.Altitude == 0 {
				altitude = findFirstNonNilAltitude(&f)
			}

			latVal *= scale
			longVal *= scale

			// Sometimes watches deliver wrong results.
			// We bridge the gap by using the last valid value
			// until OK. A better approach would be to interpolate
			// the values.
			// TODO: interpolate values
			if latVal < -90 || latVal > 90 {
				latVal = lastLatVal
			}

			if longVal < -90 || longVal > 90 {
				longVal = lastLongVal
			}

			lastLatVal = latVal
			lastLongVal = longVal

			ar := &model.ActivityRecord{
				Timestamp:               ts,
				PositionLat:             latVal,
				PositionLong:            longVal,
				Distance:                rec.Distance,
				TimeFromCourse:          int(rec.TimeFromCourse),
				CompressedSpeedDistance: rec.CompressedSpeedDistance,
				HeartRate:               int(rec.HeartRate),
				Altitude:                altitude,
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

		if a.BoundaryEast < -90 || a.BoundaryEast > 90 || a.BoundaryWest < -90 || a.BoundaryWest > 90 {
			a.BoundaryEast = 0
			a.BoundaryWest = 0
		}

		if a.BoundaryNorth < -90 || a.BoundaryNorth > 90 || a.BoundarySouth < -90 || a.BoundarySouth > 90 {
			a.BoundaryNorth = 0
			a.BoundarySouth = 0
		}

		return &ImportResult{Activity: a, Records: recs}, nil
	}
	return nil, fmt.Errorf("Unable to import file type: %s", f.FileID.Type)
}

func findFirstNonNilPosLatLong(pff *model.PrettyFitFile) (float64, float64) {
	var lat, long float64

	for _, rec := range pff.Records {
		if rec.PositionLat != 0 && rec.PositionLong != 0 {
			return rec.PositionLat, rec.PositionLong
		}
	}

	// nothing found. Return nil
	return lat, long
}

func findFirstNonNilAltitude(pff *model.PrettyFitFile) float64 {
	for _, rec := range pff.Records {
		if rec.Altitude != 0 {
			return rec.Altitude
		}
	}
	return 0
}
