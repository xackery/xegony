package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type AccountRepository struct {
	stor storage.Storage
}

func (g *AccountRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *AccountRepository) Get(accountID int64) (account *model.Account, err error) {
	if accountID == 0 {
		err = fmt.Errorf("Invalid Account ID")
		return
	}
	account, err = g.stor.GetAccount(accountID)
	return
}

func (g *AccountRepository) Create(account *model.Account) (err error) {
	if account == nil {
		err = fmt.Errorf("Empty account")
		return
	}
	schema, err := g.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}
	account.Id = 0 //strip ID
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
	err = g.stor.CreateAccount(account)
	if err != nil {
		return
	}
	return
}

func (g *AccountRepository) Edit(accountID int64, account *model.Account) (err error) {
	schema, err := g.newSchema([]string{"name"}, nil)
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

	err = g.stor.EditAccount(accountID, account)
	if err != nil {
		return
	}
	return
}

func (g *AccountRepository) Delete(accountID int64) (err error) {
	err = g.stor.DeleteAccount(accountID)
	if err != nil {
		return
	}
	return
}

func (g *AccountRepository) List() (accounts []*model.Account, err error) {
	accounts, err = g.stor.ListAccount()
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
