package keys

var TwilioKeys = initTwilioKeys()

func initTwilioKeys() TwilioKey {
	var twilioKeys TwilioKey
	Load(&twilioKeys, "stage")

	return twilioKeys
}
