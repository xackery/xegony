package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Zone struct {
	Id           int64  `json:"id"`
	ShortName    string `json:"shortName" db:"short_name"`
	LongName     string `json:"longName" db:"long_name"`
	ZoneIdNumber int64  `json:"zoneIdNumber"`
	MinStatus    int64  `json:"minStatus" db:"min_status"`
}

func (c *Zone) ExpansionId() int64 {
	switch c.ExpansionBit() {
	case 0:
		return 0
	case 1:
		return 1
	case 2:
		return 2
	case 4:
		return 3
	case 8:
		return 4
	case 16:
		return 5
	case 32:
		return 6
	case 64:
		return 7
	case 127:
		return 8
	case 256:
		return 9
	case 512:
		return 10
	case 1024:
		return 11
	case 2048:
		return 12
	case 4096:
		return 13
	case 8192:
		return 14
	default:
		return -1
	}
	return -1
}

func (c *Zone) ExpansionBit() int64 {
	switch c.ShortName {
	case "airplane", "akanon":
		return 0 // - classic
	case "overthere":
		return 1 // - ruins of kunark
	case "thurgadina":
		return 2 // - scars of velious
	case "acrylia", "akheva":
		return 4 // - shadows of luclin
	case "poknowledge":
		return 8 // - planes of power
	case "nadox":
		return 16 // - legacy of ykesha
	case "mira":
		return 32 // - lost dungeons of norrath
	case "wallofslaughter":
		return 64 // - gates of discord
	case "abysmal", "anguish":
		return 128 // - omens of war
	case "stillmoona":
		return 256 // - dragons of norrath
	case "asd":
		return 512 // - depths of darkhallow
	case "asdf":
		return 1024 // - prophecy of ro
	case "asdfg":
		return 2048 // - serpent's spine
	case "gsdg":
		return 4096 // - the burried sea
	case "hdshd":
		return 8192 // - secrets of faydwer
	default:
		return -1
	}
	return -1
}

func (c *Zone) ExpansionName() string {
	switch c.ExpansionId() {
	case -1:
		return "Unknown"
	case 0:
		return "Classic"
	case 1:
		return "Ruins of Kunark"
	case 2:
		return "Scars of Velious"
	case 3:
		return "Shadows of Luclin"
	case 4:
		return "Planes of Power"
	case 5:
		return "Legacy of Ykesha"
	case 6:
		return "Lost Dungeons of Norrath"
	case 7:
		return "Gates of Discord"
	case 8:
		return "Omens of War"
	case 9:
		return "Dragons of Norrath"
	case 10:
		return "Depths of Darkhallow"
	case 11:
		return "Prophecy of Ro"
	case 12:
		return "Serpent's Spine"
	case 13:
		return "The Buried Sea"
	case 14:
		return "Secrets of Faydwer"
	}
	return "Unknown"
}

func (c *Zone) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *Zone) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "shortName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
