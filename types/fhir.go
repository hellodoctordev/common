package types

import (
	"cloud.google.com/go/firestore"
	"fmt"
)

type FHIRResource struct {
	ResourceType string                 `firestore:"resourceType" json:"resourceType"`
	ResourceID   string                 `firestore:"resourceID" json:"resourceID"`
	Target       *firestore.DocumentRef `firestore:"target" json:"target"`
}

func (r FHIRResource) GetReferenceString() string {
	return fmt.Sprintf("%s/%s", r.ResourceType, r.ResourceID)
}
