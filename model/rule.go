package model

import (
	"database/sql"
	"strings"
)

//Rule represents the zone table, Everquest is split into zones.
type Rule struct {
	RulesetID int64          `json:"rulesetID" db:"ruleset_id"` //`ruleset_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Name      string         `json:"ruleName" db:"rule_name"`   //`rule_name` varchar(64) NOT NULL DEFAULT '',
	Value     string         `json:"ruleValue" db:"rule_value"` //`rule_value` varchar(30) NOT NULL DEFAULT '',
	Notes     sql.NullString `json:"notes" db:"notes"`          //`notes` text,
}

//ValueParse returns parsed values
func (c *Rule) ValueParse() string {
	return c.Value
}

//Scope is what scope a rule is for
func (c *Rule) Scope() string {
	scope := ""
	if strings.Contains(c.Name, ":") {
		scope = c.Name[0:strings.Index(c.Name, ":")]
	}
	return scope
}
