package clients

import (
	"cloud.google.com/go/firestore"
	"net/http"
	"os"
	"time"
)

type NotificationClient struct {
	HttpServiceClient
}

func NewNotificationClient() *NotificationClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &NotificationClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type ChatMessageSentRequest struct {
	ChatID               string `json:"chatID"`
	ConsultationID       string `json:"consultationID"`
	SenderUserUID        string `json:"senderUserUID"`
	EncryptedMessageText string `json:"encryptedMessageText"`
}

func (client *NotificationClient) ChatMessageSent(chatID, consultationID, senderUserUID, encryptedMessageText string) (*http.Response, error) {
	req := ChatMessageSentRequest{
		SenderUserUID:        senderUserUID,
		ChatID:               chatID,
		ConsultationID:       consultationID,
		EncryptedMessageText: encryptedMessageText,
	}

	return client.Post("/notifications/chat-message-sent", req)
}

type ConsultationRequestedRequest struct {
	Chat               *firestore.DocumentRef `json:"chatRef"`
	Consultation       *firestore.DocumentRef `json:"consultationRef"`
	PatientUser        *firestore.DocumentRef `json:"patientUserRef"`
	ConsultationType   string                 `json:"consultationType"`
	RequestedStartTime time.Time              `json:"requestedStartTime"`
	RequestedEndTime   time.Time              `json:"requestedEndTime"`
}

func (client *NotificationClient) ConsultationRequested(req ConsultationRequestedRequest) (*http.Response, error) {
	return client.Post("/notifications/consultation-requested", req)
}

type VideoConsultationStartedRequest struct {
	InitiatedByUserUID string `json:"initiatedByUserUID"`
	ConsultationID     string `json:"consultationID"`
	RoomName           string `json:"RoomName"`
}

func (client *NotificationClient) VideoConsultationStarted(initiatedByUserUID string, consultationID string, roomName string) (*http.Response, error) {
	req := VideoConsultationStartedRequest{
		InitiatedByUserUID: initiatedByUserUID,
		ConsultationID:     consultationID,
		RoomName:           roomName,
	}

	return client.Post("/notifications/video-consultation-started", req)
}

type VoiceConsultationStartedRequest struct {
	InitiatedByUserUID string `json:"initiatedByUserUID"`
	ConsultationID     string `json:"consultationID"`
}

func (client *NotificationClient) VoiceConsultationStarted(initiatedByUserUID string, consultationID string) (*http.Response, error) {
	req := VoiceConsultationStartedRequest{
		InitiatedByUserUID: initiatedByUserUID,
		ConsultationID:     consultationID,
	}

	return client.Post("/notifications/voice-consultation-started", req)
}
