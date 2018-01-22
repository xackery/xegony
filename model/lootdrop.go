package model

import ()

//LootDrop is the parent of loot drop entries
// swagger:response
type LootDrop struct {
	ID      int64  `json:"id" db:"id"`     //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	Name    string `json:"name" db:"name"` //`name` varchar(255) NOT NULL DEFAULT '',
	Entries []*LootDropEntry
}
