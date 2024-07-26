package models

import "time"

type Comment struct {
	CommentID        int       `json:"commentID" bson:"commentID"`
	PostID           int       `json:"postID" bson:"postID"`
	UserID           int       `json:"userID,omitempty" bson:"userID,omitempty"`
	ProfID           int       `json:"profID,omitempty" bson:"profID,omitempty"`
	Content          string    `json:"content" bson:"content"`
	CreationDateTime time.Time `json:"creationDateTime" bson:"creationDateTime"`
}
