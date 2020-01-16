package resources

import (
	"github.com/hellodoctordev/common/emr/codes"
	"time"
)

type AllergyIntolerance struct {
	Identifier         []Identifier                               `json:"identifier"`
	ClinicalStatus     codes.AllergyIntoleranceClinicalStatus     `json:"clinicalStatus"`
	VerificationStatus codes.AllergyIntoleranceVerificationStatus `json:"verificationStatus"`
	Type               codes.AllergyIntoleranceType               `json:"type"`
	Category           []codes.AllergyIntoleranceCategory         `json:"category"`
	Criticality        codes.AllergyIntoleranceCriticality        `json:"criticality"`
	Code               codes.AllergyIntoleranceCode               `json:"code"`
	Patient            Patient                                    `json:"patient"`
	Note               []string                                   `json:"note"`
	Reaction           []AllergyIntoleranceReaction               `json:"reaction"`
}

type Condition struct {
	Identifier         []Identifier                      `json:"identifier"`
	ClinicalStatus     codes.ConditionClinicalStatus     `json:"clinicalStatus,omitempty"`
	VerificationStatus codes.ConditionVerificationStatus `json:"verificationStatus,omitempty"`
	Category           []CodeableConcept                 `json:"category,omitempty"`
	Severity           CodeableConcept                   `json:"severity,omitempty"`
	Code               CodeableConcept                   `json:"code,omitempty"`
	BodySite           []CodeableConcept                 `json:"bodySite,omitempty"`
	Subject            Reference                         `json:"subject"`
	Context            *Reference                        `json:"context,omitempty"`
	OnsetDateTime      *time.Time                        `json:"onsetDateTime,omitempty"`
	OnsetPeriod        *Period                           `json:"onsetPeriod,omitempty"`
	OnsetString        string                            `json:"onsetString,omitempty"`
	AbatementDateTime  *time.Time                        `json:"abatement,omitempty"`
	AbatementBoolean   bool                              `json:"abatementBoolean,omitempty"`
	AbatementPeriod    *Period                           `json:"abatementPeriod,omitempty"`
	AbatementString    string                            `json:"abatementString,omitempty"`
	AssertedDate       *time.Time                        `json:"assertedDate,omitempty"`
	Asserter           *Reference                        `json:"asserter,omitempty"`
	Stage              *ConditionStage                   `json:"stage,omitempty"`
	Evidence           []ConditionEvidence               `json:"evidence,omitempty"`
	Note               []Annotation                      `json:"note,omitempty"`
}

type Procedure struct {
	// http://hl7.org/implement/standards/fhir/STU3/procedure.html
	BaseResource
	Definition         []Reference          `json:"definition,omitempty"`
	BasedOn            []Reference          `json:"basedOn,omitempty"`
	PartOf             []Reference          `json:"partOf,omitempty"`
	Status             codes.EventStatus    `json:"status"`
	NotDone            bool                 `json:"notDone,omitempty"`
	NotDoneReason      *CodeableConcept     `json:"notDoneReason,omitempty"`
	Category           *CodeableConcept     `json:"category,omitempty"`
	Subject            Reference            `json:"subject"`
	Context            *Reference           `json:"context,omitempty"`
	PerformedDateTime  *time.Time           `json:"performedDateTime,omitempty"`
	PerformedPeriod    *Period              `json:"performedPeriod,omitempty"`
	Performer          []ProcedurePerformer `json:"performer,omitempty"`
	Location           *Reference           `json:"location,omitempty"`
	ReasonCode         []CodeableConcept    `json:"reasonCode,omitempty"`
	ReasonReference    []Reference          `json:"reasonReference,omitempty"`
	BodySite           []CodeableConcept    `json:"bodySite,omitempty"`
	Outcome            *CodeableConcept     `json:"outcome,omitempty"`
	Report             []Reference          `json:"report,omitempty"`
	Complication       []CodeableConcept    `json:"complication,omitempty"`
	ComplicationDetail []Reference          `json:"complicationDetail,omitempty"`
	FollowUp           []CodeableConcept    `json:"followUp,omitempty"`
	Note               []Annotation         `json:"note,omitempty"`
}

type ProcedurePerformer struct {
	Role       *CodeableConcept `json:"role,omitempty"`
	Actor      Reference        `json:"actor"`
	OnBehalfOf *Reference       `json:"onBehalfOf,omitempty"`
}

type FamilyMemberHistory struct {
	// http://hl7.org/implement/standards/fhir/STU3/familymemberhistory.html
	BaseResource
	Definition       []Reference                    `json:"definition,omitempty"`
	Status           codes.FamilyHistoryStatusCode  `json:"status"`
	NotDone          bool                           `json:"notDone,omitempty"`
	NotDoneReason    *CodeableConcept               `json:"notDoneReason,omitempty"`
	Patient          Reference                      `json:"patient"`
	Date             *time.Time                     `json:"date,omitempty"`
	Name             string                         `json:"name,omitempty"`
	Relationship     CodeableConcept                `json:"relationship,omitempty"`
	Gender           *codes.AdministrativeGender    `json:"gender,omitempty"`
	BornPeriod       *Period                        `json:"bornPeriod,omitempty"`
	BornDate         *time.Time                     `json:"bornDate,omitempty"`
	BornString       string                         `json:"bornString,omitempty"`
	AgeAge           int                            `json:"ageAge,omitempty"`
	AgeString        string                         `json:"ageString,omitempty"`
	EstimatedAge     bool                           `json:"estimatedAge,omitempty"`
	DeceasedBoolean  bool                           `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime *time.Time                     `json:"deceasedDateTime,omitempty"`
	ReasonCode       []CodeableConcept              `json:"reasonCode,omitempty"`
	ReasonReference  []Reference                    `json:"reasonReference,omitempty"`
	Note             []Annotation                   `json:"note,omitempty"`
	Condition        []FamilyMemberHistoryCondition `json:"condition"`
}

type FamilyMemberHistoryCondition struct {
	Code        CodeableConcept  `json:"code"`
	Outcome     *CodeableConcept `json:"outcome,omitempty"`
	OnsetPeriod *Period          `json:"onsetPeriod,omitempty"`
	OnsetString string           `json:"onsetString,omitempty"`
	Note        []Annotation     `json:"note"`
}

type ClinicalImpression struct {
	// http://hl7.org/implement/standards/fhir/STU3/clinicalimpression.html
	BaseResource
	Status                   codes.ClinicalImpressionStatus    `json:"status"`
	Code                     *CodeableConcept                  `json:"code,omitempty"`
	Description              string                            `json:"description,omitempty"`
	Subject                  Reference                         `json:"subject"`
	Context                  *Reference                        `json:"context,omitempty"`
	EffectiveDateTime        *time.Time                        `json:"effectiveDateTime,omitempty"` // Time of assessment
	Date                     *time.Time                        `json:"date,omitempty"`              // When the assessment was documented
	Assessor                 *Reference                        `json:"assessor,omitempty"`
	Previous                 *Reference                        `json:"previous,omitempty"`
	Problem                  []Reference                       `json:"problem,omitempty"`
	Investigation            []ClinicalImpressionInvestigation `json:"investigation,omitempty"`
	Protocol                 []string                          `json:"protocol,omitempty"`
	Summary                  string                            `json:"summary,omitempty"`
	Finding                  []ClinicalImpressionFinding       `json:"finding,omitempty"`
	PrognosisCodeableConcept []CodeableConcept                 `json:"prognosisCodeableConcept,omitempty"`
	PrognosisReference       []Reference                       `json:"prognosisReference,omitempty"`
	Action                   []Reference                       `json:"action,omitempty"`
	Note                     []Annotation                      `json:"note,omitempty"`
}

type ClinicalImpressionInvestigation struct {
	Code CodeableConcept `json:"code"`
	Item []Reference     `json:"item,omitempty"`
}

type ClinicalImpressionFinding struct {
	ItemCodeableConcept *CodeableConcept `json:"itemCodeableConcept,omitempty"`
	ItemReference       *Reference       `json:"itemReference,omitempty"`
	Basis               string           `json:"basis"`
}
