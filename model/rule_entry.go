package model

import (
	"database/sql"
)

// RuleEntrys is an array of ruleEntry
// swagger:model
type RuleEntrys []*RuleEntry

//RuleEntry represents a rule entry.
// swagger:model
type RuleEntry struct {
	ValueFloat float64 `json:"valueFloat,omitempty"`
	ValueInt   int64   `json:"valueInt,omitempty"`
	Scope      string  `json:"scope,omitempty"`

	RuleID      int64          `json:"ruleID,omitempty" db:"ruleset_id"` //`ruleset_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Name        string         `json:"name,omitempty" db:"rule_name"`    //`rule_name` varchar(64) NOT NULL DEFAULT '',
	Value       string         `json:"value,omitempty" db:"rule_value"`  //`rule_value` varchar(30) NOT NULL DEFAULT '',
	Description sql.NullString `json:"description,omitempty" db:"notes"` //`notes` text,
}
