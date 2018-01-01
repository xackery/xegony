package model

import ()

//User is a login authentication created by Xegony, binding an account ID
type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	AccountID   int64  `json:"accountID" db:"account_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"isAdmin"`
	IsModerator bool   `json:"isModerator"`
}
