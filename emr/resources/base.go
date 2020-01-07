package resources

import (
	"encoding/json"
)

type Resource interface {
	IsValid() bool
	ToMap() map[string]interface{}
	GetResourceType() string
}

type BaseResource struct {
	Identifier []Identifier `json:"identifier"`
}

func (r BaseResource) IsValid() bool { return true }

func (r BaseResource) ToMap() map[string]interface{} {
	resourceBytes, _ := json.Marshal(r)

	var asMap map[string]interface{}

	_ = json.Unmarshal(resourceBytes, &asMap)

	return asMap
}
