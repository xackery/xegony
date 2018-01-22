package model

import ()

//LootDropEntry groups together items that npcs drop
// swagger:model
type LootDropEntry struct {
	LootdropID     int64   `json:"lootdropId" db:"lootdrop_id"`         //`lootdrop_id` int(11) unsigned NOT NULL DEFAULT '0',
	ItemID         int64   `json:"itemID" db:"item_id"`                 //`item_id` int(11) NOT NULL DEFAULT '0',
	ItemCharges    int64   `json:"itemCharges" db:"item_charges"`       //`item_charges` smallint(2) unsigned NOT NULL DEFAULT '1',
	EquipItem      int64   `json:"equipItem" db:"equip_item"`           //`equip_item` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Chance         float64 `json:"chance" db:"chance"`                  //`chance` float NOT NULL DEFAULT '1',
	DisabledChance float64 `json:"disabledChance" db:"disabled_chance"` //`disabled_chance` float NOT NULL DEFAULT '0',
	Minlevel       int64   `json:"minlevel" db:"minlevel"`              //`minlevel` tinyint(3) NOT NULL DEFAULT '0',
	Maxlevel       int64   `json:"maxlevel" db:"maxlevel"`              //`maxlevel` tinyint(3) NOT NULL DEFAULT '127',
	Multiplier     int64   `json:"multiplier" db:"multiplier"`          //`multiplier` tinyint(2) unsigned NOT NULL DEFAULT '1',
}
