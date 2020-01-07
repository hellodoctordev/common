package resources

type Encounter struct{}

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
