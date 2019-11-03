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

type SendAutomatedChatMessageRequest struct {
	ConsultationID string `json:"consultationID"`
	MessageContent string `json:"messageContent"`
	ContentType    string `json:"contentType"`
}

func (client *MessagingClient) SendAutomatedChatMessage(chatID, consultationID, messageContent, contentType, messageType string) (*http.Response, error) {
	req := SendAutomatedChatMessageRequest{
		ConsultationID: consultationID,
		MessageContent: messageContent,
		ContentType:    contentType,
	}

	return client.Post(fmt.Sprintf("/messaging/internal/chats/%s", chatID), req)
}

type SendChatMessageRequest struct {
	ChatID         string `json:"chatID"`
	ConsultationID string `json:"consultationID"`
	SenderID       string `json:"senderID"`
	Message        string `json:"message"`
	ContentType    string `json:"contentType"`
}

func (client *MessagingClient) SendChatMessage(chatID, consultationID, senderID, message, contentType, messageType string) (*http.Response, error) {
	req := SendChatMessageRequest{
		SenderID:       senderID,
		ConsultationID: consultationID,
		Message:        message,
		ContentType:    contentType,
	}

	return client.Post(fmt.Sprintf("/messaging/chats/%s/messages", chatID), req)
}
