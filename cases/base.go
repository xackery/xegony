package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//BaseRepository handles BaseRepository cases and is a gateway to storage
type BaseRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *BaseRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *BaseRepository) Get(class int64, level int64) (base *model.Base, err error) {
	base, err = c.stor.GetBase(class, level)
	return
}

//Create handles logic
func (c *BaseRepository) Create(base *model.Base) (err error) {
	if base == nil {
		err = fmt.Errorf("Empty base")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(base))
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
	err = c.stor.CreateBase(base)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *BaseRepository) Edit(class int64, level int64, base *model.Base) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(base))
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

	err = c.stor.EditBase(class, level, base)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *BaseRepository) Delete(class int64, level int64) (err error) {
	err = c.stor.DeleteBase(class, level)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *BaseRepository) List() (bases []*model.Base, err error) {
	bases, err = c.stor.ListBase()
	if err != nil {
		return
	}
	return
}

func (c *BaseRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *BaseRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "itemID":
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
