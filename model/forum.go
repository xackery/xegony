package model

//Forums is an array of forum
// swagger:model
type Forums []*Forum

//Forum is the parent of topics, which groups posts together into sections
// swagger:model
type Forum struct {
	ID          int64  `json:"id,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	OwnerID     int64  `json:"ownerId" db:"owner_id"`
	Description string `json:"description" db:"description"`
	Icon        string `json:"icon" db:"icon"`
}
