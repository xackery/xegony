package model

// SpellTravelTypes is an array of SpellTravelType
// swagger:model
type SpellTravelTypes []*SpellTravelType

// SpellTravelType identifies the type of travel a spell uses
// swagger:model
type SpellTravelType struct {
	ID   int64  `json:"ID,omitempty"`
	Name string `json:"name,omitempty"`
}
