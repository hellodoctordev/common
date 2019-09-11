package types

import (
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"time"
)

type DoctorESDocument struct {
	UID                string           `json:"uid"`
	DisplayName        string           `json:"displayName"`
	FirstName          string           `json:"firstName"`
	LastName           string           `json:"lastName"`
	Specialties        []string         `json:"specialties"`
	OfficeAddress      string           `json:"officeAddress"`
	OfficePostalCode   string           `json:"officePostalCode"`
	OfficeNeighborhood string           `json:"officeNeighborhood"`
	OfficeCity         string           `json:"officeCity"`
	OfficeState        string           `json:"officeState"`
	OfficeCountry      string           `json:"officeCountry"`
	OfficeGeoLocation  *elastic.GeoPoint `json:"officeGeoLocation"`
}

func FromUserData(uid string, userData map[string]interface{}) (doc *DoctorESDocument, err error) {
	role := userData["account"].(map[string]interface{})["role"]

	if role != "provider" {
		err = errors.New(fmt.Sprintf("invalid role '%s' for doctor document", role))
		return
	}

	profileData := userData["profile"].(map[string]interface{})
	officeData := profileData["office"].(map[string]interface{})

	doc = &DoctorESDocument{}
	doc.UID = uid
	doc.DisplayName = profileData["displayName"].(string)
	doc.FirstName = profileData["firstName"].(string)
	doc.LastName = profileData["lastName"].(string)
	doc.Specialties = profileData["specialties"].([]string)
	doc.OfficeAddress = officeData["address"].(string)
	doc.OfficePostalCode = officeData["postalCode"].(string)
	doc.OfficeNeighborhood = officeData["neighborhood"].(string)
	doc.OfficeCity = officeData["city"].(string)
	doc.OfficeState = officeData["state"].(string)
	doc.OfficeCountry = officeData["country"].(string)

	geoData := officeData["geoLocation"].(map[string]float64)
	doc.OfficeGeoLocation = elastic.GeoPointFromLatLon(geoData["lat"], geoData["lon"])

	return
}

type PatientESDocument struct {
	UID       string `json:"uid"`
	FullName  string `json:"fullName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type MessageESDocument struct {
	UID         string    `json:"uid"`
	Message     string    `json:"message"`
	MessageType string    `json:"messageType"`
	Sender      string    `json:"sender"`
	SentTime    time.Time `json:"sentTime"`
}
