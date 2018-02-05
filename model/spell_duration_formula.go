package model

// SpellDurationFormulas is an array of SpellDurationFormula
// swagger:model
type SpellDurationFormulas []*SpellDurationFormula

// SpellDurationFormula is used in various areas to represent spells
// swagger:model
type SpellDurationFormula struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
