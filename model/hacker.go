package model

import (
	"time"

	"database/sql"
)

//Hacker represents hacker reports inside everquest
// swagger:response
type Hacker struct {
	Zone      *Zone
	Account   *Account
	Character *Character

	ID            int64          `json:"ID" db:"id"`               //`id` int(4) NOT NULL AUTO_INCREMENT,
	AccountName   string         `json:"accountName" db:"account"` //`account` text NOT NULL,
	CharacterName string         `json:"name" db:"name"`           //`name` text NOT NULL,
	Hacked        string         `json:"hacked" db:"hacked"`       //`hacked` text NOT NULL,
	ZoneName      sql.NullString `json:"zone" db:"zone"`           //`zone` text,
	Date          time.Time      `json:"date" db:"date"`           //`date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
}
