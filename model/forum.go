package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Forum struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	OwnerId     int64  `json:"ownerId" db:"owner_id"`
	Description string `json:"description"`
	Icon        string `json:"icon" db:"icon"`
}

func (c *Forum) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *Forum) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "ownerId":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z' ]*$"
	case "description":
		prop.Type = "string"
		prop.MaxLength = 128
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
