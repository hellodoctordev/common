package clients

import (
	"github.com/hellodoctordev/common/types"
	"log"
	"net/http"
	"os"
)

type SearchClient struct {
	HttpServiceClient
}

func NewSearchClient() *SearchClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &SearchClient{
		HttpServiceClient{
			Client: http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

func (client *SearchClient) PostDoctor(uid string, doctorData map[string]interface{}) (res *http.Response, err error) {
	doctorDocument, err := types.FromUserData(uid, doctorData)
	if err != nil {
		log.Printf("error occurred creating doctor document: %s", err)
		return
	}

	return client.Post("/search/internal/doctors", doctorDocument)
}
