package models

import "time"

// HealthJournal represents the health journal entity.
type HealthJournal struct {
	JournalID   int       `json:"journalID" bson:"journalID"`
	UserID      int       `json:"userID" bson:"userID"`
	EntryDate   time.Time `json:"entryDate" bson:"entryDate"`
	Feeling     string    `json:"feeling,omitempty" bson:"feeling,omitempty"`
	Gratitudes  string    `json:"gratitudes,omitempty" bson:"gratitudes,omitempty"`
	SelfCare    string    `json:"selfCare,omitempty" bson:"selfCare,omitempty"`
	Thoughts    string    `json:"thoughts,omitempty" bson:"thoughts,omitempty"`
	DailyRating int       `json:"dailyRating,omitempty" bson:"dailyRating,omitempty"`
}
