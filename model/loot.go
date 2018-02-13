package model

//Loots is an array of loot
// swagger:model
type Loots []*Loot

//Loot is the parent of lootentry
// swagger:model
type Loot struct {
	ID      int64  `json:"ID" db:"id"`           //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	Name    string `json:"name" db:"name"`       //`name` varchar(255) NOT NULL DEFAULT '',
	Mincash int64  `json:"mincash" db:"mincash"` //`mincash` int(11) unsigned NOT NULL DEFAULT '0',
	Maxcash int64  `json:"maxcash" db:"maxcash"` //`maxcash` int(11) unsigned NOT NULL DEFAULT '0',
	Avgcoin int64  `json:"avgcoin" db:"avgcoin"` //`avgcoin` int(10) unsigned NOT NULL DEFAULT '0',
	Done    int64  `json:"done" db:"done"`       //`done` tinyint(3) NOT NULL DEFAULT '0',
	Entries []*LootEntry
	Npcs    []*Npc
}

//MinCashName returns human readable form of mincash
func (c *Loot) MinCashName() string {
	return CashName(c.Mincash)
}

//MaxCashName returns human readable max cash
func (c *Loot) MaxCashName() string {
	return CashName(c.Maxcash)
}

//AvgCoinName returns human readable avg coin
func (c *Loot) AvgCoinName() string {
	return CashName(c.Avgcoin)
}
