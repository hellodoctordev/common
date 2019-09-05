package clients

import (
	"encoding/json"
	"github.com/hellodoctordev/common/keys"
	"github.com/hellodoctordev/gotwilio"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func NewTwilioClient() *gotwilio.Twilio {
	return gotwilio.NewTwilioClient(keys.TwilioKeys.AccountSID, keys.TwilioKeys.AuthToken)
}

func NewApplicationAccessToken(twilio *gotwilio.Twilio, identity string) *gotwilio.AccessToken {
	token := twilio.NewAccessToken()
	token.APIKeySid = keys.TwilioKeys.APIKeySID
	token.APIKeySecret = keys.TwilioKeys.APIKeySecret
	token.ExpiresAt = time.Now().Add(time.Hour * 4)
	token.NotBefore = time.Now()
	token.Identity = identity

	return token
}

func UnmarshalRequestBody(r *http.Request, o interface{}) (err error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		return
	}

	return json.Unmarshal(body, o)
}
