package crypto

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/asn1"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/hellodoctordev/common/firebase"
	"github.com/hellodoctordev/common/logging"
	"google.golang.org/api/iterator"
	"io"
)

var firestoreClient = firebase.NewFirestoreClient()

func GenerateChatKey(chatID string) {
	ctx := context.Background()

	chatSnapshot, err := firestoreClient.Doc(fmt.Sprintf("chats/%s", chatID)).Get(ctx)
	if err != nil {
		logging.Error("couldn't get chat %s", chatID)
		return
	}

	chatAESKey, err := generateNewAESKey()
	if err != nil {
		logging.Error("error generating new AES key: %s", err)
		return
	}

	participantsData, err := chatSnapshot.DataAt("participants")
	if err == nil && participantsData != nil {
		for _, participantData := range participantsData.([]interface{}) {
			participant := participantData.(*firestore.DocumentRef)
			registerParticipantChatKeys(ctx, chatSnapshot.Ref, participant.ID, chatAESKey)
		}
	} else {
		practitionerRefData, _ := chatSnapshot.DataAt("practitioner")
		practitionerRef := practitionerRefData.(*firestore.DocumentRef)

		patientRefData, _ := chatSnapshot.DataAt("patient")
		patientRef := patientRefData.(*firestore.DocumentRef)

		registerParticipantChatKeys(ctx, chatSnapshot.Ref, practitionerRef.ID, chatAESKey)
		registerParticipantChatKeys(ctx, chatSnapshot.Ref, patientRef.ID, chatAESKey)
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

		if !participantDeviceData.IsMasterDevice && participantDeviceData.AuthorizedBy == nil && !participantDeviceData.IsVerified {
			continue
		}

		block, _ := pem.Decode([]byte(participantDeviceData.PublicKey))
		if block == nil {
			logging.Warn("could not decode participant %s device %s public key", participantUID, participantDeviceSnapshot.Ref.ID)
			continue
		}

		var participantDevicePublicKey rsa.PublicKey
		_, err2 = asn1.Unmarshal(block.Bytes, &participantDevicePublicKey)
		if err2 != nil {
			logging.Warn("error occurred parsing participant %s public key: %s", participantUID, err2)
			continue
		}

		devicePublicKey := DevicePublicKey{
			DeviceID:    participantDeviceSnapshot.Ref.ID,
			DeviceToken: participantDeviceData.DeviceToken,
			PublicKey:   participantDevicePublicKey,
		}

		participantPublicKeys = append(participantPublicKeys, devicePublicKey)
	}
}

func registerParticipantChatKeys(ctx context.Context, chatRef *firestore.DocumentRef, participantID string, chatAESKey []byte) {
	participantDevicePublicKeys, err2 := getParticipantDevicesPublicKeys(participantID)
	if err2 != nil {
		logging.Error("failed to get participant %s public keys: %s", participantID, err2)
		return
	} else if len(participantDevicePublicKeys) == 0 {
		logging.Info("no keys found for chat %s participant %s", chatRef.ID, participantID)

		userChatRef := firestoreClient.Collection("users").Doc(participantID).Collection("chats").Doc(chatRef.ID)
		userChatData := map[string]interface{}{
			"chat": chatRef,
			"key": base64.StdEncoding.EncodeToString(chatAESKey),
		}

		userChatRef.Set(ctx, userChatData)

		return
	}

	reader := rand.Reader

	for _, publicKey := range participantDevicePublicKeys {
		encryptedChatAESKeyBytes, err2 := rsa.EncryptOAEP(sha1.New(), reader, &publicKey.PublicKey, chatAESKey, nil)
		if err2 != nil {
			logging.Warn("error occurred encrypting chat %s private key for participant %s: %s", chatRef.ID, participantID, err2)
			continue
		}

		_, _ = chatRef.Update(ctx, []firestore.Update{{
			Path:  fmt.Sprintf("device-key-%s", publicKey.DeviceID),
			Value: base64.StdEncoding.EncodeToString(encryptedChatAESKeyBytes),
		}})
	}
}

func getDemoUserAESKey() []byte {
	return []byte("cec0f2ad51a3b727444d107cf7f71072")
}

func generateNewAESKey() ([]byte, error) {
	key := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		logging.Error(err.Error())
		return nil, err
	}

	return key, nil
}
