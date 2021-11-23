package clients

import (
	"fmt"
	"net/http"
	"os"
)

type ApiClient struct {
	HttpServiceClient
}

func NewApiClient() *ApiClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &ApiClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type GetDocScoreResponse struct {
	DocScore float64                  `json:"docScore"`
	Metrics  []map[string]interface{} `json:"metrics"`
}

func (client *ApiClient) GetDocScore(practitionerID string) (*http.Response, error) {
	return client.Get(fmt.Sprintf("/docscore/%s", practitionerID))
}
