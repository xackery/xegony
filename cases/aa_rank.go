package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//AaRankRepository handles AaRankRepository cases and is a gateway to storage
type AaRankRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *AaRankRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *AaRankRepository) Get(aaRank *model.AaRank, user *model.User) (err error) {
	err = c.stor.GetAaRank(aaRank)
	if err != nil {
		err = errors.Wrap(err, "failed to get aa rank")
		return
	}
	if err != nil {
		err = errors.Wrap(err, "failed to prepare aa rank")
		return
	}
	return
}

//Create handles logic
func (c *AaRankRepository) Create(aaEntry *model.AaRank, user *model.User) (err error) {
	if aaEntry == nil {
		err = fmt.Errorf("Empty AaRank")
		return
	}

	schema, err := c.newSchema(nil, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(aaEntry))
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
	err = c.stor.CreateAaRank(aaEntry)
	if err != nil {
		return
	}
	if err != nil {
		err = errors.Wrap(err, "failed to prepare aa rank")
		return
	}
	return
}

//Edit handles logic
func (c *AaRankRepository) Edit(aaRank *model.AaRank, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(aaRank))
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

	err = c.stor.EditAaRank(aaRank)
	if err != nil {
		return
	}
	if err != nil {
		err = errors.Wrap(err, "failed to prepare aa rank")
		return
	}
	return
}

//Delete handles logic
func (c *AaRankRepository) Delete(aaRank *model.AaRank, user *model.User) (err error) {
	err = c.stor.DeleteAaRank(aaRank)
	if err != nil {
		return
	}
	if err != nil {
		err = errors.Wrap(err, "failed to prepare aa rank")
		return
	}
	return
}

//List handles logic
func (c *AaRankRepository) List(user *model.User) (aaRanks []*model.AaRank, err error) {
	aaRanks, err = c.stor.ListAaRank()
	if err != nil {
		return
	}
	for _, aaRank := range aaRanks {
		err = c.prepare(aaRank, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare aa rank")
			return
		}
	}
	return
}

func (c *AaRankRepository) prepare(aaRank *model.AaRank, user *model.User) (err error) {

	return
}

func (c *AaRankRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *AaRankRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
