package resources

import "github.com/hellodoctordev/common/emr/codes"

type CarePlan struct {
	// http://hl7.org/implement/standards/fhir/STU3/careplan.html
	BaseResource
	Definition  []Reference          `json:"definition,omitempty"`
	BasedOn     []Reference          `json:"basedOn,omitempty"`
	Replaces    []Reference          `json:"replaces,omitempty"`
	PartOf      []Reference          `json:"partOf,omitempty"`
	Status      codes.CarePlanStatus `json:"status"`
	Intent      codes.CarePlanIntent `json:"intent"`
	Category    []CodeableConcept    `json:"category,omitempty"`
	Title       string               `json:"title,omitempty"`
	Description string               `json:"description,omitempty"`
	Subject     Reference              `json:"subject"`
	Context     *Reference     `json:"context,omitempty"`
	Period      *Period              `json:"period,omitempty"`
	Author      []Reference     `json:"author,omitempty"`
	CareTeam []Reference `json:"careTeam,omitempty"`
	Addresses   []Reference          `json:"addresses,omitempty"`
	Goal        []Reference               `json:"goal,omitempty"`
	Activity    []CarePlanActivity   `json:"activity"`
}

type ProcedureRequest struct {
	// http://hl7.org/implement/standards/fhir/STU3/procedurerequest.html
}

type ReferralRequest struct {
	// http://hl7.org/implement/standards/fhir/STU3/referralrequest.html
}
