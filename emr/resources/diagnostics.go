package resources

import (
	"github.com/hellodoctordev/common/emr/codes"
	"time"
)

type DiagnosticReport struct {
	// http://hl7.org/implement/standards/fhir/STU3/diagnosticreport.html
}

type Observation struct {
	// http://hl7.org/implement/standards/fhir/STU3/observation.html
	BaseResource
	BasedOn           []Reference             `json:"basedOn,omitempty"`
	Status            codes.ObservationStatus `json:"status"`
	Category          []CodeableConcept       `json:"category,omitempty"`
	Code              CodeableConcept         `json:"code"`
	Subject           *Reference              `json:"subject,omitempty"`
	Context           *Reference              `json:"context,omitempty"`
	EffectiveDateTime *time.Time              `json:"effectiveDateTime,omitempty"`
	EffectivePeriod   *Period                 `json:"effectivePeriod,omitempty"`
	Issued            *time.Time              `json:"issued,omitempty"`
	Performer         []Reference             `json:"performer,omitempty"`
	ValueQuantity     *Quantity               `json:"valueQuantity,omitempty"`
	ValueString       string                  `json:"valueString,omitempty"`
	ValueDateTime     *time.Time              `json:"valueDateTime,omitempty"`
	Comment           string                  `json:"comment,omitempty"`
	BodySite          *CodeableConcept        `json:"bodySite,omitempty"`
	Method            *CodeableConcept        `json:"method,omitempty"`
	Specimen          *Reference              `json:"specimen,omitempty"`
}

type BodySite struct {
	Identifier  []Identifier                          `json:"identifier"`
	Active      bool                                  `json:"active"`
	Code        codes.BodyStructure                   `json:"code"`
	Qualifier   []codes.BodySiteLocationQualifierCode `json:"qualifier"`
	Description string                                `json:"description"`
	Image       []Attachment                          `json:"image"`
	Patient     *Patient                              `json:"patient"`
}
