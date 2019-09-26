package types

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	"fmt"
	"github.com/hellodoctordev/common/firebase"
	"github.com/hellodoctordev/common/logging"
	"time"
)

type Interval struct {
	Start time.Time `firestore:"start" json:"start"`
	End   time.Time `firestore:"end" json:"end"`
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

type ConsultationSessionRequest struct {
	Consultation     *firestore.DocumentRef `firestore:"consultation" json:"consultationRef"`
	PatientUser      *firestore.DocumentRef `firestore:"patientUser" json:"patientUserRef"`
	ConsultationType string                `firestore:"consultationType" json:"consultationType"`
	RequestedTime    Interval              `firestore:"requestedTime" json:"requestedTime"`
}

func (c *ConsultationSessionRequest) Type() EventType {
	return EventTypeConsultationRequest
}

func (c *ConsultationSessionRequest) Title() string {
	client := firebase.NewFirestoreClient()
	patientUserProfileSnapshot, err := client.Doc(fmt.Sprintf("profiles/%s", c.PatientUser.ID)).Get(context.Background())
	if err != nil {
		logging.Error("error getting patient %s snapshot: %s", c.PatientUser.ID, err)
		return "Consultation Request"
	}

	patientUserDisplayName, err := patientUserProfileSnapshot.DataAt("displayName")
	if err != nil {
		logging.Error("error getting patient %s display name: %s", c.PatientUser.ID, err)
		return "Consultation Request"
	}

	return fmt.Sprintf("Consultation requested by %s", patientUserDisplayName)
}

func (c *ConsultationSessionRequest) Start() time.Time {
	return c.RequestedTime.Start
}

func (c *ConsultationSessionRequest) End() time.Time {
	return c.RequestedTime.End
}

func (c *ConsultationSessionRequest) Availability() Availability {
	return Pending
}

func (c *ConsultationSessionRequest) Participants() []firestore.DocumentRef {
	return []firestore.DocumentRef{c.PatientUser}
}
