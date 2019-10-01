package commonTypes

import "cloud.google.com/go/firestore"

type ConsultationDocument struct {
	Participants    []*firestore.DocumentRef `firestore:"participants"`
	ParticipantUIDs []string                 `firestore:"_participantUIDs"`
}
