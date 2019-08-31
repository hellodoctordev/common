package clients

import (
	"github.com/hellodoctordev/gotwilio"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"time"
)

const (
	AccountSID = "AC7bc6ffc0bfc7e827324b14bab30a02ca"
	AuthToken = "dd39ec20c57107d5cc27cb29983da79a"
	APIKeySID = "SKff8edd07d763692383726364d868d921"
	APIKeySecret = "NfxxSqWorskCFNWnHacAW8WtiRrNZJ3C"
)

func NewTwilioClient() *gotwilio.Twilio {
	return gotwilio.NewTwilioClient(AccountSID, AuthToken)
}

func NewApplicationAccessToken(twilio *gotwilio.Twilio, identity string) *gotwilio.AccessToken {
	token := twilio.NewAccessToken()
	token.APIKeySid = APIKeySID
	token.APIKeySecret = APIKeySecret
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

	params, _ := url.ParseQuery(string(body))

	v := reflect.ValueOf(o).Elem()
	typeOfO := v.Type()

	for i := 0; i < v.NumField(); i++ {
		typeField := typeOfO.Field(i)

		value := params.Get(typeField.Name)

		v.FieldByName(typeField.Name).SetString(value)
	}

	return
}
