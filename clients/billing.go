package clients

import (
	"fmt"
	"net/http"
	"os"
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
	PaymentMethodID      string `json:"paymentMethodID"`
}

type CreateAuthorizedChargeResponse struct {
	PaymentIntentID string `json:"paymentIntentID"`
}

func (client *BillingClient) CreateAuthorizedCharge(consultationID, paymentMethodID string) (*http.Response, error) {
	req := CreateAuthorizedChargeRequest{
		PaymentMethodID:paymentMethodID,
	}

	return client.Post(fmt.Sprintf("/charges/consultations/%s/_authorize", consultationID), req)
}

type RequestConsultationCharge struct {

}

func (client *BillingClient) RequestConsultationChargeAuthorization(consultationID string) (*http.Response, error) {
	req := RequestConsultationCharge{}

	return client.Post(fmt.Sprintf("/charges/consultations/%s/_request_authorization", consultationID), req)
}
