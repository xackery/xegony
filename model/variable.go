package model

import (
	"time"
)

// Variables is an array of variable
// swagger:model
type Variables []*Variable

//Variable represents the zone table, Everquest is split into zones.
// swagger:model
type Variable struct {
	Name         string    `json:"name" db:"varname"`            //`varname` varchar(25) NOT NULL DEFAULT '',
	Value        string    `json:"value" db:"value"`             //`value` text NOT NULL,
	Description  string    `json:"description" db:"information"` //`information` text NOT NULL,
	ModifiedDate time.Time `json:"modifiedDate" db:"ts"`         //`ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
}
