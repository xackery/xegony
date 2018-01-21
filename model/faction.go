package model

import ()

//Faction represents everquest factions
type Faction struct {
	ID        int64  `json:"id"`
	Base      int64  `json:"base"`
	Name      string `json:"name"`
	CleanName string `json:"cleanName"`
}
