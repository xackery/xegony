package model

// Sizes represents an array of size
// swagger:model
type Sizes []*Size

//Size represents sizes inside everquest
// swagger:model
type Size struct {
	ID   int64  `json:"ID" db:"id"`
	Name string `json:"name" db:"name"`
	Icon string `json:"icon" db:"icon"`
}
