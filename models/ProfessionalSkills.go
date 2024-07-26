package models

type ProfessionalSkills struct {
	ProfID int    `json:"profID" bson:"profID"`
	Skill  string `json:"skill" bson:"skill"`
}
