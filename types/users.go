package types

type FirestoreBillingSettingsConsultationPrice struct {
	Amount   int64  `firestore:"amount" json:"amount"`
	Currency string `firestore:"currency" json:"currency"`
}

type FirestoreBillingSettingsConsultation struct {
	BillingPrice FirestoreBillingSettingsConsultationPrice `firestore:"billingPrice" json:"billingPrice"`
	BillingType  string                                    `firestore:"billingType" json:"billingType"`
	IsEnabled    bool                                      `firestore:"isEnabled" json:"isEnabled"`
}

type FirestoreBillingSettings struct {
	ChatConsultations  FirestoreBillingSettingsConsultation `firestore:"chatConsultations" json:"chatConsultations"`
	VoiceConsultations FirestoreBillingSettingsConsultation `firestore:"voiceConsultations" json:"voiceConsultations"`
	VideoConsultations FirestoreBillingSettingsConsultation `firestore:"videoConsultations" json:"videoConsultations"`
}

type FirestoreBilling struct {
	Settings FirestoreBillingSettings `firestore:"settings" json:"settings"`
}

type FirestoreDoctorUser struct {
	Billing FirestoreBilling `firestore:"billing" json:"billing"`
}
