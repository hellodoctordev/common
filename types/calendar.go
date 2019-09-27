package types

import "C"
import (
	"cloud.google.com/go/firestore"
	"time"
)

type Availability string

const (
	Busy      Availability = "busy"
	Available Availability = "available"
	Pending   Availability = "pending"
)

type EventType string

const (
	EventTypeConsultationRequest EventType = "consultationRequest"
)

type CalendarEvent interface {
	Type() EventType
	Title() string
	Description() string
	Start() time.Time
	End() time.Time
	Availability() Availability
	Participants() []*firestore.DocumentRef
}

type CalendarEventDocument struct {
	Type         EventType               `firestore:"type"`
	Title        string                  `firestore:"title"`
	Description  string                  `firestore:"description"`
	Start        time.Time               `firestore:"start"`
	End          time.Time               `firestore:"end"`
	Availability Availability            `firestore:"availability"`
	Participants []*firestore.DocumentRef `firestore:"participants"`
}
