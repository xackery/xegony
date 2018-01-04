package model

import ()

//MerchantEntry group together loot drops for npc drops
type MerchantEntry struct {
	Item *Item

	MerchantID      int64 `json:"merchantID" db:"merchantid"`             //`merchantid` int(11) NOT NULL DEFAULT '0',
	Slot            int64 `json:"slot" db:"slot"`                         //`slot` int(11) NOT NULL DEFAULT '0',
	ItemID          int64 `json:"itemID" db:"item"`                       //`item` int(11) NOT NULL DEFAULT '0',
	FactionRequired int64 `json:"factionRequired" db:"faction_required"`  //`faction_required` smallint(6) NOT NULL DEFAULT '-100',
	LevelRequired   int64 `json:"levelRequired" db:"level_required"`      //`level_required` tinyint(3) unsigned NOT NULL DEFAULT '0',
	AltCurrencyCost int64 `json:"altCurrencyCost" db:"alt_currency_cost"` //`alt_currency_cost` smallint(5) unsigned NOT NULL DEFAULT '0',
	ClassesRequired int64 `json:"classesRequired" db:"classes_required"`  //`classes_required` int(11) NOT NULL DEFAULT '65535',
	Probability     int64 `json:"probability" db:"probability"`           //`probability` int(3) NOT NULL DEFAULT '100',
}
