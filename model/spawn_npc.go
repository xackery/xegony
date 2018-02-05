package model

// SpawnNpcs is an array of SpawnNpc
// swagger:model
type SpawnNpcs []*SpawnNpc

// SpawnNpc represents the spawnentry table. This is a pivot
// swagger:model
type SpawnNpc struct {
	Npc *Npc `json:"npc,omitempty"`

	SpawnID int64 `json:"spawnID,omitempty" db:"spawngroupID"` //`spawngroupID` int(11) NOT NULL DEFAULT '0',
	NpcID   int64 `json:"npcID,omitempty" db:"npcID"`          //`npcID` int(11) NOT NULL DEFAULT '0',
	Chance  int64 `json:"chance,omitempty" db:"chance"`        //`chance` smallint(4) NOT NULL DEFAULT '0',
}
