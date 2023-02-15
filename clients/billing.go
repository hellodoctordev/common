package clients

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type BillingClient struct {
	HttpServiceClient
}

func NewBillingClient() *BillingClient {
	serviceHost := os.Getenv("BILLING_SERVICE_HOST")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &BillingClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateAuthorizedChargeRequest struct {
	PatientID       string     `json:"patientID"`
	Amount          int64      `json:"amount"`
	PaymentMethodID string     `json:"paymentMethodID"`
	ConsultationID  string     `json:"consultationID"`
	PractitionerID  string     `json:"practitionerID"`
	CreateHold      bool       `json:"createHold"`
	Service         string     `json:"service"`
	ServiceTime     *time.Time `json:"serviceTime"`
}

type CreateAuthorizedChargeResponse struct {
	AuthorizedChargeID string `json:"authorizedChargeID"`
	PaymentMethodType  string `json:"paymentMethodType"`
	Code               string `json:"code"`
	DeclineCode        string `json:"declineCode"`
}

func (client *BillingClient) CreateAuthorizedCharge(req CreateAuthorizedChargeRequest) (*http.Response, error) {
	return client.Post("/charges/_authorize", req)
}

func (client *BillingClient) CancelAuthorizedCharge(consultationID string) (*http.Response, error) {
	return client.Post(fmt.Sprintf("/charges/consultations/%s/_cancel", consultationID), nil)
}

type RequestConsultationCharge struct {
}

func (client *BillingClient) RequestConsultationChargeAuthorization(consultationID string) (*http.Response, error) {
	req := RequestConsultationCharge{}

	return client.Post(fmt.Sprintf("/charges/consultations/%s/_request_authorization", consultationID), req)
}
