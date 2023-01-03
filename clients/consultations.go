package clients

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/firestore"
)

type ConsultationsClient struct {
	HttpServiceClient
}

func NewConsultationsClient() *ConsultationsClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	log.Printf("service host: %s", serviceHost)

	return &ConsultationsClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateConsultationRequest struct {
	PatientUserUID      string                 `json:"patientUserUID"`
	PractitionerUserUID string                 `json:"practitionerUserUID"`
	ConsultationType    string                 `json:"consultationType"`   // see consultations.types for enum values
	ConsultationStatus  string                 `json:"consultationStatus"` // see consultations.status for enum values
	ScheduledStart      time.Time              `json:"requestedStart"`
	ScheduledEnd        time.Time              `json:"requestedEnd"`
	Reason              string                 `json:"reason"`
	RequestMode         string                 `json:"requestMode"`
	Specialty           string                 `json:"specialty"`
	Forms               map[string]interface{} `json:"forms"`
}

type CreateConsultationResponse struct {
	Chat         *firestore.DocumentRef `json:"chatRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

func (client *ConsultationsClient) CreateConsultation(req CreateConsultationRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations", req)
}

func (client *ConsultationsClient) CancelConsultation(consultationID string) (*http.Response, error) {
	return client.Post(fmt.Sprintf("/consultations/internal/consultations/%s/_cancel", consultationID), nil)
}

type StartChatConsultationRequest struct {
	Practitioner *firestore.DocumentRef `json:"practitionerRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

type StartChatConsultationResponse struct {
	Chat         *firestore.DocumentRef `json:"chatRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

func (client *ConsultationsClient) StartChatConsultation(req StartChatConsultationRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations/start-chat", req)
}

type EmitConsultationEventRequest struct {
	ConsultationID string `json:"consultationID"`
	ActorID        string `json:"actorID"`
	Type           string `json:"type"`
	Value          string `json:"value"`
}

func (client *ConsultationsClient) EmitConsultationEvent(req EmitConsultationEventRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/events", req)
}
