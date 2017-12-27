package model

import (
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type Npc struct {
	Id          int64          `json:"id"`
	AccountId   int64          `json:"accountId" db:"account_id"`
	Name        string         `json:"name"`
	LastName    sql.NullString `json:"lastName" db:"lastname"`
	Hp          int64          `json:"hp"`
	Level       int64          `json:"level"`
	Class       int64          `json:"class" db:"class"`
	LootTableId int64          `json:"lootTableId" db:"loottable_id"`
}

func (c *Npc) ZoneId() int64 {
	if c.Id > 1000 {
		return c.Id / 1000
	}
	return 0
}

func (c *Npc) CleanName() string {
	var re = regexp.MustCompile(`[^0-9A-Za-z_]+`)
	cleanName := strings.Replace(c.Name, " ", "_", -1)
	cleanName = strings.Replace(cleanName, "#", "", -1)
	cleanName = strings.TrimSpace(re.ReplaceAllString(cleanName, ""))
	cleanName = strings.Replace(cleanName, "_", " ", -1)
	return cleanName
}

func (c *Npc) ClassName() string {
	switch c.Class {
	case 1:
		return "Warrior"
	case 2:
		return "Cleric"
	case 3:
		return "Paladin"
	case 4:
		return "Ranger"
	case 5:
		return "Shadowknight"
	case 6:
		return "Druid"
	case 7:
		return "Monk"
	case 8:
		return "Bard"
	case 9:
		return "Rogue"
	case 10:
		return "Shaman"
	case 11:
		return "Necromancer"
	case 12:
		return "Wizard"
	case 13:
		return "Magician"
	case 14:
		return "Enchanter"
	case 15:
		return "Beastlord"
	case 16:
		return "Berserker"
	}
	return "Unknown"
}

func (c *Npc) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *Npc) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "accountId":
		prop.Type = "integer"
		prop.Minimum = 1
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneId":
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
