package models

import "time"

type ConsultationRequests struct {
	RequestID            int       `json:"requestID" bson:"requestID"`
	UserID               int       `json:"userID" bson:"userID"`
	ProfID               int       `json:"profID" bson:"profID"`
	Description          string    `json:"description,omitempty" bson:"description,omitempty"`
	CommunicationType    string    `json:"communicationType" bson:"communicationType"`
	ConsultationDateTime time.Time `json:"consultationDateTime" bson:"consultationDateTime"`
	Status               string    `json:"status" bson:"status"`
	PreferredGender      string    `json:"preferredGender,omitempty" bson:"preferredGender,omitempty"`
}
