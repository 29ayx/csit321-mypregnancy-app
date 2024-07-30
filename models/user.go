package models

type User struct {
	ID                string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName         string `json:"firstname" bson:"firstname"`
	LastName          string `json:"lastname" bson:"lastname"`
	Email             string `json:"email" bson:"email"`
	PhoneNum          int    `json:"phonenum" bson:"phonenum"`
	UserBio           string `json:"userbio" bson:"userbio"`
	PassHash          string `json:"passhash" bson:"passhash"`
	IsExpectingMother bool   `json:"isexpectingmother" bson:"isexpectingmother"`
}
