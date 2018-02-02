package model

import ()

// SpellEffectType is the effect id type, e.g. SE_CURRENT_HP_ONCE
// http://wiki.eqemulator.org/p?Spell_Effect_IDs
// swagger:model
type SpellEffectType struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
