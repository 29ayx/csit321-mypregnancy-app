package models

type HealthRecord struct {
	UserID         int    `json:"userID" bson:"userID"`
	Age            int    `json:"age" bson:"age"`
	Height         int    `json:"height" bson:"height"`
	Weight         int    `json:"weight" bson:"weight"`
	PregnancyPhase string `json:"pregnancyPhase" bson:"pregnancyPhase"`
	WeeksAlong     int    `json:"weeksAlong" bson:"weeksAlong"`
}
