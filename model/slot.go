package model

import ()

//Slot represents item slots
// swagger:model
type Slot struct {
	ID        int64  `json:"ID"`
	BitID     int64  `json:"bitID"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}
