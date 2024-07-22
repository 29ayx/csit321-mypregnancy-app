package models

type ForumBoard struct {
	BoardID     int    `json:"boardID" bson:"boardID"`
	Topic       string `json:"topic" bson:"topic"`
	Description string `json:"description" bson:"description"`
}
