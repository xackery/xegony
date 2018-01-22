package model

import ()

//Topic is the entries inside forum, grouping posts
// swagger:model
type Topic struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	ForumID int64  `json:"forumID" db:"forum_id"`
	Icon    string `json:"icon"`
}
