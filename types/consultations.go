package types

import (
	"cloud.google.com/go/firestore"
	"github.com/hellodoctordev/common/fhir/resources"
	"time"
)

type ConsultationType = string

const (
	ConsultationTypeVideo       ConsultationType = "video"
	ConsultationTypePhone       ConsultationType = "phone"
	ConsultationTypeChat        ConsultationType = "chat"
	ConsultationTypeGracePeriod ConsultationType = "gracePeriod"
	ConsultationTypeOffice      ConsultationType = "office"
)

type BillingType = string

const (
	BillingTypeFree         BillingType = "free"
	BillingTypeHourly       BillingType = "hourly"
	BillingTypeConsultation BillingType = "consultation"
)

type BillingDetails struct {
	Type             BillingType            `firestore:"type,omitempty" json:"type,omitempty"`
	Amount           int64                  `firestore:"amount" json:"amount"`
	Currency         string                 `firestore:"currency,omitempty" json:"currency,omitempty"`
	IsSubscribed     bool                   `firestore:"isSubscribed,omitempty" json:"isSubscribed,omitempty"`
	AuthorizedCharge *firestore.DocumentRef `firestore:"authorizedCharge" json:"authorizedCharge"`
}

type ConsultationStatus = string

const (
	ConsultationStatusPending   ConsultationStatus = "pending"
	ConsultationStatusAccepted  ConsultationStatus = "accepted"
	ConsultationStatusRejected  ConsultationStatus = "rejected"
	ConsultationStatusCancelled ConsultationStatus = "cancelled"
)

type ConsultationVisibility = string

const (
	ConsultationVisibilityVisible ConsultationVisibility = "visible"
	ConsultationVisibilityHidden  ConsultationVisibility = "hidden"
)

type ConsultationInterface interface {
	SetConsultationRef(ref *firestore.DocumentRef)
}

type ConsultationExaminationTemplateOverride struct {
	Action string            `firestore:"action" json:"action"`
	ItemID string            `firestore:"itemID" json:"itemID"`
	Data   map[string]string `firestore:"data" json:"data"`
}

type ConsultationExaminationTemplate struct {
	SpecialtyCode    string                                    `firestore:"specialtyCode" json:"specialtyCode"`
	SpecialtyDisplay string                                    `firestore:"specialtyDisplay" json:"specialtyDisplay"`
	Overrides        []ConsultationExaminationTemplateOverride `firestore:"overrides" json:"overrides"`
}

type RequestMode = string

const (
	RequestModeInstant      RequestMode = "instant"
	RequestModeStandard     RequestMode = "standard"
	RequestModePractitioner RequestMode = "practitioner"
)

type AcceptMode = string

const (
	AcceptModeAutomatic RequestMode = "automatic"
	AcceptModeManual    RequestMode = "manual"
)

type Consultation struct {
	Ref              *firestore.DocumentRef            `firestore:"-" json:"ref"`
	Resource         *FHIRResource                     `firestore:"resource" json:"resource"`
	Event            *firestore.DocumentRef            `firestore:"event,omitempty" json:"eventRef"`
	Patient          *firestore.DocumentRef            `firestore:"patient" json:"patientRef"`
	Practitioner     *firestore.DocumentRef            `firestore:"practitioner" json:"practitionerRef"`
	Chat             *firestore.DocumentRef            `firestore:"chat" json:"chatRef"`
	Parent           *firestore.DocumentRef            `firestore:"parent,omitempty" json:"parent"`
	Type             ConsultationType                  `firestore:"type" json:"type"`
	Location         map[string]interface{}            `firestore:"location" json:"location"`
	Templates        []ConsultationExaminationTemplate `firestore:"templates,omitempty" json:"templates"`
	Status           ConsultationStatus                `firestore:"status" json:"status"`
	RequestMode      RequestMode                       `firestore:"requestMode,omitempty" json:"requestMode"`
	AcceptMode       AcceptMode                        `firestore:"acceptMode,omitempty" json:"acceptMode"`
	ScheduledStart   time.Time                         `firestore:"scheduledStart,omitempty" json:"scheduledStart,omitempty"`
	ScheduledEnd     time.Time                         `firestore:"scheduledEnd,omitempty" json:"scheduledEnd,omitempty"`
	Start            *time.Time                        `firestore:"start,omitempty" json:"start,omitempty"`
	End              *time.Time                        `firestore:"end,omitempty" json:"end,omitempty"`
	IsFinalized      bool                              `firestore:"isFinalized,omitempty" json:"isFinalized"`
	Visibility       ConsultationVisibility            `firestore:"visibility,omitempty" json:"visibility"`
	Reason           string                            `firestore:"reason,omitempty" json:"reason"`
	Findings         []ConsultationFinding             `firestore:"findings,omitempty" json:"findings"`
	Observations     []ConsultationObservation         `firestore:"observations,omitempty" json:"observations"`
	Prescriptions    []*firestore.DocumentRef          `firestore:"prescriptions,omitempty" json:"prescriptions"`
	Invoices         []*firestore.DocumentRef          `firestore:"invoices" json:"invoices"`
	Notes            []Note                            `firestore:"notes,omitempty" json:"notes"`
	Billing          *BillingDetails                   `firestore:"billing" json:"billing"`
	CreatedTime      time.Time                         `firestore:"createdTime" json:"createdTime"`
	AcceptedTime     *time.Time                        `firestore:"acceptedTime,omitempty" json:"acceptedTime"`
	ConfirmedTime    *time.Time                        `firestore:"confirmedTime,omitempty" json:"confirmedTime"`
	FollowupMessages []map[string]interface{}          `firestore:"followupMessages,omitempty" json:"followupMessages"`
	GroupID          string                            `firestore:"groupID" json:"groupID"`
	Specialty        string                            `firestore:"specialty" json:"specialty"`
	Requirements     ConsultationRequirements          `firestore:"requirements" json:"requirements"`
}

type ConsultationRequirements struct {
	Biometrics           string `firestore:"biometrics" json:"biometrics"`
	MedicalProfile       string `firestore:"medicalProfile" json:"medicalProfile"`
	Symptoms             string `firestore:"symptoms" json:"symptoms"`
	Payment              string `firestore:"payment" json:"payment"`
	InsuranceInformation string `firestore:"insuranceInformation" json:"insuranceInformation"`
}

func (c Consultation) GetPreview() ConsultationPreview {
	preview := ConsultationPreview{
		Consultation:    c.Ref,
		Type:            c.Type,
		Status:          c.Status,
		NeedsFinalizing: !c.IsFinalized,
		TotalFindings:   len(c.Findings),
		Prescriptions:   c.Prescriptions,
	}

	if c.Start != nil {
		preview.Start = *c.Start
	} else {
		preview.Start = c.ScheduledStart
	}

	for _, observation := range c.Observations {
		switch observation.Code.Code {
		case LOINCBodyWeightCode:
			preview.BodyWeightKG = observation.Value.Value
		case LOINCBodyTemperatureCode:
			preview.TemperatureCelsius = observation.Value.Value
		case LOINCHeartRateCode:
			preview.HeartRate = observation.Value.Value
		case LOINCBloodPressureCode:
			for _, component := range observation.Components {
				switch component.Code.Code {
				case LOINCSystolicCode:
					preview.Systolic = component.Value.Value
				case LOINCDiastolicCode:
					preview.Diastolic = component.Value.Value
				}
			}
		}
	}

	return preview
}

const (
	LOINCBodyWeightCode      = "29463-7"
	LOINCBodyTemperatureCode = "8310-5"
	LOINCHeartRateCode       = "8867-4"
	LOINCBloodPressureCode   = "85354-9"
	LOINCSystolicCode        = "8480-6"
	LOINCDiastolicCode       = "8462-4"
)

type PatientPreview struct {
	ID          string `firestore:"id" json:"id"`
	DisplayName string `firestore:"displayName" json:"displayName"`
}

type ConsultationPreview struct {
	Consultation       *firestore.DocumentRef   `firestore:"consultation" json:"consultationRef"`
	Start              time.Time                `firestore:"start" json:"start"`
	Status             string                   `firestore:"status" json:"status"`
	Type               string                   `firestore:"type" json:"type"`
	NeedsFinalizing    bool                     `firestore:"needsFinalizing" json:"needsFinalizing"`
	BodyWeightKG       float64                  `firestore:"bodyWeightKG" json:"bodyWeightKG"`
	TemperatureCelsius float64                  `firestore:"temperatureCelsius,omitempty" json:"temperatureCelsius"`
	HeartRate          float64                  `firestore:"heartRate" json:"heartRate"`
	Systolic           float64                  `firestore:"systolic" json:"systolic"`
	Diastolic          float64                  `firestore:"diastolic" json:"diastolic"`
	TotalFindings      int                      `firestore:"totalFindings" json:"totalFindings"`
	Prescriptions      []*firestore.DocumentRef `firestore:"prescriptions" json:"prescriptionsRefs"`
}

type CreateConsultationPayload struct {
	PractitionerUserUID  string `json:"practitionerUserUID"`
	PatientUserUID       string `json:"patientUserUID"`
	ParentConsultationID string `json:"parentConsultationID"`
	GroupID              string `json:"groupID"`
	Specialty            string `json:"specialty"`
	LocationID           string `json:"locationID"`
	Consultation
}

func NewConsultationFromSnapshot(consultationSnapshot *firestore.DocumentSnapshot) (consultation Consultation, err error) {
	if err := consultationSnapshot.DataTo(&consultation); err != nil {
		return consultation, err
	}

	consultation.Ref = consultationSnapshot.Ref

	return consultation, nil
}

func (c *Consultation) SetConsultationRef(ref *firestore.DocumentRef) {
	c.Ref = ref
}

type VideoConsultation struct {
	// additional metadata for video consultations
	Consultation
	RoomName string `firestore:"roomName" json:"roomName"`
	RoomSid  string `firestore:"roomSid" json:"roomSid"`
}

func (c *VideoConsultation) SetConsultationRef(ref *firestore.DocumentRef) {
	c.Ref = ref
}

type VoiceConsultation struct {
	// additional metadata for voice consultations
	Consultation
	CallSID string `firestore:"callSID" json:"callSID"`
}

func (c *VoiceConsultation) SetConsultationRef(ref *firestore.DocumentRef) {
	c.Ref = ref
}

type ConsultationRoom struct {
	Consultation  *firestore.DocumentRef `firestore:"consultation" json:"consultation"`
	RoomType      ConsultationType       `firestore:"roomType" json:"roomType"`
	TwilioRoomSID string                 `firestore:"twilioRoomSID" json:"twilioRoomSID"`
	Status        string                 `firestore:"status" json:"status"`
}

type ConsultationReason = resources.Coding

type ConsultationFinding struct {
	Finding        Finding    `firestore:"finding" json:"finding"`
	ClinicalStatus string     `firestore:"clinicalStatus" json:"clinicalStatus"`
	Severity       string     `firestore:"severity" json:"severity"`
	BodySites      []BodySite `firestore:"bodySites" json:"bodySites"`
	Onset          *time.Time `firestore:"onset" json:"abatement"`
	Abatement      *time.Time `firestore:"abatement" json:"abatement"`
	Notes          []Note     `firestore:"notes" json:"notes"`
}

type Finding = resources.Coding
type BodySite = resources.Coding

type Note struct {
	Time *time.Time `firestore:"time" json:"time"`
	Text string     `firestore:"text" json:"text"`
}

type Prescription struct {
	Patient            *firestore.DocumentRef    `firestore:"patient" json:"patient"`
	Practitioner       *firestore.DocumentRef    `firestore:"practitioner" json:"practitioner"`
	Consultation       *firestore.DocumentRef    `firestore:"consultation" json:"consultation"`
	Status             string                    `firestore:"status" json:"status"`
	Medication         Medication                `firestore:"medication" json:"medication"`
	AuthoredOn         *time.Time                `firestore:"authoredOn" json:"authoredOn"`
	Notes              []Note                    `firestore:"notes" json:"notes"`
	DosageInstructions []Dosage                  `firestore:"dosageInstructions" json:"dosageInstructions"`
	DispenseRequest    *DispenseRequest          `firestore:"dispenseRequest" json:"dispenseRequest"`
	Substitution       *PrescriptionSubstitution `firestore:"substitution" json:"substitution"`
}

type Medication = resources.Coding
type PrescriptionReason = resources.Coding
type Dosage = resources.Dosage
type DispenseRequest = resources.DispenseRequest

type PrescriptionSubstitution struct {
	Allowed bool              `firestore:"allowed" json:"allowed"`
	Reason  *resources.Coding `firestore:"reason" json:"reason"`
}

type ConsultationObservation struct {
	Code           resources.Coding          `firestore:"code" json:"code"`
	Category       resources.Coding          `firestore:"category" json:"category"`
	Components     []ObservationComponent    `firestore:"components" json:"components"`
	Value          resources.Quantity        `firestore:"value" json:"value"`
	Interpretation resources.CodeableConcept `firestore:"interpretation" json:"interpretation"`
	Notes          []Note                    `firestore:"notes" json:"notes"`
}

type ObservationComponent struct {
	Code           resources.Coding          `firestore:"code" json:"code"`
	Value          resources.Quantity        `firestore:"value" json:"value"`
	Interpretation resources.CodeableConcept `firestore:"interpretation" json:"interpretation"`
}
