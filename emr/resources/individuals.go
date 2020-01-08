package resources

import (
	"github.com/hellodoctordev/common/emr/codes"
	"time"
)

type Person struct {
	// http://hl7.org/implement/standards/fhir/STU3/person.html
	BaseResource
	Identifier []Identifier               `json:"identifier"`
	Active     bool                       `json:"active,omitempty"`
	Name       []HumanName                `json:"name,omitempty"`
	Telecom    []ContactPoint             `json:"telecom,omitempty"`
	Gender     codes.AdministrativeGender `json:"gender,omitempty"`
	BirthDate  *Date                      `json:"birthDate,omitempty"`
	Address    []Address                  `json:"address,omitempty"`
	Photo      []Attachment               `json:"photo,omitempty"`
	Link       []PersonLink               `json:"link,omitempty"`
}

func (p Person) GetResourceType() string { return "Person" }
func (p Person) IsPersonLinkTarget()     {}

type Patient struct {
	// http://hl7.org/implement/standards/fhir/STU3/patient.html
	Person
	DeceasedBoolean      bool             `json:"deceasedBoolean,omitempty"`
	DeceasedDateTime     *time.Time       `json:"deceasedDateTime,omitempty"`
	MaritalStatus        *CodeableConcept `json:"maritalStatus"`
	Contact              []Contact        `json:"contact"`
	Communication        []Communication  `json:"communication"`
	GeneralPractitioner  []Reference      `json:"generalPractitioner"`  // Reference Types: Organization, Practitioner
	ManagingOrganization *Reference       `json:"managingOrganization"` // Reference Types: Organization
	Link                 []PatientLink    `json:"link"`
}

func (p Patient) GetResourceType() string { return "Patient" }

func (p Patient) IsValid() bool {
	return p.DeceasedBoolean == false || p.DeceasedDateTime == nil
}
func (p Patient) IsPatientLinkOther()          {}
func (p Patient) IsConditionSubject()          {}
func (p Patient) IsConditionAsserter()         {}
func (p Patient) IsProcedureSubject()          {}
func (p Patient) IsProcedurePerformer()        {}
func (p Patient) IsObservationSubject()        {}
func (p Patient) IsObservationPerformer()      {}
func (p Patient) IsCarePlanAuthor()            {}
func (p Patient) IsGoalSubject()               {}
func (p Patient) IsGoalExpressedBy()           {}
func (p Patient) IsAuthorReference()           {}
func (p Patient) IsCarePlanActivityPerformer() {}

type RelatedPerson struct {
	// http://hl7.org/implement/standards/fhir/STU3/relatedperson.html
	Person
	Relationship CodeableConcept `json:"relationship"`
}

func (rp RelatedPerson) IsPatientLinkOther()          {}
func (rp RelatedPerson) IsConditionAsserter()         {}
func (rp RelatedPerson) IsProcedurePerformer()        {}
func (rp RelatedPerson) IsObservationPerformer()      {}
func (rp RelatedPerson) IsCarePlanAuthor()            {}
func (rp RelatedPerson) IsGoalExpressedBy()           {}
func (rp RelatedPerson) IsAuthorReference()           {}
func (rp RelatedPerson) IsCarePlanActivityPerformer() {}

type Practitioner struct {
	// http://hl7.org/implement/standards/fhir/STU3/practitioner.html
	Person
	Qualification []Qualification   `json:"qualification"`
	Communication []CodeableConcept `json:"communication"`
}

func (p Practitioner) GetResourceType() string      { return "Practitioner" }
func (p Practitioner) IsPersonLinkTarget()          {}
func (p Practitioner) IsConditionAsserter()         {}
func (p Practitioner) IsProcedurePerformer()        {}
func (p Practitioner) IsObservationPerformer()      {}
func (p Practitioner) IsCarePlanAuthor()            {}
func (p Practitioner) IsGoalExpressedBy()           {}
func (p Practitioner) IsAuthorReference()           {}
func (p Practitioner) IsCarePlanActivityPerformer() {}

type PractitionerRole struct {
	// http://hl7.org/implement/standards/fhir/STU3/practitionerrole.html
	Identifier        []Identifier                 `json:"identifier"`
	Active            bool                         `json:"active"`
	Organization      Organization                 `json:"organization"` // Organization where the roles are available
	Code              []codes.PractitionerRoleCode `json:"code"`
	Speciality        []codes.Specialty            `json:"specialty"`
	Location          []Location                   `json:"location"`
	HealthcareService []HealthcareService          `json:"healthcareService"`
	Telecom           []ContactPoint               `json:"telecom"`
}
