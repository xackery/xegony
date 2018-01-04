package model

import ()

//Fishing represents items inside everquest
type Fishing struct {
	Item *Item
	Zone *Zone
	Npc  *Npc

	ID         int64 `json:"id" db:"id"`                  //`id` int(11) NOT NULL AUTO_INCREMENT,
	ZoneID     int64 `json:"zoneid" db:"zoneid"`          //`zoneid` int(4) NOT NULL DEFAULT '0',
	ItemID     int64 `json:"Itemid" db:"Itemid"`          //`Itemid` int(11) NOT NULL DEFAULT '0',
	SkillLevel int64 `json:"skillLevel" db:"skill_level"` //`skill_level` smallint(6) NOT NULL DEFAULT '0',
	Chance     int64 `json:"chance" db:"chance"`          //`chance` smallint(6) NOT NULL DEFAULT '0',
	NpcID      int64 `json:"npcId" db:"npc_id"`           //`npc_id` int(11) NOT NULL DEFAULT '0',
	NpcChance  int64 `json:"npcChance" db:"npc_chance"`   //`npc_chance` int(11) NOT NULL DEFAULT '0',
}
