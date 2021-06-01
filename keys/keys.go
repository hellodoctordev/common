package keys

var TwilioKeys = initTwilioKeys()
var InternalServiceKeys = initInternalServiceKeys()
var AdminKeys = initAdminKeys()
var GoogleApiKeys = initGoogleApiKeys()
var GoogleOAuthKeys = initGoogleAuthKeys()

func initTwilioKeys() TwilioKey {
	var keys TwilioKey
	Load(&keys)

	return keys
}

func initInternalServiceKeys() InternalServiceKey {
	var keys InternalServiceKey
	Load(&keys)

	return keys
}

func initAdminKeys() AdminKey {
	var keys AdminKey
	Load(&keys)

	return keys
}

func initGoogleApiKeys() GoogleApiKey {
	var keys GoogleApiKey
	Load(&keys)

	return keys
}

func initGoogleAuthKeys() GoogleOAuthKey {
	var keys GoogleOAuthKey
	Load(&keys)

	return keys
}

