package clients

import (
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

func (client *SearchClient) PostDoctor(doctorData map[string]interface{}) (*http.Response, error) {
	return client.Post("/search/internal/doctors", doctorData)
}
