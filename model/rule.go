package model

import (
	"database/sql"
)

//Rule represents the zone table, Everquest is split into zones.
// swagger:model
type Rule struct {
	ValueFloat float64 `json:"valueFloat"`
	ValueInt   int64   `json:"valueInt"`
	Scope      string  `json:"scope"`

	RulesetID int64          `json:"rulesetID" db:"ruleset_id"` //`ruleset_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Name      string         `json:"ruleName" db:"rule_name"`   //`rule_name` varchar(64) NOT NULL DEFAULT '',
	Value     string         `json:"ruleValue" db:"rule_value"` //`rule_value` varchar(30) NOT NULL DEFAULT '',
	Notes     sql.NullString `json:"notes" db:"notes"`          //`notes` text,
}
