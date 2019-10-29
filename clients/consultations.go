package clients

import (
	"cloud.google.com/go/firestore"
	"log"
	"net/http"
	"os"
	"time"
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
	PatientUserUID     string     `json:"patientUserUID"`
	ProviderUserUID    string     `json:"providerUserUID"`
	ConsultationType   string     `json:"consultationType"`   // see consultations.types for enum values
	ConsultationStatus string     `json:"consultationStatus"` // see consultations.types for enum values
	RequestMode        string     `json:"requestMode"`        // see scheduling.types for enum values
	RequestedStart     *time.Time `json:"requestedStart"`
	RequestedEnd       *time.Time `json:"requestedEnd"`
}

type CreateConsultationResponse struct {
	Chat         *firestore.DocumentRef `json:"chatRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

func (client *ConsultationsClient) CreateConsultation(req CreateConsultationRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations", req)
}

type StartChatConsultationRequest struct {
	Provider     *firestore.DocumentRef `json:"providerRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

type StartChatConsultationResponse struct {
	Chat         *firestore.DocumentRef `json:"chatRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

func (client *ConsultationsClient) StartChatConsultation(req StartChatConsultationRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations/start-chat", req)
}
