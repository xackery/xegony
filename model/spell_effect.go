package model

import ()

// SpellEffect stores details about a spell's effects
// There is 1 to 12 range of effects per spell
// swagger:model
type SpellEffect struct {
	Formula *SpellFormula `json:"formula"`

	SpellID   int64 `json:"spellID"`
	ID        int64 `json:"ID"`
	FormulaID int64 `json:"formulaID"`  //`formula1` int(11) NOT NULL DEFAULT '100',
	Base      int64 `json:"baseValue"`  //`effect_base_value1` int(11) NOT NULL DEFAULT '100',
	Limit     int64 `json:"limitValue"` //`effect_limit_value1` int(11) NOT NULL DEFAULT '0',
	Max       int64 `json:"max"`        //`max1` int(11) NOT NULL DEFAULT '0',
	EffectID  int64 `json:"effectID"`   //`effectid1` int(11) NOT NULL DEFAULT '254',
	Class     int64 `json:"class"`      //`classes1` int(11) NOT NULL DEFAULT '255',
}
