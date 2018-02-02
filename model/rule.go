package model

import ()

// Rules is an array of rule
// swagger:model
type Rules []*Rule

//Rule represents the zone table, Everquest is split into zones.
// swagger:model
type Rule struct {
	ID   int64  `json:"ID,omitempty" db:"ruleset_id"` //`ruleset_id` tinyint(3) unsigned NOT NULL DEFAULT '0',
	Name string `json:"ruleName,omitempty" db:"name"` //`name` varchar(255) NOT NULL DEFAULT '',
}
