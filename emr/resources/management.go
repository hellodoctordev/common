package resources

import (
	"errors"
	"fmt"
)

type EncounterStatus string

const (
	EncounterStatusPlanned    EncounterStatus = "planned"
	EncounterStatusArrived    EncounterStatus = "arrived"
	EncounterStatusTriaged    EncounterStatus = "triaged"
	EncounterStatusInProgress EncounterStatus = "in-progress"
	EncounterStatusOnLeave    EncounterStatus = "on-leave"
	EncounterStatusFinished   EncounterStatus = "finished"
)

func (e EncounterStatus) GetCodes() []Code {
	return []Code{EncounterStatusPlanned, EncounterStatusArrived, EncounterStatusTriaged, EncounterStatusInProgress, EncounterStatusOnLeave, EncounterStatusFinished}
}
func (e EncounterStatus) String() string { return string(e) }
func (e *EncounterStatus) UnmarshalJSON(b []byte) error {
	raw := string(b)[1 : len(b)-1]
	unmarshalled := UnmarshalCode(e, raw)

	if unmarshalled != nil {
		*e = unmarshalled.(EncounterStatus)
		return nil
	} else {
		return errors.New(fmt.Sprintf("%s is not a valid EncounterStatus", b))
	}
}

type EncounterStatusHistory struct {
	Status EncounterStatus `json:"status"`
	Period Period          `json:"period"`
}

type EncounterParticipant struct {
	Type       []CodeableConcept `json:"type"`
	Period     *Period           `json:"period,omitempty"`
	Individual *Reference        `json:"individual,omitempty"`
}

type EncounterDiagnosis struct {
	Condition *CodeableConcept `json:"condition,omitempty"`
	Role      *Reference       `json:"role,omitempty"`
	Rank      float64          `json:"rank,omitempty"`
}

type Encounter struct {
	BaseResource
	Status        EncounterStatus          `json:"status"`
	StatusHistory []EncounterStatusHistory `json:"statusHistory,omitempty"`
	Class         Coding                   `json:"class,omitempty"`
	Type          []CodeableConcept        `json:"type,omitempty"`
	Priority      *CodeableConcept         `json:"priority,omitempty"`
	Subject       *Reference               `json:"subject,omitempty"`
	EpisodeOfCare []Reference              `json:"episodeOfCare,omitempty"`
	Participant   []EncounterParticipant   `json:"participant,omitempty"`
	Appointment   *Reference               `json:"appointment,omitempty"`
	Period        *Period                  `json:"period,omitempty"`
	Reason        []CodeableConcept        `json:"reason,omitempty"`
	Diagnosis     []EncounterDiagnosis     `json:"diagnosis,omitempty"`
	PartOf        *Reference               `json:"partOf,omitempty"`
}

func (e Encounter) GetResourceType() string      { return "Encounter" }
func (e Encounter) IsConditionContext()          {}
func (e Encounter) IsProcedureContext()          {}
func (e Encounter) IsClinicalImpressionContext() {}
func (e Encounter) IsObservationContext()        {}
func (e Encounter) IsCarePlanContext()           {}

type EpisodeOfCare struct{}

func (e EpisodeOfCare) IsConditionContext()          {}
func (e EpisodeOfCare) IsProcedureContext()          {}
func (e EpisodeOfCare) IsClinicalImpressionContext() {}
func (e EpisodeOfCare) IsObservationContext()        {}
func (e EpisodeOfCare) IsCarePlanContext()           {}
