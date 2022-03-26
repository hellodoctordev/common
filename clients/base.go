package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hellodoctordev/common/keys"
)

const (
	DefaultServiceHost = "http://localhost:8080"
)

type HttpServiceClient struct {
	*http.Client
	ServiceHost string
}

func (client *HttpServiceClient) Get(path string) (resp *http.Response, err error) {
	return client.doRequest("GET", path, nil)
}

func (client *HttpServiceClient) Post(path, body interface{}) (resp *http.Response, err error) {
	return client.doRequest("POST", path, body)
}

func (client *HttpServiceClient) Put(path, body interface{}) (resp *http.Response, err error) {
	return client.doRequest("PUT", path, body)
}

func (client *HttpServiceClient) Delete(path, body interface{}) (resp *http.Response, err error) {
	return client.doRequest("DELETE", path, body)
}

func (client *HttpServiceClient) doRequest(method string, path, body interface{}) (resp *http.Response, err error) {
	url := fmt.Sprintf("%s%s", client.ServiceHost, path)

	var req *http.Request

	if body != nil {
		reqBody, marshallErr := json.Marshal(body)
		if err != nil {
			log.Printf("error occurred marshalling interface: %s", marshallErr)
			return nil, marshallErr
		}

		req, err = http.NewRequest(method, url, bytes.NewBuffer(reqBody))
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Printf("error occurred creating new request: %s", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Internal-Authorization", keys.InternalServiceKeys.ServiceToken)

	return client.Do(req)
}
