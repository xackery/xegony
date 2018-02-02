package model

import ()

// SpellTravelType identifies the type of travel a spell uses
// swagger:model
type SpellTravelType struct {
	ID   int64  `json:"ID"`
	Name string `json:"name"`
}
