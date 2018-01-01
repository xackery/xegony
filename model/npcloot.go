package model

import ()

//NpcLoot is a cache table used to speed up lookup of items for NPCs
type NpcLoot struct {
	NpcID   int64  `json:"npcID" db:"npc_id"`
	ItemID  int64  `json:"itemID" db:"item_id"`
	NpcName string `json:"npcName" db:"npc_name"`
	*Item
	*Npc
}

//NpcCleanName returns an NPC as a clean name
func (c *NpcLoot) NpcCleanName() string {
	return CleanName(c.NpcName)
}
