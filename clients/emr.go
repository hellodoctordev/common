package clients

import (
	"fmt"
	"github.com/hellodoctordev/common/emr/resources"
	"github.com/hellodoctordev/common/utils"
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
	Resource resources.Resource
}

type CreateResourceResponse struct {
	ID           string `json:"id"`
	ResourceType string `json:"resourceType"`
}

func (client *EMRClient) CreateResource(req CreateResourceRequest) (response CreateResourceResponse, err error) {
	res, err := client.Post(fmt.Sprintf("/emr/resources/%s", req.Resource.GetResourceType()), req.Resource)
	if err != nil {
		return response, err
	}

	err = utils.ReadBody(res.Body, &response)

	return response, nil
}
