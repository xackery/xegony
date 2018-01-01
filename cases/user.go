package cases

import (
	"fmt"

	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
	"github.com/xeipuuv/gojsonschema"
)

type UserRepository struct {
	stor storage.Storage
}

func (g *UserRepository) Initialize(stor storage.Storage) (err error) {
	if stor == nil {
		err = fmt.Errorf("Invalid storage type")
		return
	}
	g.stor = stor
	return
}

func (g *UserRepository) Get(userID int64) (user *model.User, err error) {
	if userID == 0 {
		err = fmt.Errorf("Invalid User ID")
		return
	}
	user, err = g.stor.GetUser(userID)
	return
}

func (g *UserRepository) Create(user *model.User) (err error) {
	if user == nil {
		err = fmt.Errorf("Empty user")
		return
	}
	schema, err := g.newSchema([]string{"name", "password", "email", "accountID"}, nil)
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
	err = g.stor.CreateUser(user)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) Login(username string, password string) (user *model.User, err error) {
	user = &model.User{
		Name:     username,
		Password: password,
	}
	schema, err := g.newSchema([]string{"name", "password"}, nil)
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
	user, err = g.stor.LoginUser(user.Name, user.Password)
	if err != nil {
		return
	}

	return
}

func (g *UserRepository) Edit(userID int64, user *model.User) (err error) {
	schema, err := g.newSchema([]string{"name"}, nil)
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

	err = g.stor.EditUser(userID, user)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) Delete(userID int64) (err error) {
	err = g.stor.DeleteUser(userID)
	if err != nil {
		return
	}
	return
}

func (g *UserRepository) List() (users []*model.User, err error) {
	users, err = g.stor.ListUser()
	if err != nil {
		return
	}
	return
}

func (u *UserRepository) newSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = u.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = u.getSchemaProperty(field); err != nil {
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

func (u *UserRepository) getSchemaProperty(field string) (prop model.Schema, err error) {
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
