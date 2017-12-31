package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type LootDropEntry struct {
	LootdropId     int64   `json:"lootdropId" db:"lootdrop_id"`         //`lootdrop_id` int(11) unsigned NOT NULL DEFAULT '0',
	ItemId         int64   `json:"itemID" db:"item_id"`                 //`item_id` int(11) NOT NULL DEFAULT '0',
	ItemCharges    int64   `json:"itemCharges" db:"item_charges"`       //`item_charges` smallint(2) unsigned NOT NULL DEFAULT '1',
	EquipItem      int64   `json:"equipItem" db:"equip_item"`           //`equip_item` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Chance         float64 `json:"chance" db:"chance"`                  //`chance` float NOT NULL DEFAULT '1',
	DisabledChance float64 `json:"disabledChance" db:"disabled_chance"` //`disabled_chance` float NOT NULL DEFAULT '0',
	Minlevel       int64   `json:"minlevel" db:"minlevel"`              //`minlevel` tinyint(3) NOT NULL DEFAULT '0',
	Maxlevel       int64   `json:"maxlevel" db:"maxlevel"`              //`maxlevel` tinyint(3) NOT NULL DEFAULT '127',
	Multiplier     int64   `json:"multiplier" db:"multiplier"`          //`multiplier` tinyint(2) unsigned NOT NULL DEFAULT '1',
}

func (c *LootDropEntry) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]Schema)
	var field string
	var prop Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

func (c *LootDropEntry) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "lootDropID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
