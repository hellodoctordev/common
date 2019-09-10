package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadBody(r *http.Request, o interface{}) (err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading body: %s", err)
		return
	}

	err = json.Unmarshal(body, o)
	if err != nil {
		log.Printf("error unmarshaling payload: %s", err)
	}

	return
}
