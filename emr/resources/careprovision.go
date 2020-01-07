package resources

import "hellodoctor/emr/x/codes"

type CarePlan struct {
	Identifier  []Identifier                 `json:"identifier"`
	Definition  []CarePlanDefinition         `json:"definition"`
	BasedOn     []CarePlan                   `json:"basedOn"`
	Replaces    []CarePlan                   `json:"replaces"`
	PartOf      []CarePlan                   `json:"partOf"`
	Status      codes.CarePlanStatus         `json:"status"`
	Intent      codes.CarePlanIntent         `json:"intent"`
	Category    []codes.CarePlanCategoryCode `json:"category"`
	Title       string                       `json:"title"`
	Description string                       `json:"description"`
	Subject     Patient                      `json:"subject"`
	Context     *CarePlanContext             `json:"context"`
	Period      *Period                      `json:"period"`
	Author      []CarePlanAuthor             `json:"author"`
	Addresses   []Condition                  `json:"addresses"`
	Goal        []Goal                       `json:"goal"`
	Activity    []CarePlanActivity           `json:"activity"`
}

func (c CarePlan) IsObservationBasedOnReference() {}

type ProcedureRequest struct {
	// http://hl7.org/implement/standards/fhir/STU3/procedurerequest.html
}

type ReferralRequest struct {
	// http://hl7.org/implement/standards/fhir/STU3/referralrequest.html
}
