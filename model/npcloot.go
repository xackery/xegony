package model

import ()

//NpcLoot is a cache table used to speed up lookup of items for NPCs
type NpcLoot struct {
	NpcID  int64 `json:"npcID" db:"npc_id"`
	ItemID int64 `json:"itemID" db:"item_id"`
	Item   *Item `json:"item"`
	Npc    *Npc  `json:"npc"`
}
