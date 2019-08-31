package clients

import (
	"hellodoctor/api"
	"net/http"
)

type MessagingClient struct {
	common.HttpServiceClient
}

func NewMessagingClient() *MessagingClient {
	return &MessagingClient{
		common.HttpServiceClient{
			ServiceHost: common.DefaultServiceHost,
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
