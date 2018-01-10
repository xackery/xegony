package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//AaRepository handles AaRepository cases and is a gateway to storage
type AaRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *AaRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *AaRepository) Get(aaID int64) (aa *model.Aa, err error) {
	if aaID == 0 {
		err = fmt.Errorf("Invalid Aa ID")
		return
	}
	aa, err = c.stor.GetAa(aaID)
	return
}

//Create handles logic
func (c *AaRepository) Create(aa *model.Aa) (err error) {
	if aa == nil {
		err = fmt.Errorf("Empty aa")
		return
	}
	schema, err := c.newSchema([]string{"shortName"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(aa))
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
	err = c.stor.CreateAa(aa)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *AaRepository) Edit(aaID int64, aa *model.Aa) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(aa))
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

	err = c.stor.EditAa(aaID, aa)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *AaRepository) Delete(aaID int64) (err error) {
	err = c.stor.DeleteAa(aaID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *AaRepository) List() (aas []*model.Aa, err error) {
	aas, err = c.stor.ListAa()
	if err != nil {
		return
	}
	return
}

func (c *AaRepository) prepare(aa *model.Aa) (err error) {

	return
}

func (c *AaRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *AaRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
