package model

// ZoneExpansions is an array of zoneExpansion
// swagger:model
type ZoneExpansions []*ZoneExpansion

// ZoneExpansion represents the zoneExpansion table, Everquest is split into zoneExpansions
// swagger:model
type ZoneExpansion struct {
	ID        int64  `json:"ID,omitempty" yaml:"ID,omitempty" db:"id"`
	ShortName string `json:"shortName,omitempty" yaml:"shortName,omitempty" db:"short_name"`
	Name      string `json:"name,omitempty" yaml:"name,omitempty" db:"name"`
	Bit       int64  `json:"bit,omitempty" yaml:"bit,omitempty" db:"bit"`
}
