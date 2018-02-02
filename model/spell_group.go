package model

import ()

// SpellGroup is used to group together spells of similar types.
// http://wiki.eqemulator.org/p?Spell_Groups&frm=spells_new
// swagger:model
type SpellGroup struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
