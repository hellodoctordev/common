package crypto

import "crypto/rsa"

type UserPublicKeyData struct {
	UserUID     string `firestore:"userUID"`
	PublicKey   string `firestore:"publicKey"`
	DeviceToken string `firestore:"deviceToken"`
}

type ChatPublicKeyData struct {
	ChatID    string `firestore:"chatID"`
	PublicKey string `firestore:"publicKey"`
}

type ChatParticipantPrivateKeyData struct {
	ParticipantUID             string `firestore:"participantUID"`
	DeviceToken                string `firestore:"deviceToken"`
	ChatID                     string `firestore:"chatID"`
	ChatPublicKey              string `firestore:"chatPublicKey"`
	EncodedEncryptedPrivateKey string `firestore:"encodedEncryptedPrivateKey"`
	EncodedEncryptedAESKey     string `firestore:"encodedEncryptedAESKey"`
	EncodedAESIV               string `firestore:"encodedAESIV"`
}

type DevicePublicKey struct {
	DeviceToken string
	PublicKey   rsa.PublicKey
}
