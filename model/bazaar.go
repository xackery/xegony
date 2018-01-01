package model

import ()

//Bazaar is an item store on the website
type Bazaar struct {
	ID     int64  `json:"id"`
	ItemID int64  `json:"itemID" db:"itemid"`
	Name   string `json:"name"`
}
