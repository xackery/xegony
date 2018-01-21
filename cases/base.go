package cases

import (
	"fmt"

	"github.com/pkg/errors"
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
func (c *BaseRepository) Get(base *model.Base, user *model.User) (err error) {
	err = c.stor.GetBase(base)
	if err != nil {
		err = errors.Wrap(err, "failed to get base")
		return
	}
	err = c.prepare(base, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare base")
		return
	}
	return
}

//Create handles logic
func (c *BaseRepository) Create(base *model.Base, user *model.User) (err error) {
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
	err = c.prepare(base, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare base")
		return
	}
	return
}

//Edit handles logic
func (c *BaseRepository) Edit(base *model.Base, user *model.User) (err error) {
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

	err = c.stor.EditBase(base)
	if err != nil {
		return
	}
	err = c.prepare(base, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare base")
		return
	}
	return
}

//Delete handles logic
func (c *BaseRepository) Delete(base *model.Base, user *model.User) (err error) {
	err = c.stor.DeleteBase(base)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *BaseRepository) List(user *model.User) (bases []*model.Base, err error) {
	bases, err = c.stor.ListBase()
	if err != nil {
		return
	}
	for _, base := range bases {
		err = c.prepare(base, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare base")
			return
		}
	}
	return
}

func (c *BaseRepository) prepare(base *model.Base, user *model.User) (err error) {

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
