package clients

import (
	"cloud.google.com/go/firestore"
	"net/http"
	"os"
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
	ChatID             string `json:"chatID"`
	ConsultationID     string `json:"consultationID"`
	SenderUserUID      string `json:"senderUserUID"`
	MessageContent     string `json:"messageContent"`
	MessageContentType string `json:"messageContentType"`
}

func (client *NotificationClient) ChatMessageSent(req ChatMessageSentRequest) (*http.Response, error) {
	return client.Post("/notifications/chat-message-sent", req)
}

type ConsultationNotificationPayload struct {
	ConsultationID string `json:"consultationID"`
}

func (client *NotificationClient) ConsultationRequested(consultationID string) (*http.Response, error) {
	req := ConsultationNotificationPayload{
		ConsultationID: consultationID,
	}

	return client.Post("/notifications/consultation-requested", req)
}

func (client *NotificationClient) ConsultationRescheduled(consultationID string) (*http.Response, error) {
	req := ConsultationNotificationPayload{
		ConsultationID: consultationID,
	}

	return client.Post("/notifications/consultation-rescheduled", req)
}

func (client *NotificationClient) ConsultationCancelled(consultationID string) (*http.Response, error) {
	req := ConsultationNotificationPayload{
		ConsultationID: consultationID,
	}

	return client.Post("/notifications/consultation-cancelled", req)
}

func (client *NotificationClient) ConsultationCompleted(consultationID string) (*http.Response, error) {
	req := ConsultationNotificationPayload{
		ConsultationID: consultationID,
	}

	return client.Post("/notifications/consultation-completed", req)
}

type VideoConsultationStartedRequest struct {
	InitiatedByUserUID string `json:"initiatedByUserUID"`
	ConsultationID     string `json:"consultationID"`
	RoomName           string `json:"roomName"`
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

type VoiceCallStatusRequest struct {
	ConsultationID string `json:"consultationID"`
	CallStatus     string `json:"callStatus"`
}

func (client *NotificationClient) VoiceCallStatus(req VoiceCallStatusRequest) (*http.Response, error) {
	return client.Post("/notifications/voice-call-status", req)
}

type RequestConsultationFollowupRequest struct {
	ConsultationID string `json:"consultationID"`
}

func (client *NotificationClient) RequestConsultationFollowup(req RequestConsultationFollowupRequest) (*http.Response, error) {
	return client.Post("/notifications/request-consultation-followup", req)
}

type ConsultationReminderRequest struct {
	Target       *firestore.DocumentRef `json:"target"`
	Consultation *firestore.DocumentRef `json:"consultation"`
}

func (client *NotificationClient) ConsultationReminder(req ConsultationReminderRequest) (*http.Response, error) {
	return client.Post("/notifications/consultation-reminder", req)
}
