package utils

import (
	"encoding/json"
	"github.com/hellodoctordev/common/logging"
	"io/ioutil"
	"log"
	"net/http"
)

func ReadRequestBody(r *http.Request, o interface{}) (err error) {
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

func WriteJSONResponse(w http.ResponseWriter, response interface{}) {
	js, err := json.Marshal(response)
	if err != nil {
		logging.Error("could not marshal CreateChatConsultationResponse: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}