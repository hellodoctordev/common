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

type Code interface {
	GetCodes() []Code
	String() string
}

func UnmarshalCode(c Code, raw string) Code {
	for _, code := range c.GetCodes() {
		if raw == code.String() {
			return code
		}
	}

	return nil
}
