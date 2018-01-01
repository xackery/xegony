package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//SpawnEntry group together loot drops for npc drops
type SpawnEntry struct {
	SpawngroupID int64 `json:"spawngroupID" db:"spawngroupID"` //`spawngroupID` int(11) NOT NULL DEFAULT '0',
	NpcID        int64 `json:"npcID" db:"npcID"`               //`npcID` int(11) NOT NULL DEFAULT '0',
	Chance       int64 `json:"chance" db:"chance"`             //`chance` smallint(4) NOT NULL DEFAULT '0',
}

func (c *SpawnEntry) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *SpawnEntry) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "spawnID":
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
