package clients

import (
	"github.com/hellodoctordev/common/types"
	"log"
	"net/http"
)

type SearchClient struct {
	HttpServiceClient
}

func NewSearchClient() *SearchClient {
	return &SearchClient{
		HttpServiceClient{
			Client: http.DefaultClient,
			ServiceHost: DefaultServiceHost,
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
