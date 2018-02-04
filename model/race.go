package model

// Races is an array of Race
// swagger:model
type Races []*Race

// Race holds data about races (including models)
// swagger:model
type Race struct {
	ID      int64  `json:"ID,omitempty" yaml:"ID"`
	Bit     int64  `json:"bit,omitempty" yaml:"bit"`
	Name    string `json:"name,omitempty" yaml:"name"`
	Male    string `json:"male,omitempty" yaml:"male"`
	Female  string `json:"female,omitempty" yaml:"female"`
	Neutral string `json:"neutral,omitempty" yaml:"neutral"`
	Icon    string `json:"icon,omitempty" yaml:"icon"`
}
