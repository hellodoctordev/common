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

type ConsultationMessageSentRequest struct {
	SenderUserUID  string `json:"senderUserUID"`
	ConsultationID string `json:"consultationID"`
	MessageText    string `json:"messageText"`
}

func (client *NotificationClient) ConsultationMessageSent(senderUserUID string, consultationID string, messageText string) (*http.Response, error) {
	req := ConsultationMessageSentRequest{
		SenderUserUID:  senderUserUID,
		ConsultationID: consultationID,
		MessageText:    messageText,
	}

	return client.Post("/notifications/consultation-message-sent", req)
}

type ConsultationSessionRequestedRequest struct {
	Consultation       *firestore.DocumentRef `json:"consultationRef"`
	PatientUser        *firestore.DocumentRef `json:"patientUserRef"`
	ConsultationType   string                 `json:"consultationType"`
	RequestedStartTime time.Time              `json:"requestedStartTime"`
	RequestedEndTime   time.Time              `json:"requestedEndTime"`
}

func (client *NotificationClient) ConsultationSessionRequested(req ConsultationSessionRequestedRequest) (*http.Response, error) {
	return client.Post("/notifications/consultation-requested", req)
}

type VideoConsultationStartedRequest struct {
	InitiatedByUserUID string `json:"initiatedByUserUID"`
	ConsultationID     string `json:"consultationID"`
	RoomSID            string `json:"roomSID"`
}

func (client *NotificationClient) VideoConsultationStarted(initiatedByUserUID string, consultationID string, roomSID string) (*http.Response, error) {
	req := VideoConsultationStartedRequest{
		InitiatedByUserUID: initiatedByUserUID,
		ConsultationID:     consultationID,
		RoomSID:            roomSID,
	}

	return client.Post("/notifications/video-consultation-started", req)
}
