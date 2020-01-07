package resources

import (
	"hellodoctor/emr/x/codes"
	"time"
)

type Reference struct {
	Reference  string      `json:"reference"`
	Identifier *Identifier `json:"identifier"`
	Display    string      `json:"display"`
}

type Date string

type Identifier struct {
	Use      codes.IdentifierUse `json:"use"`
	Type     CodeableConcept     `json:"type"`
	System   string              `json:"system"` // The namespace for the identifier value
	Value    string              `json:"value"`
	Assigner *Reference          `json:"assigner"` // Reference Types: Organization
}

type HumanName struct {
	Use    codes.NameUse `json:"use"`
	Text   string        `json:"text,omitempty"`
	Family string        `json:"family,omitempty"`
	Given  []string      `json:"given,omitempty"`
	Prefix string        `json:"prefix,omitempty"`
	Suffix string        `json:"suffix,omitempty"`
}

type ContactPoint struct {
	System *codes.ContactPointSystem `json:"system"`
	Value  string                    `json:"value,omitempty"`
	Use    *codes.ContactPointUse    `json:"use"`
	Rank   int                       `json:"rank,omitempty"`
}

type Address struct {
	Use        codes.AddressUse  `json:"use"`
	Type       codes.AddressType `json:"type"`
	Text       string            `json:"text"`
	Line       []string          `json:"line"`
	City       string            `json:"city"`
	District   string            `json:"district"`
	State      string            `json:"state"`
	PostalCode string            `json:"postalCode"`
	Country    string            `json:"country"`
}

type Attachment struct {
	ID          string         `json:"id"`
	ContentType codes.MimeType `json:"contentType"`
	Language    codes.Language `json:"language"`
	URL         string         `json:"url"`
	Size        uint64         `json:"size"`
	Title       string         `json:"title"`
	Creation    time.Time      `json:"creation"`
}

type Contact struct {
	Relationship []codes.ContactRole         `json:"relationship"`
	Name         *HumanName                  `json:"name"`
	Telecom      []ContactPoint              `json:"telecom"`
	Address      *Address                    `json:"address"`
	Gender       *codes.AdministrativeGender `json:"gender"`
}

type Communication struct {
	Language  CodeableConcept `json:"language"`
	Preferred bool            `json:"preferred"`
}

type PatientLinkOther interface {
	IsPatientLinkOther()
}

type PatientLink struct {
	Other Reference      `json:"other"` // Reference Types: Patient, RelatedPerson
	Type  codes.LinkType `json:"type"`
}

type Qualification struct {
	Identifier Identifier              `json:"identifier"`
	Code       codes.QualificationCode `json:"code"`
	Issuer     *Organization           `json:"issuer"` // Organization that regulates and issues the qualification
}

type PersonLinkTarget interface {
	IsPersonLinkTarget()
}

type PersonLink struct {
	Target    PersonLinkTarget             `json:"target"`
	Assurance codes.IdentityAssuranceLevel `json:"assurance"`
}

type AllergyIntoleranceReaction struct {
	Substance     codes.SubstanceCode              `json:"substance"`
	Manifestation []codes.ClinicalFindingCode      `json:"manifestation"`
	Description   string                           `json:"description"`
	Onset         *time.Time                       `json:"onset"`
	Severity      codes.AllergyIntoleranceSeverity `json:"severity"`
	ExposureRoute codes.RouteCode                  `json:"exposureRoute"`
	Note          []string                         `json:"note"`
}

type Onset struct {
	OnsetDateTime *time.Time `json:"onsetDateTime"`
	OnsetString   string     `json:"onsetString"`
}

type Abatement struct {
	AbatementDateTime *time.Time `json:"abatementDateTime"`
}

type ConditionStageAssessment interface {
	IsConditionStageAssessment()
}

type ConditionStage struct {
	Summary    codes.ConditionStageCode   `json:"summary"`
	Assessment []ConditionStageAssessment `json:"assessment"`
}

type ConditionEvidence struct {
	Code   []codes.ManifestationAndSymptomCode `json:"code"`
	Detail []struct{}
}

type ProcedureDefinition interface {
	IsProcedureDefinition()
}

type ProcedureSubject interface {
	IsProcedureSubject()
}

type ProcedureContext interface {
	IsProcedureContext()
}

type Period struct {
	Start *time.Time `json:"start"`
	End   *time.Time `json:"end"`
}

type ProcedurePerformed struct {
	PerformedDateTime *time.Time `json:"performedDateTime"`
	PerformedPeriod   *Period    `json:"performedPeriod"`
}

type ProcedurePerformerActor interface {
	IsProcedurePerformer()
}

type ProcedurePerformer struct {
	Role  codes.ProcedurePerformerRoleCode `json:"role"`
	Actor ProcedurePerformerActor          `json:"actor"`
}

type ProcedureReasonReference interface {
	IsProcedureReasonReference()
}

type FamilyMemberHistoryDefinition interface {
	IsFamilyMemberHistoryDefinition()
}

type ApproximateBirthDate struct {
	BornPeriod *Period    `json:"bornPeriod"`
	BornDate   *time.Time `json:"bornDate"`
	BornString string     `json:"bornString"`
}

type ApproximateAge struct {
	AgeAge    int    `json:"ageAge"`
	AgeString string `json:"ageString"`
}

type FamilyHistoryReasonReference interface {
	IsFamilyHistoryReasonReference()
}

type FamilyMemberHistoryCondition struct {
	Code    codes.ConditionCode        `json:"code"`
	Outcome codes.ConditionOutcomeCode `json:"outcome"`
	Onset   Onset                      `json:"onset"`
	Note    []string                   `json:"note"`
}

type ClinicalImpressionContext interface {
	IsClinicalImpressionContext()
}

type ClinicalImpressionProblem interface {
	IsClinicalImpressionProblem()
}

type InvestigationItem interface {
	IsInvestigationItem()
}

type ClinicalImpressionInvestigation struct {
	Code codes.InvestigationTypeCode `json:"code"`
	Item []InvestigationItem         `json:"item"`
}

type Coding struct {
	System       string `json:"system"`
	Version      string `json:"version"`
	Code         string `json:"code"`
	Display      string `json:"display"`
	UserSelected bool   `json:"userSelected"`
}

type CodeableConcept struct {
	Coding []Coding `json:"coding"`
	Text   string   `json:"text"`
}

type FindingItemReference interface {
	IsFindingItemReference()
}

type InvestigationFinding struct {
	ItemCodeableConcept *CodeableConcept      `json:"itemCodeableConcept,omitempty"`
	ItemReference       *FindingItemReference `json:"itemReference,omitempty"`
	Basis               string                `json:"basis"`
}

func (i InvestigationFinding) IsValid() bool {
	return i.ItemCodeableConcept == nil || i.ItemReference == nil
}

type ClinicalImpressionAction interface {
	IsClinicalImpressionAction()
}

type ObservationBasedOnReference interface {
	IsObservationBasedOnReference()
}

type ObservationSubject interface {
	IsObservationSubject()
}

type ObservationContext interface {
	IsObservationContext()
}

type ObservationPerformer interface {
	IsObservationPerformer()
}

type QuantityComparatorCode = string

const (
	QuantityComparatorCodeLessThan           QuantityComparatorCode = "<"
	QuantityComparatorCodeLessThanOrEqual    QuantityComparatorCode = "<="
	QuantityComparatorCodeGreaterThan        QuantityComparatorCode = ">"
	QuantityComparatorCodeGreaterThanOrEqual QuantityComparatorCode = ">="
)

type Quantity struct {
	Value      float64                `json:"value"`
	Comparator QuantityComparatorCode `json:"comparator"`
	Unit       string                 `json:"unit"`
	System     string                 `json:"system"`
	Code       string                 `json:"code"`
}

type Range struct {
	Low  *Quantity `json:"low"`
	High *Quantity `json:"high"`
}

type CarePlanDefinition interface {
	IsCarePlanDefinition()
}

type CarePlanContext interface {
	IsCarePlanContext()
}

type CarePlanAuthor interface {
	IsCarePlanAuthor()
}

type GoalSubject interface {
	IsGoalSubject()
}

type GoalTarget struct {
	Measure        *codes.ObservationCode `json:"measure"`
	DetailQuantity *Quantity              `json:"detailQuantity,omitempty"`
	DetailRange    *Range                 `json:"range"`
	DueDate        *time.Time             `json:"dueDate"`
}

func (gt GoalTarget) IsValid() bool {
	return gt.DetailQuantity == nil || gt.DetailRange == nil
}

type GoalExpressedBy interface {
	IsGoalExpressedBy()
}

type GoalAddresses interface {
	IsGoalAddresses()
}

type Goal struct {
	Identifier       []Identifier                `json:"identifier"`
	Status           codes.GoalStatus            `json:"status"`
	Category         codes.GoalCategory          `json:"category"`
	Priority         codes.GoalPriority          `json:"priority"`
	Description      string                      `json:"description"`
	Subject          GoalSubject                 `json:"subject"`
	StartDate        *time.Time                  `json:"startDate"`
	Target           *GoalTarget                 `json:"target"`
	StatusDate       *time.Time                  `json:"statusDate"`
	StatusReason     string                      `json:"statusReason"`
	ExpressedBy      *GoalExpressedBy            `json:"expressedBy"`
	Addresses        []GoalAddresses             `json:"addresses"`
	Note             []string                    `json:"note"`
	OutcomeCode      []codes.ClinicalFindingCode `json:"outcomeCode"`
	OutcomeReference []Observation               `json:"outcomeReference"`
}

type AuthorReference interface {
	IsAuthorReference()
}

type Annotation struct {
	AuthorReference *AuthorReference `json:"authorReference"`
	AuthorString    string           `json:"authorString"`
	Time            *time.Time       `json:"time"`
	Text            string           `json:"text"`
}

type CarePlanActivityDetailDefinition interface {
	IsCarePlanActivityDetailDefinition()
}

type CarePlanActivityPerformer interface {
	IsCarePlanActivityPerformer()
}

type ProductReference interface {
	IsProductReference()
}

type CarePlanActivityDetail struct {
	Category               *codes.CarePlanActivityDetailCategory `json:"category"`
	Definition             *CarePlanActivityDetailDefinition     `json:"definition"`
	Code                   *codes.CarePlanActivityCode           `json:"code"`
	ReasonCode             *codes.ActivityReasonCode             `json:"reasonCode"`
	ReasonReference        []Condition                           `json:"reasonReference"`
	Goal                   []Goal                                `json:"goal"`
	Status                 codes.CarePlanActivityStatus          `json:"status"`
	StatusReason           string                                `json:"statusReason"`
	Prohibited             bool                                  `json:"prohibited"`
	ScheduledPeriod        *Period                               `json:"scheduledPeriod"`
	ScheduledString        string                                `json:"scheduledString"`
	Location               *Location                             `json:"location"`
	Performer              []CarePlanActivityPerformer           `json:"performer"`
	ProductCodeableConcept *CodeableConcept                      `json:"productCodeableConcept"`
	ProductReference       *ProductReference                     `json:"productReference"`
	DailyAmount            *Quantity                             `json:"dailyAmount"`
	Quantity               *Quantity                             `json:"quantity"`
	Description            string                                `json:"description"`
}

func (c CarePlanActivityDetail) IsValid() bool {
	return (c.ScheduledPeriod == nil || c.ScheduledString == "") && (c.ProductCodeableConcept == nil || c.ProductReference == nil)
}

type CarePlanActivity struct {
	Outcome          []codes.CarePlanActivityOutcome `json:"outcomeCodeableConcept"`
	OutcomeReference []struct{}                      `json:"outcomeReference"`
	Progress         []Annotation                    `json:"progress"`
	Reference        *struct{}                       `json:"reference"`
	Detail           *CarePlanActivityDetail         `json:"detail"`
	Note             []Annotation                    `json:"note"`
}
