package models

type User struct {
	ID                string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName         string `json:"firstname" bson:"firstname"`
	LastName          string `json:"lastname" bson:"lastname"`
	Email             string `json:"email" bson:"email"`
	PhoneNum          int    `json:"phonenum" bson:"phonenum"`
	UserBio           string `json:"userbio" bson:"userbio"`
	IsExpectingMother bool   `json:"isexpectingmother" bson:"isexpectingmother"`
	IsProfessional    bool   `json:"isprofessional" bson:"isprofessional"`
}
