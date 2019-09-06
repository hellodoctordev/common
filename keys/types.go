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

type InternalServiceKey struct {
	KeyType
	ServiceToken	string
	ServiceSecret	string
}

func (internalKey *InternalServiceKey) GetKeyFilePrefix() string {
	return "internal"
}

type GoogleOAuthKey struct {
	KeyType
	ClientID		string
	ClientSecret	string
}

func (googleKey *GoogleOAuthKey) GetKeyFilePrefix() string {
	return "google-oauth"
}
