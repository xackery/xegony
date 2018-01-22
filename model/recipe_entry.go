package model

import ()

//RecipeEntry group together loot drops for npc drops
// swagger:response
type RecipeEntry struct {
	Item   *Item   `json:"item"`
	Recipe *Recipe `json:"recipe"`

	ID             int64 `json:"id" db:"id"`                         //`id` int(11) NOT NULL AUTO_INCREMENT,
	RecipeID       int64 `json:"recipeId" db:"recipe_id"`            //`recipe_id` int(11) NOT NULL DEFAULT '0',
	ItemID         int64 `json:"itemId" db:"item_id"`                //`item_id` int(11) NOT NULL DEFAULT '0',
	Successcount   int64 `json:"successcount" db:"successcount"`     //`successcount` tinyint(2) NOT NULL DEFAULT '0',
	Failcount      int64 `json:"failcount" db:"failcount"`           //`failcount` tinyint(2) NOT NULL DEFAULT '0',
	Componentcount int64 `json:"componentcount" db:"componentcount"` //`componentcount` tinyint(2) NOT NULL DEFAULT '1',
	Salvagecount   int64 `json:"salvagecount" db:"salvagecount"`     //`salvagecount` tinyint(2) NOT NULL DEFAULT '0',
	Iscontainer    int64 `json:"iscontainer" db:"iscontainer"`       //`iscontainer` tinyint(1) NOT NULL DEFAULT '0',
}
