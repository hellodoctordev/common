package keys

type KeyType interface {
	GetKeyFilePrefix()	string
}

type TwilioKey struct {
	KeyType
	AccountSID        	string
	AuthToken         	string
	APIKeySID         	string
	APIKeySecret      	string
	PushCredentialSID 	string
	TwimlApplicationSID string
}

func (twilioKey *TwilioKey) GetKeyFilePrefix() string {
	return "twilio"
}
