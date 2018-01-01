package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ZoneLevelRepository handles ZoneLevelRepository cases and is a gateway to storage
type ZoneLevelRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *ZoneLevelRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *ZoneLevelRepository) Get(zoneID int64) (zoneLevel *model.ZoneLevel, err error) {
	zoneLevel, err = c.stor.GetZoneLevel(zoneID)
	return
}

//Truncate handles logic
func (c *ZoneLevelRepository) Truncate() (err error) {
	err = c.stor.TruncateZoneLevel()
	return
}

//Create handles logic
func (c *ZoneLevelRepository) Create(zoneLevel *model.ZoneLevel) (err error) {
	if zoneLevel == nil {
		err = fmt.Errorf("Empty zoneLevel")
		return
	}
	schema, err := c.newSchema([]string{}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(zoneLevel))
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
	err = c.stor.CreateZoneLevel(zoneLevel)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *ZoneLevelRepository) Edit(zoneID int64, zoneLevel *model.ZoneLevel) (err error) {
	schema, err := c.newSchema(nil, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(zoneLevel))
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

	err = c.stor.EditZoneLevel(zoneID, zoneLevel)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *ZoneLevelRepository) Delete(zoneID int64) (err error) {
	err = c.stor.DeleteZoneLevel(zoneID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ZoneLevelRepository) List() (zoneLevels []*model.ZoneLevel, err error) {
	zoneLevels, err = c.stor.ListZoneLevel()
	if err != nil {
		return
	}
	return
}

func (c *ZoneLevelRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *ZoneLevelRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "accountID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "zoneID":
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
