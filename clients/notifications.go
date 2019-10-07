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
	ChatID        string `json:"chatID"`
	SenderUserUID string `json:"senderUserUID"`
	MessageText   string `json:"messageText"`
}

func (client *NotificationClient) ChatMessageSent(chatID, senderUserUID, messageText string) (*http.Response, error) {
	req := ChatMessageSentRequest{
		SenderUserUID: senderUserUID,
		ChatID:        chatID,
		MessageText:   messageText,
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
