package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Base struct {
	Level   int64   `json:"level" db:"level"`      //`level` int(10) unsigned NOT NULL,
	Class   int64   `json:"class" db:"class"`      //`class` int(10) unsigned NOT NULL,
	Hp      float64 `json:"hp" db:"hp"`            //`hp` double NOT NULL,
	Mana    float64 `json:"mana" db:"mana"`        //`mana` double NOT NULL,
	End     float64 `json:"end" db:"end"`          //`end` double NOT NULL,
	Unk1    float64 `json:"unk1" db:"unk1"`        //`unk1` double NOT NULL,
	Unk2    float64 `json:"unk2" db:"unk2"`        //`unk2` double NOT NULL,
	HpFac   float64 `json:"hpFac" db:"hp_fac"`     //`hp_fac` double NOT NULL,
	ManaFac float64 `json:"manaFac" db:"mana_fac"` //`mana_fac` double NOT NULL,
	EndFac  float64 `json:"endFac" db:"end_fac"`   //`end_fac` double NOT NULL,
}

func (c *Base) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *Base) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "itemId":
		prop.Type = "integer"
		prop.Minimum = 1
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
