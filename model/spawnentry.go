package model

import ()

//SpawnEntry group together loot drops for npc drops
// swagger:model
type SpawnEntry struct {
	SpawngroupID int64 `json:"spawngroupID" db:"spawngroupID"` //`spawngroupID` int(11) NOT NULL DEFAULT '0',
	NpcID        int64 `json:"npcID" db:"npcID"`               //`npcID` int(11) NOT NULL DEFAULT '0',
	Chance       int64 `json:"chance" db:"chance"`             //`chance` smallint(4) NOT NULL DEFAULT '0',
}
