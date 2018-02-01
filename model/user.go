package model

import ()

// Users is an array of account
// swagger:model
type Users []*User

// User represents the user for this application, and binds to accountID
//
// swagger:model
type User struct {
	Accounts         []*Account   `json:"accounts,omitempty"`
	Characters       []*Character `json:"characters,omitempty"`
	PrimaryAccount   *Account     `json:"primaryAccount,omitempty"`
	PrimaryCharacter *Character   `json:"primaryCharacter,omitempty"`

	ID                 int64  `json:"ID,omitempty" db:"id"`
	DisplayName        string `json:"displayName,omitempty" db:"display_name"`
	PrimaryAccountID   int64  `json:"primaryAccountID,omitempty" db:"primary_account_id"`
	PrimaryCharacterID int64  `json:"primaryCharacterID,omitempty" db:"primary_character_id"`
	Email              string `json:"email,omitempty" db:"email"`
	Password           string `json:"password,omitempty" db:"password"`
	GoogleToken        string `json:"googleToken,omitempty" db:"google_token"`
}

//IsAdmin returns an error if not admin
func (u *User) IsAdmin() (err error) {
	var highestStatus int64
	for _, account := range u.Accounts {
		if account.Status > highestStatus {
			highestStatus = account.Status
		}
	}
	if u.PrimaryAccount.Status > highestStatus {
		highestStatus = u.PrimaryAccount.Status
	}
	if highestStatus < 200 {
		err = &ErrPermission{
			Message: "Admin level access required",
		}
		return
	}
	return
}

//IsGuide returns an err if not guide
func (u *User) IsGuide() (err error) {
	var highestStatus int64
	for _, account := range u.Accounts {
		if account.Status > highestStatus {
			highestStatus = account.Status
		}
	}
	if u.PrimaryAccount.Status > highestStatus {
		highestStatus = u.PrimaryAccount.Status
	}
	if highestStatus < 100 {
		err = &ErrPermission{
			Message: "Guide level access required",
		}
		return
	}
	return
}
