package model

import ()

//User is a login authentication created by Xegony, binding an account ID
type User struct {
	Status int64 `json:"status"`

	ID          int64  `json:"id"`
	Name        string `json:"name"`
	AccountID   int64  `json:"accountID" db:"account_id"`
	CharacterID int64  `json:"characterID" db:"character_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

//IsAdmin returns an error if not admin
func (u *User) IsAdmin() (err error) {
	if u.Status < 200 {
		err = &ErrPermission{
			Message: "Admin level access required",
		}
		return
	}
	return
}

//IsGuide returns an err if not guide
func (u *User) IsGuide() (err error) {
	if u.Status < 100 {
		err = &ErrPermission{
			Message: "Guide level access required",
		}
		return
	}
	return
}
