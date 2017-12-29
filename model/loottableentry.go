package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type LootTableEntry struct {
	LoottableId int64   `json:"loottableId" db:"loottable_id"` //`loottable_id` int(11) unsigned NOT NULL DEFAULT '0',
	LootdropId  int64   `json:"lootdropId" db:"lootdrop_id"`   //`lootdrop_id` int(11) unsigned NOT NULL DEFAULT '0',
	Multiplier  int64   `json:"multiplier" db:"multiplier"`    //`multiplier` tinyint(2) unsigned NOT NULL DEFAULT '1',
	Droplimit   int64   `json:"droplimit" db:"droplimit"`      //`droplimit` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Mindrop     int64   `json:"mindrop" db:"mindrop"`          //`mindrop` tinyint(2) unsigned NOT NULL DEFAULT '0',
	Probability float64 `json:"probability" db:"probability"`  //`probability` float NOT NULL DEFAULT '100',
}

func (c *LootTableEntry) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *LootTableEntry) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "lootTableId":
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
