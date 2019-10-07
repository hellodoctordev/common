package clients

import (
	"cloud.google.com/go/firestore"
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

	return &ConsultationsClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateConsultationRequest struct {
	PatientUserUID     string    `json:"patientUserUID"`
	ProviderUserUID    string    `json:"providerUserUID"`
	ConsultationType   string    `json:"consultationType"`   // see consultations.types for enum values
	ConsultationStatus string    `json:"consultationStatus"` // see consultations.types for enum values
	RequestMode        string    `json:"requestMode"`        // see scheduling.types for enum values
	RequestedStart     time.Time `json:"requestedStart"`
	RequestedEnd       time.Time `json:"requestedEnd"`
}

type CreateConsultationResponse struct {
	Chat         *firestore.DocumentRef `json:"chatRef"`
	Consultation *firestore.DocumentRef `json:"consultationRef"`
}

func (client *ConsultationsClient) CreateConsultation(req CreateConsultationRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations", req)
}
