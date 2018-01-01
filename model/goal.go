package model

import ()

//Goal represnts goallist, used by the everquest Task system
type Goal struct {
	ListID  int64 `json:"listID" db:"listid"`
	EntryID int64 `json:"entryID" db:"entry"`
}
