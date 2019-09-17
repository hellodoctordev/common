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

func (i *Interval) Duration() time.Duration {
	return i.End.Sub(i.Start)
}

func (i *Interval) Equal(i2 Interval) bool {
	return i.Start.Equal(i2.Start) && i.End.Equal(i2.End)
}

func (i *Interval) Abuts(i2 Interval) bool {
	return i.Start.Equal(i2.End) || i.End.Equal(i2.Start)
}

func (i *Interval) Contains(i2 Interval) bool {
	return (i.Start.Before(i2.Start) || i.Start.Equal(i2.Start)) && (i.End.After(i2.End) || i.End.Equal(i2.End))
}

func (i *Interval) Overlaps(i2 Interval) bool {
	if i.Equal(i2) || i.Contains(i2) || i2.Contains(*i) {
		return true
	}

	if i.Abuts(i2) {
		return false
	}

	return i.Start.Equal(i2.Start) && i.End.Before(i2.End) ||
		i.Start.After(i2.Start) && i.End.Equal(i2.End) ||
		i.Start.Before(i2.Start) && i.End.After(i2.Start) ||
		i2.Start.Before(i.Start) && i2.End.After(i.Start)
}


type ConsultationRequest struct {
	ConsultationID   string   `json:"consultationID"`
	ConsultationType string   `json:"consultationType"`
	RequestedTime    Interval `json:"requestedTime"`
	PatientUserID    string   `json:"patientUserID"`
}
