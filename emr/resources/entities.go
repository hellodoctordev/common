package resources

import (
	"github.com/hellodoctordev/common/emr/codes"
)

type Organization struct {
	BaseResource
	Active  bool              `json:"active"`
	Type    []CodeableConcept `json:"type"`
	Name    string            `json:"name"`
	Alias   []string          `json:"alias"`
	Telecom []ContactPoint    `json:"telecom"`
	Address []Address         `json:"address"`
	PartOf  Reference         `json:"partOf,omitempty"`
}

func (o Organization) GetResourceType() string { return "Organization" }
func (o Organization) IsProcedurePerformer()        {}
func (o Organization) IsObservationPerformer()      {}
func (o Organization) IsCarePlanAuthor()            {}
func (o Organization) IsGoalSubject()               {}
func (o Organization) IsCarePlanActivityPerformer() {}

type Location struct {
	// http://hl7.org/implement/standards/fhir/STU3/location.html
	Identifier   []Identifier                          `json:"identifier"`
	Status       codes.LocationStatus                  `json:"status"`
	Name         string                                `json:"name"`
	Mode         codes.LocationMode                    `json:"mode"`
	Type         codes.ServiceDeliveryLocationRoleType `json:"type"`
	Telecom      []ContactPoint                        `json:"telecom"`
	Address      Address                               `json:"address"`
	PhysicalType codes.LocationType                    `json:"physicalType"`
}

func (l Location) IsObservationSubject() {}

type HealthcareService struct {
	Identifier Identifier            `json:"identifier"`
	Active     bool                  `json:"active"`
	Category   codes.ServiceCategory `json:"category"`
	Type       []codes.ServiceType   `json:"type"`
	Specialty  []codes.Specialty     `json:"specialty"`
	Location   []Location            `json:"location"`
	Name       string                `json:"name"`
	Comment    string                `json:"comment"`
	Photo      Attachment            `json:"photo"`
	Telecom    []ContactPoint        `json:"telecom"`
}

func (hs HealthcareService) IsProcedureDefinition() {}

type Substance struct{}

func (s Substance) IsProductReference() {}
