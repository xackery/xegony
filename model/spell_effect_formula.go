package model

// SpellEffectFormulas is an array of SpellEffectFormula
// swagger:model
type SpellEffectFormulas []*SpellEffectFormula

// SpellEffectFormula is used in various areas to represent spells
// swagger:model
type SpellEffectFormula struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
