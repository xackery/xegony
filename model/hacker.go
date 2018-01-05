package model

import (
	"time"

	"database/sql"
)

//Hacker represents hacker reports inside everquest
type Hacker struct {
	Account   *Account
	ID        int64          `json:"ID" db:"id"`             //`id` int(4) NOT NULL AUTO_INCREMENT,
	AccountID string         `json:"accountID" db:"account"` //`account` text NOT NULL,
	Name      string         `json:"name" db:"name"`         //`name` text NOT NULL,
	Hacked    string         `json:"hacked" db:"hacked"`     //`hacked` text NOT NULL,
	Zone      sql.NullString `json:"zone" db:"zone"`         //`zone` text,
	Date      time.Time      `json:"date" db:"date"`         //`date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
}
