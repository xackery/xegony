package model

//Forums is an array of forum
// swagger:model
type Forums []*Forum

//Forum is the parent of topics, which groups posts together into sections
// swagger:model
type Forum struct {
	ID          int64  `json:"ID,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	UserID      int64  `json:"userID" db:"user_id"`
	Sort        int64  `json:"sort" db:"sort"`
	Description string `json:"description" db:"description"`
	Icon        string `json:"icon" db:"icon"`
}
