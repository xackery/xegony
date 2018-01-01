package model

import ()

//Forum is the parent of topics, which groups posts together into sections
type Forum struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	OwnerID     int64  `json:"ownerId" db:"owner_id"`
	Description string `json:"description"`
	Icon        string `json:"icon" db:"icon"`
}
