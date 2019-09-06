package keys

var TwilioKeys = initTwilioKeys()
var InternalServiceKeys = initInternalServiceKeys()
var GoogleOAuthKeys = initGoogleAuthKeys()

func initTwilioKeys() TwilioKey {
	var twilioKeys TwilioKey
	Load(&twilioKeys, "stage")

	return twilioKeys
}

func initInternalServiceKeys() InternalServiceKey {
	var internalServiceKey InternalServiceKey
	Load(&internalServiceKey, "stage")

	return internalServiceKey
}

func initGoogleAuthKeys() GoogleOAuthKey {
	var googleOAuthKey GoogleOAuthKey
	Load(&googleOAuthKey, "stage")

	return googleOAuthKey
}
