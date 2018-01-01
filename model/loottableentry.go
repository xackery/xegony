package model

import ()

//LootTableEntry group together loot drops for npc drops
type LootTableEntry struct {
	LoottableId int64   `json:"loottableId" db:"loottable_id"` //`loottable_id` int(11) unsigned NOT NULL DEFAULT '0',
	LootdropId  int64   `json:"lootdropId" db:"lootdrop_id"`   //`lootdrop_id` int(11) unsigned NOT NULL DEFAULT '0',
	Multiplier  int64   `json:"multiplier" db:"multiplier"`    //`multiplier` tinyint(2) unsigned NOT NULL DEFAULT '1',
	Droplimit   int64   `json:"droplimit" db:"droplimit"`      //`droplimit` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Mindrop     int64   `json:"mindrop" db:"mindrop"`          //`mindrop` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Probability float64 `json:"probability" db:"probability"`  //`probability` float NOT NULL DEFAULT '100',
}
