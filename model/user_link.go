package model

import "time"

// UserLinks is an array of account
// swagger:model
type UserLinks []*UserLink

// UserLink represents the user for this application, and binds to accountID
//
// swagger:model
type UserLink struct {
	Account   *Account   `json:"account,omitempty"`
	Character *Character `json:"character,omitempty"`

	ID          int64     `json:"ID,omitempty" db:"id"`
	Link        string    `json:"link,omitempty" db:"link"`
	AccountID   int64     `json:"accountID,omitempty" db:"account_id"`
	CharacterID int64     `json:"characterID,omitempty" db:"character_id"`
	CreateDate  time.Time `json:"createDate,omitempty" db:"create_date"`
}
