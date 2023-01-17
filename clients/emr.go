package clients

import (
	"log"
	"net/http"
	"os"
)

type EMRServiceClient struct {
	HttpServiceClient
}

func NewEMRServiceClient() *EMRServiceClient {
	serviceHost := os.Getenv("EMR_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	log.Printf("service host: %s", serviceHost)

	return &EMRServiceClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateConsultationIntakeFormRequest struct {
	UserID         string `json:"userID"`
	ConsultationID string `json:"consultationID"`
}

func (client *ConsultationsClient) CreateUserAssessment(req CreateConsultationIntakeFormRequest) (*http.Response, error) {
	return client.Post("/consultations/internal/consultations", req)
}
