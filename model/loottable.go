package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//LootTable is the parent of loottableentry
type LootTable struct {
	Id      int64  `json:"id" db:"id"`           //`id` int(11) unsigned NOT NULL AUTO_INCREMENT,
	Name    string `json:"name" db:"name"`       //`name` varchar(255) NOT NULL DEFAULT '',
	Mincash int64  `json:"mincash" db:"mincash"` //`mincash` int(11) unsigned NOT NULL DEFAULT '0',
	Maxcash int64  `json:"maxcash" db:"maxcash"` //`maxcash` int(11) unsigned NOT NULL DEFAULT '0',
	Avgcoin int64  `json:"avgcoin" db:"avgcoin"` //`avgcoin` int(10) unsigned NOT NULL DEFAULT '0',
	Done    int64  `json:"done" db:"done"`       //`done` tinyint(3) NOT NULL DEFAULT '0',
	Entries []*LootTableEntry
	Npcs    []*Npc
}

func (c *LootTable) MinCashName() string {
	return CashName(c.Mincash)
}
func (c *LootTable) MaxCashName() string {
	return CashName(c.Maxcash)
}

func (c *LootTable) AvgCoinName() string {
	return CashName(c.Avgcoin)
}

func (c *LootTable) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *LootTable) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "lootTableID":
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
