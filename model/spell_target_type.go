package model

import ()

// SpellTargetType is used to identify who are valid targets for a spell
// http://wiki.eqemulator.org/p?Target_Types&frm=spells_new
// swagger:model
type SpellTargetType struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
