package model

import (
	"time"
)

//Bazaar is an item store on the website
// swagger:model
type Bazaar struct {
	Item       *Item
	ID         int64     `json:"id" db:"id"`
	ItemID     int64     `json:"itemID" db:"itemid"`
	AccountID  int64     `json:"accountID" db:"accountid"`
	Price      int64     `json:"price" db:"price"`
	CreateDate time.Time `json:"createDate" db:"createdate"`
}
