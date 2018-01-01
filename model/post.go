package model

import ()

//Post are topic entries for forums
type Post struct {
	ID      int64  `json:"id"`
	Body    string `json:"body"`
	TopicID int64  `json:"topicID" db:"topic_id"`
	OwnerID int64  `json:"ownerId" db:"owner_id"`
}
