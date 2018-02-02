package model

import ()

// SpellReagent represents the no expend spell reagents used for a spell.
// Ranges from 1 to 4
// swagger:model
type SpellReagent struct {
	ID      int64 `json:"ID"`
	SpellID int64 `json:"spellID"`
	ItemID  int64 `json:"itemID"` //`NoexpendReagent1` int(11) NOT NULL DEFAULT '-1',
}
