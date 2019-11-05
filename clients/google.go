package clients

import (
	"github.com/hellodoctordev/common/logging"
	"googlemaps.github.io/maps"
)

func NewGoogleMapsClient() *maps.Client {
	client, err := maps.NewClient(maps.WithAPIKey("Insert-API-Key-Here"))
	if err != nil {
		logging.Error("could not create google maps api client: %s", err)
		return nil
	}

	return client
}
