package model

import ()

// SpellNumHitsType describes the type of numhits on a spell
// http://wiki.eqemulator.org/p?Numhit_Types&frm=spells_new
// swagger:model
type SpellNumHitsType struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
