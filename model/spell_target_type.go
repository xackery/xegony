package model

// SpellTargetTypes is an array of SpellTargetType
// swagger:model
type SpellTargetTypes []*SpellTargetType

// SpellTargetType is used to identify who are valid targets for a spell
// http://wiki.eqemulator.org/p?Target_Types&frm=spells_new
// swagger:model
type SpellTargetType struct {
	ID   int64  `json:"ID,omitempty"`
	Name string `json:"name,omitempty"`
}
