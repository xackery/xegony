package model

import ()

// DBStrType represents the category types found inside dbstr, and is refered to in entries.
// http://wiki.eqemulator.org/p?dbstr_us.txt
// swagger:model
type DBStrType struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
