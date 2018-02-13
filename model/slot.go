package model

// Slots represents an array of slot
// swagger:model
type Slots []*Slot

//Slot represents item slots
// swagger:model
type Slot struct {
	ID        int64  `json:"ID,omitempty"`
	Bit       int64  `json:"bit,omitempty"`
	Name      string `json:"name,omitempty"`
	ShortName string `json:"shortName,omitempty"`
	Icon      string `json:"icon,omitempty"`
}
