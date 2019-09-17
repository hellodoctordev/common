package types

import (
	"errors"
	"time"
)

type Interval struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func NewInterval(start time.Time, end time.Time) (interval Interval, err error) {
	if start.After(end) {
		err = errors.New("start time cannot be after end time")
		return
	}

	interval = Interval{
		Start: start,
		End:   end,
	}

	return
}

type ConsultationRequest struct {
	ConsultationID   string   `json:"consultationID"`
	ConsultationType string   `json:"consultationType"`
	RequestedTime    Interval `json:"requestedTime"`
	PatientUserID    string   `json:"patientUserID"`
}
