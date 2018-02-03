package model

import ()

//SpellAnimationTypes is an array of spellAnimationType
// swagger:model
type SpellAnimationTypes []*SpellAnimationType

// SpellAnimationType represents animation types on spells
// http://www.eqemulator.org/forums/showthread.php?t=30731
// swagger:model
type SpellAnimationType struct {
	ID        int64  `json:"ID,omitempty" yaml:"ID"`
	Name      string `json:"name,omitempty" yaml:"name"`
	ShortName string `json:"shortName,omitempty" yaml:"shortName"`
}
