package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//AccountRepository handles AccountRepository cases and is a gateway to storage
type AccountRepository struct {
	stor storage.Storage
}

//Initialize handles logic
func (c *AccountRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

//Get handles logic
func (c *AccountRepository) Get(accountID int64) (account *model.Account, err error) {
	if accountID == 0 {
		err = fmt.Errorf("Invalid Account ID")
		return
	}
	account, err = c.stor.GetAccount(accountID)
	return
}

//Create handles logic
func (c *AccountRepository) Create(account *model.Account) (err error) {
	if account == nil {
		err = fmt.Errorf("Empty account")
		return
	}
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	account.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(account))
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
	err = c.stor.CreateAccount(account)
	if err != nil {
		return
	}
	return
}

//Edit handles logic
func (c *AccountRepository) Edit(accountID int64, account *model.Account) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(account))
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

	err = c.stor.EditAccount(accountID, account)
	if err != nil {
		return
	}
	return
}

//Delete handles logic
func (c *AccountRepository) Delete(accountID int64) (err error) {
	err = c.stor.DeleteAccount(accountID)
	if err != nil {
		return
	}
	return
}

//List handles logic
func (c *AccountRepository) List() (accounts []*model.Account, err error) {
	accounts, err = c.stor.ListAccount()
	if err != nil {
		return
	}
	return
}

func (c *AccountRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

func (c *AccountRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "status":
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
