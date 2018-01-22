package model

import ()

//Skill represents skills inside everquest
// swagger:model
type Skill struct {
	ID   int64  `json:"ID" db:"id"`
	Name string `json:"name" db:"name"`
	Type int64  `json:"type" db:"type"`
}
