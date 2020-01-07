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

func (a AllergyIntolerance) IsFamilyHistoryReasonReference() {}
func (a AllergyIntolerance) IsClinicalImpressionProblem()    {}

type ConditionSubject interface {
	IsConditionSubject()
}

type ConditionContext interface {
	IsConditionContext()
}

type ConditionAsserter interface {
	IsConditionAsserter()
}

type Condition struct {
	Identifier         []Identifier                      `json:"identifier"`
	ClinicalStatus     codes.ConditionClinicalStatus     `json:"clinicalStatus"`
	VerificationStatus codes.ConditionVerificationStatus `json:"verificationStatus"`
	Category           []codes.ConditionCategory         `json:"category"`
	Severity           codes.ConditionSeverity           `json:"severity"`
	Code               codes.ConditionCode               `json:"code"`
	BodySite           []codes.BodyStructure             `json:"bodySite"`
	Subject            ConditionSubject                  `json:"subject"`
	Context            ConditionContext                  `json:"context"`
	Onset              Onset                             `json:"onset"`
	Abatement          Abatement                         `json:"abatement"`
	AssertedDate       *time.Time                        `json:"assertedDate"`
	Asserter           ConditionAsserter                 `json:"asserter"`
	Stage              ConditionStage                    `json:"stage"`
	Evidence           []ConditionEvidence               `json:"evidence"`
	Note               []string                          `json:"note"`
}

func (c Condition) IsProcedureReasonReference()     {}
func (c Condition) IsFamilyHistoryReasonReference() {}
func (c Condition) IsClinicalImpressionProblem()    {}
func (c Condition) IsInvestigationItem()            {}
func (c Condition) IsFindingItemReference()         {}
func (c Condition) IsGoalAddresses()                {}

type Procedure struct {
	Identifier         []Identifier                      `json:"identifier"`
	Definition         []ProcedureDefinition             `json:"definition"`
	Status             codes.EventStatus                 `json:"status"`
	NotDone            bool                              `json:"notDone"`
	NotDoneReason      codes.ProcedureNotPerformedReason `json:"notDoneReason"`
	Category           codes.ProcedureCategory           `json:"category"`
	Subject            ProcedureSubject                  `json:"subject"`
	Context            ProcedureContext                  `json:"context"`
	Performed          ProcedurePerformed                `json:"performed"`
	Performer          []ProcedurePerformer              `json:"performer"`
	Location           *Location                         `json:"location"`
	ReasonCode         []codes.ProcedureReasonCode       `json:"reasonCode"`
	ReasonReference    []ProcedureReasonReference        `json:"reasonReference"`
	BodySite           codes.BodyStructure               `json:"bodySite"`
	Outcome            codes.ProcedureOutcomeCode        `json:"outcome"`
	Report             []DiagnosticReport                `json:"report"`
	Complication       []codes.ConditionCode             `json:"complication"`
	ComplicationDetail []Condition                       `json:"complicationDetail"`
	FollowUp           []codes.ProcedureFollowUpCode     `json:"followUp"`
	Note               []string                          `json:"note"`
}

func (p Procedure) IsClinicalImpressionAction() {}

type FamilyMemberHistory struct {
	// http://hl7.org/implement/standards/fhir/STU3/familymemberhistory.html
	Identifier       []Identifier                     `json:"identifier"`
	Definition       FamilyMemberHistoryDefinition    `json:"definition"`
	Status           codes.FamilyHistoryStatusCode    `json:"status"`
	NotDone          bool                             `json:"notDone"`
	NotDoneReason    codes.FamilyHistoryNotDoneReason `json:"notDoneReason"`
	Patient          *Patient                         `json:"patient"`
	Date             *time.Time                       `json:"date"`
	Name             string                           `json:"name"`
	Relationship     codes.FamilyMemberCode           `json:"relationship"`
	Gender           codes.AdministrativeGender       `json:"gender"`
	Born             ApproximateBirthDate             `json:"born"`
	Age              ApproximateAge                   `json:"age"`
	EstimatedAge     bool                             `json:"estimatedAge"`
	DeceasedBoolean  bool                             `json:"deceasedBoolean"`
	DeceasedDateTime *time.Time                       `json:"deceasedDateTime"`
	ReasonCode       []codes.ClinicalFindingCode      `json:"reasonCode"`
	ReasonReference  []FamilyHistoryReasonReference   `json:"reasonReference"`
	Note             []string                         `json:"note"`
	Condition        []FamilyMemberHistoryCondition   `json:"condition"`
}

func (f FamilyMemberHistory) IsValid() bool {
	return f.DeceasedBoolean == false || f.DeceasedDateTime == nil
}

func (f FamilyMemberHistory) IsInvestigationItem() {}

type ClinicalImpression struct {
	// http://hl7.org/implement/standards/fhir/STU3/clinicalimpression.html
	Identifier               []Identifier                        `json:"identifier"`
	Status                   codes.ClinicalImpressionStatus      `json:"status"`
	Code                     string                              `json:"code"`
	Description              string                              `json:"description"`
	Subject                  *Patient                            `json:"subject"`
	Context                  ClinicalImpressionContext           `json:"context"`
	EffectiveDateTime        *time.Time                          `json:"effectiveDateTime"` // Time of assessment
	Date                     *time.Time                          `json:"date"`              // When the assessment was documented
	Assessor                 *Practitioner                       `json:"assessor"`
	Previous                 *ClinicalImpression                 `json:"previous"`
	Problem                  []ClinicalImpressionProblem         `json:"problem"`
	Investigation            []ClinicalImpressionInvestigation   `json:"investigation"`
	Protocol                 []string                            `json:"protocol"`
	Summary                  string                              `json:"summary"`
	Finding                  []InvestigationFinding              `json:"finding"`
	PrognosisCodeableConcept []codes.ClinicalImpressionPrognosis `json:"prognosisCodeableConcept"`
	Action                   []ClinicalImpressionAction          `json:"action"`
	Note                     []string                            `json:"note"`
}

func (ci ClinicalImpression) IsConditionStageAssessment() {}
