package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//User is a login authentication created by Xegony, binding an account ID
type User struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	AccountID   int64  `json:"accountID" db:"account_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsAdmin     bool   `json:"isAdmin"`
	IsModerator bool   `json:"isModerator"`
}

func (u *User) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]Schema)
	var field string
	var prop Schema
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

func (u *User) getSchemaProperty(field string) (prop Schema, err error) {
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
