package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

//UserRepository handles UserRepository cases and is a gateway to storage
type UserRepository struct {
	stor storage.Storage
}

// Initialize handles functions
func (c *UserRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	c.stor = stor
	return
}

// Get handles functions
func (c *UserRepository) Get(user *model.User, authUser *model.User) (err error) {

	err = c.stor.GetUser(user)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}
	err = c.prepare(user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}
	return
}

// Create handles functions
func (c *UserRepository) Create(user *model.User, authUser *model.User) (err error) {
	if user == nil {
		err = fmt.Errorf("Empty user")
		return
	}
	schema, err := c.newSchema([]string{"name", "password", "email", "accountID"}, nil)
	if err != nil {
		return
	}
	user.ID = 0 //strip ID
	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
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
	err = c.stor.CreateUser(user)
	if err != nil {
		return
	}
	err = c.prepare(user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}
	return
}

// Login handles functions
func (c *UserRepository) Login(user *model.User, passwordConfirm string) (err error) {
	schema, err := c.newSchema([]string{"name", "password"}, nil)
	if err != nil {
		return
	}
	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
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
	err = c.stor.LoginUser(user, passwordConfirm)
	if err != nil {
		return
	}
	err = c.prepare(user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}

	return
}

// Edit handles functions
func (c *UserRepository) Edit(user *model.User, authIser *model.User) (err error) {
	schema, err := c.newSchema([]string{"name"}, nil)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(user))
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

	err = c.stor.EditUser(user)
	if err != nil {
		return
	}
	err = c.prepare(user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}
	return
}

// Delete handles functions
func (c *UserRepository) Delete(user *model.User, authUser *model.User) (err error) {
	err = c.stor.DeleteUser(user)
	if err != nil {
		return
	}
	return
}

// List handles functions
func (c *UserRepository) List(user *model.User) (users []*model.User, err error) {
	users, err = c.stor.ListUser()
	if err != nil {
		return
	}
	for _, user := range users {
		err = c.prepare(user)
		if err != nil {
			err = errors.Wrap(err, "failed to prepare user")
			return
		}
	}
	return
}

func (c *UserRepository) prepare(user *model.User) (err error) {

	user.Password = ""
	return
}

// newSchema handles functions
func (c *UserRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
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

// getSchemaProperty handles functions
func (c *UserRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 32
		prop.Pattern = "^[a-zA-Z' ]*$"
	case "password":
		prop.Type = "string"
		prop.MinLength = 6
		prop.MaxLength = 32
		prop.Pattern = `^[a-zA-Z]\w{3,14}$`
	case "passwordConfirm":
		prop.Type = "string"
		prop.MinLength = 6
		prop.MaxLength = 32
		prop.Pattern = `^[a-zA-Z]\w{3,14}$`
	case "email":
		prop.Type = "email"
	case "accountID":
		prop.Type = "integer"
		prop.Minimum = 1
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
