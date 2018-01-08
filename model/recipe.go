package model

import (
	"database/sql"
	"fmt"
	"html/template"
)

//Recipe represents tradeskill recipes inside everquest
type Recipe struct {
	Entrys []*RecipeEntry

	ID               int64          `json:"id" db:"id"`                              //`id` int(11) NOT NULL AUTO_INCREMENT,
	Name             string         `json:"name" db:"name"`                          //`name` varchar(64) NOT NULL DEFAULT '',
	Tradeskill       int64          `json:"tradeskill" db:"tradeskill"`              //`tradeskill` smallint(6) NOT NULL DEFAULT '0',
	Skillneeded      int64          `json:"skillneeded" db:"skillneeded"`            //`skillneeded` smallint(6) NOT NULL DEFAULT '0',
	Trivial          int64          `json:"trivial" db:"trivial"`                    //`trivial` smallint(6) NOT NULL DEFAULT '0',
	Nofail           int64          `json:"nofail" db:"nofail"`                      //`nofail` tinyint(1) NOT NULL DEFAULT '0',
	ReplaceContainer int64          `json:"replaceContainer" db:"replace_container"` //`replace_container` tinyint(1) NOT NULL DEFAULT '0',
	Notes            sql.NullString `json:"notes" db:"notes"`                        //`notes` tinytext,
	MustLearn        int64          `json:"mustLearn" db:"must_learn"`               //`must_learn` tinyint(4) NOT NULL DEFAULT '0',
	Quest            int64          `json:"quest" db:"quest"`                        //`quest` tinyint(1) NOT NULL DEFAULT '0',
	Enabled          int64          `json:"enabled" db:"enabled"`                    //`enabled` tinyint(1) NOT NULL DEFAULT '1',
}

//ProfitMarginName takes difference of a tradeskill cost to item reward
func (c *Recipe) ProfitMarginName() string {
	//Get reagent cost
	cost := int64(0)
	for _, entry := range c.Entrys {
		if entry.Componentcount > 0 && entry.Successcount == 0 && entry.Failcount == 0 && entry.Item != nil {
			cost += entry.Item.Price
		}
	}
	//Get reward price
	price := int64(0)
	for _, entry := range c.Entrys {
		if entry.Successcount > 0 && entry.Item != nil {
			price = entry.Item.Price
		}
	}
	//Take difference
	price = price - cost
	if price == 0 {
		return "None"
	}
	if price < 0 {
		return fmt.Sprintf("Loss %s", CashName(-price))
	}
	return CashName(price)
}

//RewardItem returns the primary reward item for a recipe
func (c *Recipe) RewardItem() *Item {
	for _, entry := range c.Entrys {
		if entry.Successcount > 0 && entry.Item != nil {
			return entry.Item
		}
	}
	return nil
}

//SkillName returns a skill in human readable format
func (c *Recipe) SkillName() string {
	return SkillName(c.Tradeskill)
}

//ReagentIconList returns icons as <span> elements
func (c *Recipe) ReagentIconList() template.HTML {
	icons := ""
	for _, entry := range c.Entrys {
		if entry.Componentcount > 0 && entry.Successcount == 0 && entry.Failcount == 0 && entry.Item != nil {
			icons += fmt.Sprintf(`<span item="%d" class="item icon-%d-sm"></span>`, entry.Item.ID, entry.Item.Icon)
		}
	}
	return template.HTML(icons)
}

//ReagentPriceList returns a total human readable price for reagents
func (c *Recipe) ReagentPriceList() string {
	price := int64(0)
	for _, entry := range c.Entrys {
		if entry.Componentcount > 0 && entry.Successcount == 0 && entry.Failcount == 0 && entry.Item != nil {
			price += entry.Item.Price
		}
	}
	return CashName(price)
}

//ToolIconList returns icons as <span> elements
func (c *Recipe) ToolIconList() template.HTML {
	icons := ""
	for _, entry := range c.Entrys {
		if entry.Iscontainer > 0 && entry.Item != nil {
			icons += fmt.Sprintf(`<span item="%d" class="item icon-%d-sm"></span>`, entry.Item.ID, entry.Item.Icon)
		}
	}
	return template.HTML(icons)
}
