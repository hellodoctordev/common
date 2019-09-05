package keys

var TwilioKeys = initTwilioKeys()
var InternalServiceKeys = initInternalServiceKeys()

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
