package clients

import (
	"github.com/hellodoctordev/common/utils"
	"net/http"
	"os"
	"time"
)

type SchedulingClient struct {
	HttpServiceClient
}

func NewSchedulingClient() *MessagingClient {
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

type CreateNewEventRequest struct {
	Title        string    `json:"title"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	Availability string    `json:"availability"`
	EventType    string    `json:"eventType"`
	IsAllDay     bool      `json:"isAllDay"`
	Participants []string  `json:"participants"`
	Description  string    `json:"description"`
}


type CreateNewEventResponse struct {
	EventID string `json:"eventID"`
}

func (client *SchedulingClient) CreateNewEvent(req CreateNewEventRequest) (res CreateNewEventResponse, err error) {
	r, err := client.Post("/scheduling/internal/events", req)
	if err == nil {
		err = utils.ReadBody(r.Body, &res)
	}

	return
}
