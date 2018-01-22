package model

import (
	"fmt"
	"html/template"
)

//Merchant represents tradeskill merchants inside everquest
// swagger:response
type Merchant struct {
	Entrys []*MerchantEntry
	Npcs   []*Npc
	ID     int64 `json:"merchantID" db:"merchantid"` //`merchantid` int(11) NOT NULL DEFAULT '0',
}

//ItemIconList returns icons as <span> elements
// swagger:response
func (c *Merchant) ItemIconList() template.HTML {
	icons := ""

	itemCount := 0
	for _, entry := range c.Entrys {
		item := entry.Item
		icons += fmt.Sprintf(`<span item="%d" class="item icon-%d-sm"></span>`, item.ID, item.Icon)
		itemCount++
		if itemCount > 10 {
			break
		}
	}
	return template.HTML(icons)
}
