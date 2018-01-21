package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//HackerRepository handles HackerRepository cases and is a gateway to storage
type HackerRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *HackerRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *HackerRepository) Get(hacker *model.Hacker, user *model.User) (err error) {

	err = c.stor.GetHacker(hacker)
	if err != nil {
		err = errors.Wrap(err, "failed to get hacker")
		return
	}

	err = c.prepare(hacker, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare hacker")
		return
	}
	return
}

//Search handles logic
func (c *HackerRepository) SearchByMessage(hacker *model.Hacker, user *model.User) (hackers []*model.Hacker, err error) {
	hackers, err = c.stor.SearchHackerByMessage(hacker)
	if err != nil {
		return
	}
	for _, hacker := range hackers {
		err = c.prepare(hacker, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare hacker")
			return
		}
	}
	return
}

//Create handles logic
func (c *HackerRepository) Create(hacker *model.Hacker, user *model.User) (err error) {
	if hacker == nil {
		err = fmt.Errorf("Empty hacker")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	hacker.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(hacker))
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
	err = c.stor.CreateHacker(hacker)
	if err != nil {
		return
	}
	err = c.prepare(hacker, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare hacker")
		return
	}
	return
}

//Edit handles logic
func (c *HackerRepository) Edit(hacker *model.Hacker, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(hacker))
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

	err = c.stor.EditHacker(hacker)
	if err != nil {
		return
	}
	err = c.prepare(hacker, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare hacker")
		return
	}
	return
}

//Delete handles logic
func (c *HackerRepository) Delete(hacker *model.Hacker, user *model.User) (err error) {
	err = c.stor.DeleteHacker(hacker)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *HackerRepository) List(pageSize int64, pageNumber int64, user *model.User) (hackers []*model.Hacker, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	hackers, err = c.stor.ListHacker(pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, hacker := range hackers {
		err = c.prepare(hacker, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare hacker")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *HackerRepository) ListCount(user *model.User) (count int64, err error) {

	count, err = c.stor.ListHackerCount()
	if err != nil {
		return
	}
	return
}

func (c *HackerRepository) prepare(hacker *model.Hacker, user *model.User) (err error) {

	return
}
func (c *HackerRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *HackerRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
