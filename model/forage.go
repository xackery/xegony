package model

import ()

//Forage represents items inside everquest
type Forage struct {
	Item *Item
	Zone *Zone

	ID     int64 `json:"id" db:"id"`         //`id` int(11) NOT NULL AUTO_INCREMENT,
	ZoneID int64 `json:"zoneid" db:"zoneid"` //`zoneid` int(4) NOT NULL DEFAULT '0',
	ItemID int64 `json:"Itemid" db:"Itemid"` //`Itemid` int(11) NOT NULL DEFAULT '0',
	Level  int64 `json:"level" db:"level"`   //`level` smallint(6) NOT NULL DEFAULT '0',
	Chance int64 `json:"chance" db:"chance"` //`chance` smallint(6) NOT NULL DEFAULT '0',
}
