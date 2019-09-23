package types

type FirestoreBillingSettingsConsultationPrice struct {
	Amount   int64  `firestore:"amount"`
	Currency string `firestore:"currency"`
}

type FirestoreBillingSettingsConsultation struct {
	BillingType string `firestore:"billingType"`
	IsEnabled   bool   `firestore:"isEnabled"`
}

type FirestoreBillingSettings struct {
	ChatConsultations  FirestoreBillingSettingsConsultation `firestore:"chatConsultations"`
	VoiceConsultations FirestoreBillingSettingsConsultation `firestore:"voiceConsultations"`
	VideoConsultations FirestoreBillingSettingsConsultation `firestore:"videoConsultations"`
}

type FirestoreBilling struct {
	Billing FirestoreBillingSettings `firestore:"billing"`
}

type FirestoreDoctorUser struct {
	Billing FirestoreBilling `firestore:"billing"`
}
