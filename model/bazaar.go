package model

import ()

//Bazaar is an item store on the website
type Bazaar struct {
	Id     int64  `json:"id"`
	ItemId int64  `json:"itemID" db:"itemid"`
	Name   string `json:"name"`
}
