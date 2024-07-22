package models // Post represents the post entity.
import "time"

type Post struct {
	PostID           int       `json:"postID" bson:"postID"`
	BoardID          int       `json:"boardID" bson:"boardID"`
	UserID           int       `json:"userID" bson:"userID"`
	ProfID           int       `json:"profID,omitempty" bson:"profID,omitempty"`
	Content          string    `json:"content" bson:"content"`
	CreationDateTime time.Time `json:"creationDateTime" bson:"creationDateTime"`
	EditDateTime     time.Time `json:"editDateTime,omitempty" bson:"editDateTime,omitempty"`
	NumOfReplies     int       `json:"numOfReplies,omitempty" bson:"numOfReplies,omitempty"`
}
