package models

type ConsultationNotes struct {
	RequestID int    `json:"requestID" bson:"requestID"`
	Notes     string `json:"notes" bson:"notes"`
}
