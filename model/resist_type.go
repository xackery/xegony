package model

import ()

// ResistType is used to describe resist type.
// Used by spells and items.
// https://github.com/Shendare/EQArchitect/blob/master/lists/resistance_types.csv
// swagger:model
type ResistType struct {
	ID          int64  `json:"ID"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
