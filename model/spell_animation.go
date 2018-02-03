package model

import ()

//SpellAnimations is an array of spellAnimation
// swagger:model
type SpellAnimations []*SpellAnimation

// SpellAnimation represents animations on spells
// http://www.eqemulator.org/forums/showthread.php?t=30731
// swagger:model
type SpellAnimation struct {
	Type *SpellAnimationType `json:"type,omitempty" yaml:"-"`

	ID     int64  `json:"ID,omitempty" yaml:"ID"`
	TypeID int64  `json:"typeID,omitempty" yaml:"typeID"`
	Name   string `json:"name,omitempty" yaml:"name"`
}
