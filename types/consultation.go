package types

import "cloud.google.com/go/firestore"

type Consultation struct {
	Participants    []*firestore.DocumentRef `firestore:"participants"`
	ParticipantUIDs []string                 `firestore:"_participantUIDs"`
}
