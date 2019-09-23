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
	ChatConsultation  FirestoreBillingSettingsConsultation `firestore:"chatConsultation" json:"chatConsultation"`
	VoiceConsultation FirestoreBillingSettingsConsultation `firestore:"voiceConsultation" json:"voiceConsultation"`
	VideoConsultation FirestoreBillingSettingsConsultation `firestore:"videoConsultation" json:"videoConsultation"`
}

type FirestoreBilling struct {
	Settings FirestoreBillingSettings `firestore:"settings" json:"settings"`
}

type FirestoreDoctorUser struct {
	Billing FirestoreBilling `firestore:"billing" json:"billing"`
}
