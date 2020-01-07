package clients

import (
	"fmt"
	"github.com/hellodoctordev/common/emr/resources"
	"log"
	"net/http"
	"os"
)

type EMRClient struct {
	HttpServiceClient
}

func NewEMRClient() *EMRClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	log.Printf("service host: %s", serviceHost)

	return &EMRClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateResourceRequest struct {
	ResourceType string
	Resource resources.Resource
}

type CreateResourceResponse struct {
	Payload map[string]interface{} `json:"payload"`
}

func (client *EMRClient) CreateResource(req CreateResourceRequest) (*http.Response, error) {
	return client.Post(fmt.Sprintf("/emr/resources/%s", req.ResourceType), req.Resource)
}