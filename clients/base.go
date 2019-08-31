package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	DefaultServiceHost = "http://stage.api.hellodoctor.com.mx"
)

type HttpServiceClient struct {
	*http.Client
	ServiceHost string
}

func (client *HttpServiceClient) Post(path, body interface{}) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s/%s", client.ServiceHost, path)

	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("error occurred marshalling interface: %s", err)
	}

	return http.DefaultClient.Post(url, "application/json", bytes.NewBuffer(reqBody))
}
