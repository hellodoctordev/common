package crypto

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"github.com/hellodoctordev/common/firebase"
	"github.com/hellodoctordev/common/logging"
	"io"
	"strings"
)

var firestoreClient = firebase.NewFirestoreClient()

const keyBitSize = 2048

type UserPublicKeyData struct {
	UserUID   string `firestore:"userUID"`
	PublicKey string `firestore:"publicKey"`
}

type ConsultationPublicKeyData struct {
	ConsultationID string `firestore:"consultationID"`
	PublicKey      string `firestore:"publicKey"`
}

type ConsultationParticipantPrivateKeyData struct {
	ParticipantUID             string `firestore:"participantUID"`
	ConsultationID             string `firestore:"consultationID"`
	ConsultationPublicKey      string `firestore:"consultationPublicKey"`
	EncodedEncryptedPrivateKey string `firestore:"encodedEncryptedPrivateKey"`
	EncodedEncryptedAESKey     string `firestore:"encodedEncryptedAESKey"`
	EncodedAESIV               string `firestore:"encodedAESIV"`
}

func GenerateConsultationKeys(consultationID string, participantRefs []*firestore.DocumentRef) {
	ctx := context.Background()

	reader := rand.Reader

	consultationKey, err := rsa.GenerateKey(reader, keyBitSize)
	if err != nil {
		logging.Error("error generating keys for consultation %s: %s", consultationID, err)
		return
	}

	publicPem := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&consultationKey.PublicKey),
	}

	var consultationPublicKey strings.Builder

	err = pem.Encode(&consultationPublicKey, &publicPem)
	if err != nil {
		logging.Error("error creating consultation public key string: %s", err)
		return
	}

	consultationPublicKeyData := ConsultationPublicKeyData{
		ConsultationID: consultationID,
		PublicKey:      consultationPublicKey.String(),
	}

	_, _, err = firestoreClient.Collection("publicKeys").Add(ctx, consultationPublicKeyData)
	if err != nil {
		logging.Error("error storing public consultationKey for consultation %s: %s", consultationID, err)
		return
	}

	consultationPrivateKeyBytes, err := x509.MarshalPKCS8PrivateKey(consultationKey)
	if err != nil {
		logging.Error("error marshaling consultation %s private key: %s", consultationID, err)
		return
	}

	consultationAESKey, err := generateNewAESKey()
	if err != nil {
		logging.Error("error generating new AES key: %s", err)
		return
	}

	encryptedConsultationPrivateKey, aesIV, err := encryptConsultationPrivateKey(consultationAESKey, consultationPrivateKeyBytes)
	if err != nil {
		logging.Error("error encrypting private key: %s", err)
		return
	}

	for _, participantRef := range participantRefs {
		participantPublicKey, err2 := getParticipantPublicKey(participantRef.ID)
		if err2 != nil {
			continue
		}

		encryptedConsultationAESKeyBytes, err2 := rsa.EncryptPKCS1v15(reader, &participantPublicKey, consultationAESKey)
		if err2 != nil {
			logging.Warn("error occurred encrypting consultation %s private key for participant %s: %s", consultationID, participantRef.ID, err2)
			continue
		}

		consultationParticipantPrivateKey := ConsultationParticipantPrivateKeyData{
			ParticipantUID:             participantRef.ID,
			ConsultationID:             consultationID,
			ConsultationPublicKey:      consultationPublicKey.String(),
			EncodedEncryptedPrivateKey: hex.EncodeToString(encryptedConsultationPrivateKey),
			EncodedEncryptedAESKey:     hex.EncodeToString(encryptedConsultationAESKeyBytes),
			EncodedAESIV:               hex.EncodeToString(aesIV),
		}

		_, _, err2 = firestoreClient.Collection("encryptedPrivateKeys").Add(ctx, consultationParticipantPrivateKey)
		if err2 != nil {
			logging.Warn("error occurred storing consultation %s private key for participant %s: %s", consultationID, participantRef.ID, err2)
			continue
		}
	}
}

func getParticipantPublicKey(participantUID string) (participantPublicKey rsa.PublicKey, err error) {
	participantPublicKeySnapshot, err := firestoreClient.Collection("publicKeys").
		Where("userUID", "==", participantUID).
		Documents(context.Background()).
		Next()

	if err != nil {
		logging.Warn("error occurred getting participant %s public consultationKey: %s", participantUID, err)
		return
	}

	var participantPublicKeyData UserPublicKeyData

	err = participantPublicKeySnapshot.DataTo(&participantPublicKeyData)
	if err != nil {
		logging.Warn("error occurred getting participant %s public consultationKey data: %s", participantUID, err)
		return
	}

	block, _ := pem.Decode([]byte(participantPublicKeyData.PublicKey))

	_, err = asn1.Unmarshal(block.Bytes, &participantPublicKey)
	if err != nil {
		logging.Warn("error occurred parsing participant %s public key: %s", participantUID, err)
		return
	}

	return
}

func generateNewAESKey() ([]byte, error) {
	key := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		logging.Error(err.Error())
		return nil, err
	}

	return key, nil
}

func encryptConsultationPrivateKey(aesKey []byte, consultationPrivateKeyBytes []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		logging.Error(err.Error())
		return nil, nil, err
	}

	paddedConsultationPrivateKeyBytes := pad(consultationPrivateKeyBytes)
	encryptedConsultationPrivateKeyBytes := make([]byte, aes.BlockSize+len(paddedConsultationPrivateKeyBytes))

	iv := encryptedConsultationPrivateKeyBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	cfb := cipher.NewCBCEncrypter(block, iv)
	cfb.CryptBlocks(encryptedConsultationPrivateKeyBytes[aes.BlockSize:], paddedConsultationPrivateKeyBytes)

	return encryptedConsultationPrivateKeyBytes, iv, nil
}

func pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize
	padBytes := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(src, padBytes...)
}
