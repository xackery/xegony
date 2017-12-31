package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type ZoneRepository struct {
	stor storage.Storage
}

func (g *ZoneRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *ZoneRepository) Get(zoneID int64) (zone *model.Zone, err error) {
	if zoneID == 0 {
		err = fmt.Errorf("Invalid Zone ID")
		return
	}
	zone, err = g.stor.GetZone(zoneID)
	return
}

func (g *ZoneRepository) Create(zone *model.Zone) (err error) {
	if zone == nil {
		err = fmt.Errorf("Empty zone")
		return
	}
	schema, err := g.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}
	if zone.ZoneIDNumber < 1 {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		vErr.Reasons["accountID"] = "Account ID must be greater than 0"
		err = vErr
		return
	}
	zone.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(zone))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}
	err = g.stor.CreateZone(zone)
	if err != nil {
		return
	}
	return
}

func (g *ZoneRepository) Edit(zoneID int64, zone *model.Zone) (err error) {
	schema, err := g.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(zone))
	if err != nil {
		return
	}
	if !result.Valid() {
		vErr := &model.ErrValidation{
			Message: "invalid",
		}
		vErr.Reasons = map[string]string{}
		for _, res := range result.Errors() {
			vErr.Reasons[res.Field()] = res.Description()
		}
		err = vErr
		return
	}

	err = g.stor.EditZone(zoneID, zone)
	if err != nil {
		return
	}
	return
}

func (g *ZoneRepository) Delete(zoneID int64) (err error) {
	err = g.stor.DeleteZone(zoneID)
	if err != nil {
		return
	}
	return
}

func (g *ZoneRepository) List() (zones []*model.Zone, err error) {
	zones, err = g.stor.ListZone()
	if err != nil {
		return
	}
	return
}

func (g *ZoneRepository) ListByHotzone() (zones []*model.Zone, err error) {
	zones, err = g.stor.ListZoneByHotzone()
	if err != nil {
		return
	}
	return
}

func (c *ZoneRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
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

func (c *ZoneRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
