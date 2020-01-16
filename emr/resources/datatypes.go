package resources

import (
	"github.com/hellodoctordev/common/emr/codes"
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
	System codes.ContactPointSystem `json:"system"`
	Value  string                   `json:"value,omitempty"`
	Use    codes.ContactPointUse    `json:"use"`
	Rank   int                      `json:"rank,omitempty"`
}

type Address struct {
	Use        codes.AddressUse  `json:"use"`
	Type       codes.AddressType `json:"type"`
	Text       string            `json:"text,omitempty"`
	Line       []string          `json:"line"`
	City       string            `json:"city"`
	District   string            `json:"district,omitempty"`
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

type PatientLink struct {
	Other Reference      `json:"other"` // Reference Types: Patient, RelatedPerson
	Type  codes.LinkType `json:"type"`
}

type Qualification struct {
	Identifier Identifier      `json:"identifier"`
	Code       CodeableConcept `json:"code"`
	Issuer     *Reference      `json:"issuer"` // Organization that regulates and issues the qualification
}

type PersonLink struct {
	Target    Reference                    `json:"target"`
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

type ConditionStage struct {
	Summary    CodeableConcept `json:"summary,omitempty"`
	Assessment []Reference     `json:"assessment,omitempty"`
}

type ConditionEvidence struct {
	Code   []CodeableConcept `json:"code,omitempty"`
	Detail []Reference       `json:"detail,omitempty"`
}

type Period struct {
	Start *time.Time `json:"start"`
	End   *time.Time `json:"end"`
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

type Annotation struct {
	AuthorReference *Reference `json:"authorReference,omitempty"`
	AuthorString    string     `json:"authorString,omitempty"`
	Time            *time.Time `json:"time,omitempty"`
	Text            string     `json:"text"`
}

type CarePlanActivityDetail struct {
	Category               *CodeableConcept             `json:"category,omitempty"`
	Definition             *Reference                   `json:"definition,omitempty"`
	Code                   *CodeableConcept             `json:"code,omitempty"`
	ReasonCode             []CodeableConcept            `json:"reasonCode,omitempty"`
	ReasonReference        []Reference                  `json:"reasonReference,omitempty"`
	Goal                   []Reference                  `json:"goal,omitempty"`
	Status                 codes.CarePlanActivityStatus `json:"status"`
	StatusReason           string                       `json:"statusReason,omitempty"`
	Prohibited             bool                         `json:"prohibited,omitempty"`
	ScheduledPeriod        *Period                      `json:"scheduledPeriod,omitempty"`
	ScheduledString        string                       `json:"scheduledString,omitempty"`
	Location               *Reference                   `json:"location,omitempty"`
	Performer              []Reference                  `json:"performer,omitempty"`
	ProductCodeableConcept *CodeableConcept             `json:"productCodeableConcept,omitempty"`
	ProductReference       *Reference                   `json:"productReference,omitempty"`
	DailyAmount            *Quantity                    `json:"dailyAmount,omitempty"`
	Quantity               *Quantity                    `json:"quantity,omitempty"`
	Description            string                       `json:"description,omitempty"`
}

type CarePlanActivity struct {
	OutcomeCodeableConcept []CodeableConcept       `json:"outcomeCodeableConcept,omitempty"`
	OutcomeReference       []Reference             `json:"outcomeReference,omitempty"`
	Progress               []Annotation            `json:"progress,omitempty"`
	Reference              *Reference              `json:"reference,omitempty"`
	Detail                 *CarePlanActivityDetail `json:"detail,omitempty"`
	Note                   []Annotation            `json:"note,omitempty"`
}
