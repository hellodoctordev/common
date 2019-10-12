package clients

import (
	"cloud.google.com/go/firestore"
	"github.com/hellodoctordev/common/utils"
	"net/http"
	"os"
	"time"
)

type SchedulingClient struct {
	HttpServiceClient
}

func NewSchedulingClient() *SchedulingClient {
	serviceHost := os.Getenv("BASE_SERVICE_URL")

	if len(serviceHost) == 0 {
		serviceHost = DefaultServiceHost
	}

	return &SchedulingClient{
		HttpServiceClient{
			Client:      http.DefaultClient,
			ServiceHost: serviceHost,
		},
	}
}

type InternalCreateNewEventRequest struct {
	UserUID      string                   `json:"userUID"`
	Title        string                   `json:"title"`
	Start        time.Time                `json:"start"`
	End          time.Time                `json:"end"`
	Availability string                   `json:"availability"`
	EventType    string                   `json:"eventType"`
	IsAllDay     bool                     `json:"isAllDay"`
	Participants []*firestore.DocumentRef `json:"participants"`
	Description  string                   `json:"description"`
}

type InternalCreateNewEventResponse struct {
	Event *firestore.DocumentRef `json:"eventRef"`
}

func (client *SchedulingClient) InternalCreateNewEvent(req InternalCreateNewEventRequest) (res InternalCreateNewEventResponse, err error) {
	r, err := client.Post("/scheduling/internal/events", req)
	if err == nil {
		err = utils.ReadBody(r.Body, &res)
	}

	return
}
