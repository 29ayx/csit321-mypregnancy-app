package models

type UserPassword struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	PassHash string `json:"passhash" bson:"passhash"`
}
