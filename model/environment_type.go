package model

import ()

// EnvironmentType describes the type of environment this value works in.
// Used by spells.
// swagger:model
type EnvironmentType struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
