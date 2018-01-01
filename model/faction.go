package model

import ()

//Faction represents everquest factions
type Faction struct {
	Id   int64  `json:"id"`
	Base int64  `json:"base"`
	Name string `json:"name"`
}

func (c *Faction) CleanName() string {
	return CleanName(c.Name)
}
