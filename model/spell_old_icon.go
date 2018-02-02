package model

import ()

// SpellOldIcon is used by spells to describe an icon for older clients (pre luclin)
// swagger:model
type SpellOldIcon struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
