package models

type HealthCareProfessional struct {
	ProfID       int    `json:"profID" bson:"profID"`
	FirstName    string `json:"firstName" bson:"firstName"`
	LastName     string `json:"lastName" bson:"lastName"`
	EmailAddress string `json:"emailAddress" bson:"emailAddress"`
	PhoneNum     string `json:"phoneNum" bson:"phoneNum"`
	WorkPhoneNum string `json:"workPhoneNum" bson:"workPhoneNum"`
	ProfBio      string `json:"profBio" bson:"profBio"`
	ABN          string `json:"ABN" bson:"ABN"`
	IsConsultant bool   `json:"isConsultant" bson:"isConsultant"`
}
