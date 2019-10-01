package commonTypes

import "cloud.google.com/go/firestore"

// TODO Perhaps get rid of this entirely
type ConsultationDocument struct {
	Participants    []*firestore.DocumentRef `firestore:"participants"`
	ParticipantUIDs []string                 `firestore:"_participantUIDs"`
}
