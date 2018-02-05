package model

// SpellEffectTypes is an array of SpellEffectType
// swagger:model
type SpellEffectTypes []*SpellEffectType

// SpellEffectType is the effect id type, e.g. SE_CURRENT_HP_ONCE
// http://wiki.eqemulator.org/p?Spell_Effect_IDs
// swagger:model
type SpellEffectType struct {
	ID   int64  `json:"ID,omitmpty"`
	Type int64  `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}
