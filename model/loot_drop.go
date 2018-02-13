package model

//LootDrops is an array of lootDrop
// swagger:model
type LootDrops []*LootDrop

//LootDrop is the parent of loot drop entries
// swagger:model
type LootDrop struct {
	Entries []*LootDropEntry `json:"entries,omitempty"`

	ID   int64  `json:"ID" db:"id"`     //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	Name string `json:"name" db:"name"` //`name` varchar(255) NOT NULL DEFAULT '',
}
