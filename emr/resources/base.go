package resources

type Resource interface {
	IsValid() bool
	GetResourceType() string
}

type BaseResource struct {
	ID string `json:"id,omitempty"`
	Identifier []Identifier `json:"identifier"`
}

func (r BaseResource) IsValid() bool { return true }
