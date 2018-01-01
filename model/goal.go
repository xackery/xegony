package model

import ()

//Goal represnts goallist, used by the everquest Task system
type Goal struct {
	ListId  int64 `json:"listID" db:"listid"`
	EntryId int64 `json:"entryID" db:"entry"`
}
