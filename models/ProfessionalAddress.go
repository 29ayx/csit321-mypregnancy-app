package models

type ProfessionalAddress struct {
	ProfID     int    `json:"profID" bson:"profID"`
	UnitNumber string `json:"unitNumber" bson:"unitNumber"`
	StreetNum  string `json:"streetNum" bson:"streetNum"`
	StreetName string `json:"streetName" bson:"streetName"`
	Suburb     string `json:"suburb" bson:"suburb"`
	State      string `json:"state" bson:"state"`
}
