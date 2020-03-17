package crypto

import (
	"bytes"
	"cloud.google.com/go/firestore"
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"github.com/hellodoctordev/common/firebase"
	"github.com/hellodoctordev/common/logging"
	"google.golang.org/api/iterator"
	"io"
	"strings"
)

var firestoreClient = firebase.NewFirestoreClient()

const keyBitSize = 2048

func GenerateChatKeys(chatID string, participantRefs []*firestore.DocumentRef) {
	ctx := context.Background()

	reader := rand.Reader

	chatKey, err := rsa.GenerateKey(reader, keyBitSize)
	if err != nil {
		logging.Error("error generating keys for chat %s: %s", chatID, err)
		return
	}

	publicPem := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&chatKey.PublicKey),
	}

	var chatPublicKey strings.Builder
	err = pem.Encode(&chatPublicKey, &publicPem)
	if err != nil {
		logging.Error("error creating chat public key string: %s", err)
		return
	}

	chatPublicKeyData := ChatPublicKeyData{
		ChatID:    chatID,
		PublicKey: chatPublicKey.String(),
	}

	_, _, err = firestoreClient.Collection("publicKeys").Add(ctx, chatPublicKeyData)
	if err != nil {
		logging.Error("error storing public chatKey for chat %s: %s", chatID, err)
		return
	}

	chatPrivateKeyBytes, err := x509.MarshalPKCS8PrivateKey(chatKey)
	if err != nil {
		logging.Error("error marshaling chat %s private key: %s", chatID, err)
		return
	}

	chatAESKey, err := generateNewAESKey()
	if err != nil {
		logging.Error("error generating new AES key: %s", err)
		return
	}

	encryptedChatPrivateKey, aesIV, err := encryptChatPrivateKey(chatAESKey, chatPrivateKeyBytes)
	if err != nil {
		logging.Error("error encrypting private key: %s", err)
		return
	}

	for _, participantRef := range participantRefs {
		participantPublicKeys, err2 := getParticipantDevicesPublicKeys(participantRef.ID)
		if err2 != nil {
			continue
		}

		for _, participantDevicePublicKey := range participantPublicKeys {
			//encryptedChatAESKeyBytes, err2 := rsa.EncryptPKCS1v15(reader, &participantDevicePublicKey.PublicKey, chatAESKey)
			encryptedChatAESKeyBytes, err2 := rsa.EncryptOAEP(sha256.New(), reader, &participantDevicePublicKey.PublicKey, chatAESKey, nil)
			if err2 != nil {
				logging.Warn("error occurred encrypting chat %s private key for participant %s: %s", chatID, participantRef.ID, err2)
				continue
			}

			chatParticipantPrivateKey := ChatParticipantPrivateKeyData{
				ParticipantUID:             participantRef.ID,
				DeviceToken:                participantDevicePublicKey.DeviceToken,
				ChatID:                     chatID,
				ChatPublicKey:              chatPublicKey.String(),
				EncodedEncryptedPrivateKey: hex.EncodeToString(encryptedChatPrivateKey),
				EncodedEncryptedAESKey:     hex.EncodeToString(encryptedChatAESKeyBytes),
				EncodedAESIV:               hex.EncodeToString(aesIV),
			}

			_, _, err2 = firestoreClient.Collection("encryptedPrivateKeys").Add(ctx, chatParticipantPrivateKey)
			if err2 != nil {
				logging.Warn("error occurred storing chat %s private key for participant %s: %s", chatID, participantRef.ID, err2)
				continue
			}
		}
	}
}

func getParticipantDevicesPublicKeys(participantUID string) (participantPublicKeys []DevicePublicKey, err error) {
	participantDeviceSnapshots := firestoreClient.Collection("users").
		Doc(participantUID).
		Collection("devices").
		Documents(context.Background())

	for {
		participantDeviceSnapshot, err2 := participantDeviceSnapshots.Next()
		if err2 == iterator.Done {
			return
		} else if err2 != nil {
			logging.Warn("error occurred getting participant %s device: %s", participantUID, err2)
			continue
		}

		var participantDeviceData UserDeviceData
		err2 = participantDeviceSnapshot.DataTo(&participantDeviceData)
		if err2 != nil {
			logging.Warn("error occurred getting participant %s public chatKey data: %s", participantUID, err2)
			continue
		}

		if !participantDeviceData.IsMasterDevice && participantDeviceData.AuthorizedBy == nil {
			continue
		}

		block, _ := pem.Decode([]byte(participantDeviceData.PublicKey))

		var participantDevicePublicKey rsa.PublicKey
		_, err2 = asn1.Unmarshal(block.Bytes, &participantDevicePublicKey)
		if err2 != nil {
			logging.Warn("error occurred parsing participant %s public key: %s", participantUID, err2)
			continue
		}

		devicePublicKey := DevicePublicKey{
			DeviceToken: participantDeviceData.DeviceToken,
			PublicKey:   participantDevicePublicKey,
		}

		participantPublicKeys = append(participantPublicKeys, devicePublicKey)
	}
}

func generateNewAESKey() ([]byte, error) {
	key := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		logging.Error(err.Error())
		return nil, err
	}

	return key, nil
}

func encryptChatPrivateKey(aesKey []byte, chatPrivateKeyBytes []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(aesKey)
	if err != nil {
		logging.Error(err.Error())
		return nil, nil, err
	}

	paddedChatPrivateKeyBytes := pad(chatPrivateKeyBytes)
	encryptedChatPrivateKeyBytes := make([]byte, aes.BlockSize)

	iv := encryptedChatPrivateKeyBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, err
	}

	//gcm, err := cipher.NewGCM(block)
	//if err != nil {
	//	logging.Error("error creating new gcm: %s", err)
	//	return nil, nil, err
	//}


	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		logging.Error("error encrypting chat private key: %s", err)
		return nil, nil, err
	}

	cfb := cipher.NewCBCEncrypter(block, iv)
	cfb.CryptBlocks(encryptedChatPrivateKeyBytes[aes.BlockSize:], paddedChatPrivateKeyBytes)

	//encryptedChatPrivateKeyBytes := gcm.Seal(nil, nonce, chatPrivateKeyBytes, nil)

	return encryptedChatPrivateKeyBytes, iv, nil
}

func pad(aesKey []byte) []byte {
	padLength := aes.BlockSize - len(aesKey) % aes.BlockSize
	padBytes := bytes.Repeat([]byte{byte(padLength)}, padLength)

	return append(aesKey, padBytes...)
}
