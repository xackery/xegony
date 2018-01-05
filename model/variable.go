package model

import (
	"time"
)

//Variable represents the zone table, Everquest is split into zones.
type Variable struct {
	Name        string    `json:"name" db:"varname"`            //`varname` varchar(25) NOT NULL DEFAULT '',
	Value       string    `json:"value" db:"value"`             //`value` text NOT NULL,
	Information string    `json:"information" db:"information"` //`information` text NOT NULL,
	Ts          time.Time `json:"ts" db:"ts"`                   //`ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
}

//ValueParse returns parsed values
func (c *Variable) ValueParse() string {
	return c.Value
}
