package model

import "time"

// UserAccounts is an array of account
// swagger:model
type UserAccounts []*UserAccount

// UserAccount represents the accounts linked to a user
//
// swagger:model
type UserAccount struct {
	Account   *Account   `json:"account,omitempty"`
	Character *Character `json:"character,omitempty"`

	ID          int64     `json:"ID,omitempty" db:"id"`
	UserID      int64     `json:"userID,omitempty" db:"user_id"`
	AccountID   int64     `json:"accountID,omitempty" db:"account_id"`
	CharacterID int64     `json:"characterID,omitempty" db:"character_id"`
	CreateDate  time.Time `json:"createDate,omitempty" db:"create_date"`
}
