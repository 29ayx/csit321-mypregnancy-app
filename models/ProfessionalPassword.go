package models

type ProfessionalPassword struct {
	ProfID   int    `json:"profID" bson:"profID"`
	PassHash string `json:"passHash" bson:"passHash"`
}
