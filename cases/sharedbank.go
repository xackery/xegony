package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//SharedBankRepository handles SharedBankRepository cases and is a gateway to storage
type SharedBankRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *SharedBankRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *SharedBankRepository) Get(sharedBank *model.SharedBank, user *model.User) (err error) {
	err = c.stor.GetSharedBank(sharedBank)
	if err != nil {
		err = errors.Wrap(err, "failed to get shared bank")
		return
	}
	err = c.prepare(sharedBank, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare shared bank")
		return
	}
	return
}

//Create handles logic
func (c *SharedBankRepository) Create(sharedBank *model.SharedBank, user *model.User) (err error) {
	if sharedBank == nil {
		err = fmt.Errorf("Empty sharedBank")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(sharedBank))
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
	err = c.stor.CreateSharedBank(sharedBank)
	if err != nil {
		return
	}
	err = c.prepare(sharedBank, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare shared bank")
		return
	}
	return
}

//Edit handles logic
func (c *SharedBankRepository) Edit(sharedBank *model.SharedBank, user *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(sharedBank))
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

	err = c.stor.EditSharedBank(sharedBank)
	if err != nil {
		return
	}
	err = c.prepare(sharedBank, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare shared bank")
		return
	}
	return
}

//Delete handles logic
func (c *SharedBankRepository) Delete(sharedBank *model.SharedBank, user *model.User) (err error) {
	err = c.stor.DeleteSharedBank(sharedBank)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *SharedBankRepository) ListByAccount(account *model.Account, pageSize int64, pageNumber int64, user *model.User) (sharedBanks []*model.SharedBank, err error) {
	if pageSize < 1 {
		pageSize = 25
	}

	if pageNumber < 0 {
		pageNumber = 0
	}

	sharedBanks, err = c.stor.ListSharedBankByAccount(account, pageSize, pageNumber)
	if err != nil {
		return
	}
	for _, sharedBank := range sharedBanks {
		err = c.prepare(sharedBank, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare shared bank")
			return
		}
	}
	return
}

//ListCount handles logic
func (c *SharedBankRepository) ListCountByAccount(account *model.Account, user *model.User) (count int64, err error) {

	count, err = c.stor.ListSharedBankByAccountCount(account)
	if err != nil {
		return
	}
	return
}

//ListByItem handles logic
func (c *SharedBankRepository) ListByItem(account *model.Account, item *model.Item, user *model.User) (sharedBanks []*model.SharedBank, err error) {
	sharedBanks, err = c.stor.ListSharedBankByAccountAndItem(account, item)
	if err != nil {
		return
	}
	for _, sharedBank := range sharedBanks {
		err = c.prepare(sharedBank, user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare shared bank")
			return
		}
	}
	return
}

func (c *SharedBankRepository) prepare(sharedBank *model.SharedBank, user *model.User) (err error) {

	return
}

func (c *SharedBankRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *SharedBankRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
