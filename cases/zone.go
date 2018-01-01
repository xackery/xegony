package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ZoneRepository holds cases
type ZoneRepository struct {
	stor storage.Storage
}

//Initialize handler
func (c *ZoneRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handler
func (c *ZoneRepository) Get(zoneID int64) (zone *model.Zone, err error) {
	if zoneID == 0 {
		err = fmt.Errorf("Invalid Zone ID")
		return
	}
	zone, err = c.stor.GetZone(zoneID)
	return
}

//Create handler
func (c *ZoneRepository) Create(zone *model.Zone) (err error) {
	if zone == nil {
		err = fmt.Errorf("Empty zone")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
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
	err = c.stor.CreateZone(zone)
	if err != nil {
		return
	}
	return
}

//Edit handler
func (c *ZoneRepository) Edit(zoneID int64, zone *model.Zone) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
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

	err = c.stor.EditZone(zoneID, zone)
	if err != nil {
		return
	}
	return
}

//Delete handler
func (c *ZoneRepository) Delete(zoneID int64) (err error) {
	err = c.stor.DeleteZone(zoneID)
	if err != nil {
		return
	}
	return
}

//List handler
func (c *ZoneRepository) List() (zones []*model.Zone, err error) {
	zones, err = c.stor.ListZone()
	if err != nil {
		return
	}
	return
}

//ListByHotzone handler
func (c *ZoneRepository) ListByHotzone() (zones []*model.Zone, err error) {
	zones, err = c.stor.ListZoneByHotzone()
	if err != nil {
		return
	}
	return
}

//newSchema handler
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

//getSchemaProperty handler
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
