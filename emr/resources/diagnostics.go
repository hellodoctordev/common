package resources

import (
	"hellodoctor/emr/x/codes"
	"time"
)

type DiagnosticReport struct {
	// http://hl7.org/implement/standards/fhir/STU3/diagnosticreport.html
}

func (dr DiagnosticReport) IsConditionStageAssessment() {}
func (dr DiagnosticReport) IsInvestigationItem()        {}

type Observation struct {
	Identifier        []Identifier                    `json:"identifier"`
	BasedOn           []ObservationBasedOnReference   `json:"basedOn"`
	Status            codes.ObservationStatus         `json:"status"`
	Category          []codes.ObservationCategoryCode `json:"category"`
	Code              codes.ObservationCode           `json:"code"`
	Subject           ObservationSubject              `json:"subject"`
	Context           ObservationContext              `json:"context"`
	EffectiveDateTime *time.Time                      `json:"effectiveDateTime"`
	EffectivePeriod   *Period                         `json:"effectivePeriod"`
	Issued            *time.Time                      `json:"issued"`
	Performer         []ObservationPerformer          `json:"performer"`
	ValueQuantity     *Quantity                       `json:"valueQuantity"`
	ValueString       string                          `json:"valueString"`
	ValueDateTime     *time.Time                      `json:"valueDateTime"`
	Comment           string                          `json:"comment"`
	BodySite          *codes.BodyStructure            `json:"bodySite"`
	Method            *codes.ObservationMethodCode    `json:"method"`
}

func (o Observation) IsValid() bool {
	effectiveStarIsValid := o.EffectiveDateTime == nil || o.EffectivePeriod == nil
	valueStarIsValid := o.ValueQuantity == nil && o.ValueString == "" || o.ValueString == "" && o.ValueDateTime == nil

	return effectiveStarIsValid && valueStarIsValid
}
func (o Observation) IsConditionStageAssessment()     {}
func (o Observation) IsProcedureReasonReference()     {}
func (o Observation) IsFamilyHistoryReasonReference() {}
func (o Observation) IsFindingItemReference()         {}
func (o Observation) IsGoalAddresses()                {}

type BodySite struct {
	Identifier  []Identifier                          `json:"identifier"`
	Active      bool                                  `json:"active"`
	Code        codes.BodyStructure                   `json:"code"`
	Qualifier   []codes.BodySiteLocationQualifierCode `json:"qualifier"`
	Description string                                `json:"description"`
	Image       []Attachment                          `json:"image"`
	Patient     *Patient                              `json:"patient"`
}
