package clients

import (
	"net/http"
)

type MessagingClient struct {
	HttpServiceClient
}

func NewMessagingClient() *MessagingClient {
	return &MessagingClient{
		HttpServiceClient{
			ServiceHost: DefaultServiceHost,
		},
	}
}

func (client *MessagingClient) SendConsultationMessage(senderId string, consultationID string, message string, contentType string, messageType string) (*http.Response, error) {
	type sendConsultationMessageRequest struct {
		SenderID		string		`json:"senderID"`
		ConsultationID	string		`json:"consultationID"`
		Message			string		`json:"message"`
		ContentType		string		`json:"contentType"`
		MessageType		string		`json:"messageType"`
	}

	req := sendConsultationMessageRequest{
		SenderID: senderId,
		ConsultationID: consultationID,
		Message: message,
		ContentType: contentType,
		MessageType: messageType,
	}

	return client.Post("/messages/send", req)
}
