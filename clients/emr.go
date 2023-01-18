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
	ConsultationID string `json:"consultationID"`
}

func (client *EMRServiceClient) CreateConsultationIntakeForm(userID, consultationID string) (*http.Response, error) {
	header := Header{
		Key:   "X-On-Behalf-Of",
		Value: userID,
	}

	req := CreateConsultationIntakeFormRequest{
		ConsultationID: consultationID,
	}

	return client.Post("/forms", req, header)
}

type CreateUserDiagnosisRequest struct {
	ConsultationID string                 `json:"consultationID"`
	Diagnosis      map[string]interface{} `json:"diagnosis"`
}

func (client *EMRServiceClient) CreateUserDiagnosis(userID, consultationID string, diagnosis map[string]interface{}) (*http.Response, error) {
	header := Header{
		Key:   "X-On-Behalf-Of",
		Value: userID,
	}

	req := CreateUserDiagnosisRequest{
		ConsultationID: consultationID,
		Diagnosis:      diagnosis,
	}

	return client.Post("/diagnoses", req, header)
}
