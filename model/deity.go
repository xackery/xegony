package model

//Deitys is an array of deity
// swagger:model
type Deitys []*Deity

//Deity represents deities in EQ
// swagger:model
type Deity struct {
	ID        int64  `json:"ID,omitempty" yaml:"ID"`
	Bit       int64  `json:"bit" yaml:"bit"`
	ShortName string `json:"shortName" yaml:"shortName"`
	SpellID   int64  `json:"spellID,omitempty" yaml:"spellID"`
	Name      string `json:"name" yaml:"name"`
	Icon      string `json:"icon" yaml:"icon"`
}
