package model

import ()

// DBStr represents the dbstr_us.txt file.
// swagger:model
type DBStr struct {
	Type *DBStrType `json:"type"`

	ID     int64  `json:"ID"`
	TypeID int64  `json:"typeID"`
	Name   string `json:"name"`
}
