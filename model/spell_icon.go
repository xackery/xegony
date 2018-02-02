package model

import ()

// SpellIcon is used by spells to describe an icon
// swagger:model
type SpellIcon struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
