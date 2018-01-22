package model

import ()

//SpellEffect stores details about a spell's effects
// swagger:response
type SpellEffect struct {
	ID       int64
	Type     int64
	TypeName string

	FormulaName string
	Name        string
	BaseValue   int64
	Description string
}
