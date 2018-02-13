package model

//LootEntrys is an array of lootEntry
// swagger:model
type LootEntrys []*LootEntry

//LootEntry group together loot drops for npc drops
// swagger:model
type LootEntry struct {
	DropEntrys  []*LootDropEntry `json:"dropEntrys,omitempty"`
	LootID      int64            `json:"lootID" db:"loottable_id"`     //`loottable_id` int(11) unsigned NOT NULL DEFAULT '0',
	LootDropID  int64            `json:"lootDropID" db:"lootdrop_id"`  //`lootdrop_id` int(11) unsigned NOT NULL DEFAULT '0',
	Multiplier  int64            `json:"multiplier" db:"multiplier"`   //`multiplier` tinyint(2) unsigned NOT NULL DEFAULT '1',
	DropLimit   int64            `json:"dropLimit" db:"droplimit"`     //`droplimit` tinyint(2) unsigned NOT NULL DEFAULT '0',
	MinDrop     int64            `json:"minDrop" db:"mindrop"`         //`mindrop` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Probability float64          `json:"probability" db:"probability"` //`probability` float NOT NULL DEFAULT '100',
}
