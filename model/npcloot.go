package model

import ()

//NpcLoot is a cache table used to speed up lookup of items for NPCs
type NpcLoot struct {
	NpcId   int64  `json:"npcID" db:"npc_id"`
	ItemId  int64  `json:"itemID" db:"item_id"`
	NpcName string `json:"npcName" db:"npc_name"`
	*Item
	*Npc
}

func (c *NpcLoot) NpcCleanName() string {
	return CleanName(c.NpcName)
}
