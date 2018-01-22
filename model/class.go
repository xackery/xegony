package model

import ()

//Class represents classes in EQ
// swagger:response
type Class struct {
	ID        int64  `json:"id"`
	Bit       int64  `json:"bit"`
	ShortName string `json:"shortName"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
}
