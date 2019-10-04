package clients

import (
	"net/http"
	"os"
)

type MessagingClient struct {
	HttpServiceClient
}

func NewMessagingClient() *MessagingClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &MessagingClient{
		HttpServiceClient{
			Client: http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateChatRequest struct {
	PatientUserUID    string `json:"patientUserUID"`
	DoctorUserUID string `json:"doctorUserUID"`
}

type CreateChatResponse struct {
	ChatID string `json:"chatID"`
}

func (client *MessagingClient) CreateChat(patientUID, doctorUserUID string) (*http.Response, error) {
	req := CreateChatRequest{
		PatientUserUID:    patientUID,
		DoctorUserUID: doctorUserUID,
	}

	return client.Post("/messaging/internal/chats", req)
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
