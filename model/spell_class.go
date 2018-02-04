package model

import ()

// SpellClass stores minimum levels of each class
// ranges 1 to 60
// swagger:model
type SpellClass struct {
	ClassID int64 `json:"classID,omitempty"`
	Level   int64 `json:"level,omitempty"`
}
