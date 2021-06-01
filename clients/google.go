package clients

import (
	"github.com/hellodoctordev/common/keys"
	"googlemaps.github.io/maps"
)

func NewGoogleMapsClient() (*maps.Client, error) {
	serverKey := keys.GoogleApiKeys.ServerKey

	return maps.NewClient(maps.WithAPIKey(serverKey))
}
