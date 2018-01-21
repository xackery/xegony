package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//ErrorRepository handles ErrorRepository cases and is a gateway to storage
type ErrorRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *ErrorRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *ErrorRepository) Get(errorStruct *model.Error, user *model.User) (err error) {

	err = c.stor.GetError(errorStruct)
	if err != nil {
		err = errors.Wrap(err, "failed to get error")
		return
	}
	err = c.prepare(errorStruct, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare error")
		return
	}
	return
}

//Search handles logic
func (c *ErrorRepository) SearchByMessage(errorStruct *model.Error, user *model.User) (errorStructs []*model.Error, err error) {
	errorStructs, err = c.stor.SearchErrorByMessage(errorStruct)
	if err != nil {
		return
	}
	for _, errorStruct := range errorStructs {
		err = c.prepare(errorStruct, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare error")
			return
		}
	}
	return
}

//Create handles logic
func (c *ErrorRepository) Create(errorStruct *model.Error, user *model.User) (err error) {
	if errorStruct == nil {
		err = fmt.Errorf("Empty error")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	errorStruct.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(errorStruct))
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
	err = c.stor.CreateError(errorStruct)
	if err != nil {
		return
	}
	err = c.prepare(errorStruct, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare error")
		return
	}
	return
}

//Edit handles logic
func (c *ErrorRepository) Edit(errorStruct *model.Error, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(errorStruct))
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

	err = c.stor.EditError(errorStruct)
	if err != nil {
		return
	}
	err = c.prepare(errorStruct, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare error")
		return
	}
	return
}

//Delete handles logic
func (c *ErrorRepository) Delete(errorStruct *model.Error, user *model.User) (err error) {
	err = c.stor.DeleteError(errorStruct)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *ErrorRepository) List(pageSize int64, pageNumber int64, user *model.User) (errorStructs []*model.Error, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	errorStructs, err = c.stor.ListError(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, errorStruct := range errorStructs {
		err = c.prepare(errorStruct, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare error")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *ErrorRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListErrorCount()
	if err != nil {
		return
	}
	return
}

func (c *ErrorRepository) prepare(errorStruct *model.Error, user *model.User) (err error) {

	return
}

func (c *ErrorRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *ErrorRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
