package model

import ()

// SpellDeity represents the deities field inside spells_new.
// There is a range of 0 to 16 to represent deities
// swagger:model
type SpellDeity struct {
	SpellID int64 `json:"spellID"`
	ID      int64 `json:"ID"`
	Deity   int64 `json:"deity"` //`deities0` int(11) NOT NULL DEFAULT '0',
}
