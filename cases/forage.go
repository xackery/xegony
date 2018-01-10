package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ForageRepository handles ForageRepository cases and is a gateway to storage
type ForageRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *ForageRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *ForageRepository) Get(forageID int64) (forage *model.Forage, err error) {
	if forageID == 0 {
		err = fmt.Errorf("Invalid Forage ID")
		return
	}
	forage, err = c.stor.GetForage(forageID)
	return
}

//Create handles logic
func (c *ForageRepository) Create(forage *model.Forage) (err error) {
	if forage == nil {
		err = fmt.Errorf("Empty forage")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	forage.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(forage))
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
	err = c.stor.CreateForage(forage)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *ForageRepository) Edit(forageID int64, forage *model.Forage) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(forage))
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

	err = c.stor.EditForage(forageID, forage)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *ForageRepository) Delete(forageID int64) (err error) {
	err = c.stor.DeleteForage(forageID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ForageRepository) List(pageSize int64, pageNumber int64) (forages []*model.Forage, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	forages, err = c.stor.ListForage(pageSize, pageNumber)
	if err != nil {
		return
	}
	return
}

//ListCount handles logic
func (c *ForageRepository) ListCount() (count int64, err error) {

	count, err = c.stor.ListForageCount()
	if err != nil {
		return
	}
	return
}

//GetByZone handles logic
func (c *ForageRepository) GetByZone(zoneID int64) (forages []*model.Forage, err error) {
	forages, err = c.stor.ListForageByZone(zoneID)
	if err != nil {
		return
	}
	return
}

//GetByItem handles logic
func (c *ForageRepository) GetByItem(itemID int64) (forages []*model.Forage, err error) {
	forages, err = c.stor.ListForageByItem(itemID)
	if err != nil {
		return
	}
	return
}

func (c *ForageRepository) prepare(forage *model.Forage) (err error) {

	return
}

func (c *ForageRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *ForageRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
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
