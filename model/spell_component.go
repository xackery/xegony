package model

import ()

// SpellComponent represents items required or utilized for the spell.
// ranges from 1 to 4
// swagger:model
type SpellComponent struct {
	SpellID int64 `json:"spellID"`
	ID      int64 `json:"ID"`    //`components1` int(11) NOT NULL DEFAULT '-1',
	Count   int64 `json:"count"` //`component_counts1` int(11) NOT NULL DEFAULT '1',
}
