package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hellodoctordev/common/keys"
	"log"
	"net/http"
)

const (
	DefaultServiceHost = "http://localhost:8080"
)

type HttpServiceClient struct {
	*http.Client
	ServiceHost string
}

func (client *HttpServiceClient) Post(path, body interface{}) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s%s", client.ServiceHost, path)
	log.Printf("url: %s", url)
	reqBody, err := json.Marshal(body)
	if err != nil {
		log.Printf("error occurred marshalling interface: %s", err)
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("error occurred creating new request: %s", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Internal-Authorization", keys.InternalServiceKeys.ServiceToken)
	log.Printf("doing req")
	return client.Do(req)
}
