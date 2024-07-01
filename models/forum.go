package models

import "time"

type Forum struct {
    ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
    Title     string    `json:"title" bson:"title"`
    Content   string    `json:"content" bson:"content"`
    UserID    string    `json:"user_id" bson:"user_id"`
    CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
