package model

import ()

//Race holds data about races (including models)
// swagger:model
type Race struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Male    string `json:"male"`
	Female  string `json:"female"`
	Neutral string `json:"neutral"`
	Icon    string `json:"icon"`
}
