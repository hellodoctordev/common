package clients

import (
	"fmt"
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
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type CreateChatRequest struct {
	PatientUserUID  string `json:"patientUserUID"`
	ProviderUserUID string `json:"providerUserUID"`
}

type CreateChatResponse struct {
	ChatID string `json:"chatID"`
}

func (client *MessagingClient) CreateChat(patientUID, providerUserUID string) (*http.Response, error) {
	req := CreateChatRequest{
		PatientUserUID:  patientUID,
		ProviderUserUID: providerUserUID,
	}

	return client.Post("/messaging/internal/chats", req)
}

type SendChatMessageRequest struct {
	ChatID      string `json:"chatID"`
	SenderUID   string `json:"senderUID"`
	Message     string `json:"message"`
	ContentType string `json:"contentType"`
	MessageType string `json:"messageType"`
}

func (client *MessagingClient) SendChatMessage(chatID, senderUID, message, contentType, messageType string) (*http.Response, error) {
	req := SendChatMessageRequest{
		SenderUID:   senderUID,
		Message:     message,
		ContentType: contentType,
		MessageType: messageType,
	}

	return client.Post(fmt.Sprintf("/messaging/internal/chats/%s", chatID), req)
}

func (client *MessagingClient) SendConsultationMessage(senderId string, consultationID string, message string, contentType string, messageType string) (*http.Response, error) {
	type sendConsultationMessageRequest struct {
		SenderID       string `json:"senderID"`
		ConsultationID string `json:"consultationID"`
		Message        string `json:"message"`
		ContentType    string `json:"contentType"`
		MessageType    string `json:"messageType"`
	}

	req := sendConsultationMessageRequest{
		SenderID:       senderId,
		ConsultationID: consultationID,
		Message:        message,
		ContentType:    contentType,
		MessageType:    messageType,
	}

	return client.Post("/messages/send", req)
}
