package clients

import (
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
			ServiceHost: serviceHost,
		},
	}
}

func (client *NotificationClient) ConsultationMessageSent(senderUserUID string, consultationID string, messageText string) (*http.Response, error) {
	type consultationMessageSentRequest struct {
		SenderUserUID		string		`json:"senderUserUID"`
		ConsultationID	string		`json:"consultationID"`
		MessageText			string		`json:"messageText"`
	}

	req := consultationMessageSentRequest{
		SenderUserUID: senderUserUID,
		ConsultationID: consultationID,
		MessageText: messageText,
	}

	return client.Post("/notifications/consultation-message-sent", req)
}
