package resources

type Questionnaire struct{}

func (q Questionnaire) IsFamilyMemberHistoryDefinition()    {}
func (q Questionnaire) IsCarePlanDefinition()               {}
func (q Questionnaire) IsCarePlanActivityDetailDefinition() {}

type PlanDefinition struct{}

func (pd PlanDefinition) IsCarePlanDefinition()               {}
func (pd PlanDefinition) IsCarePlanActivityDetailDefinition() {}

type ActivityDefinition struct{}

func (ad ActivityDefinition) IsCarePlanActivityDetailDefinition() {}
