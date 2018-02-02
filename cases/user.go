package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListUser lists all users accessible by provided user
func ListUser(page *model.Page, user *model.User) (users []*model.User, err error) {
	err = validateOrderByUserField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("user")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for user")
		return
	}

	page.Total, err = reader.ListUserTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list user toal count")
		return
	}

	users, err = reader.ListUser(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list user")
		return
	}
	for i, focusUser := range users {
		err = sanitizeUser(focusUser, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize user element %d", i)
			return
		}
	}
	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}

	return
}

//ListUserBySearch will request any user matching the pattern of name
func ListUserBySearch(page *model.Page, focusUser *model.User, user *model.User) (users []*model.User, err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list user by search without guide+")
		return
	}

	err = validateOrderByUserField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre user")
		return
	}

	err = validateUser(focusUser, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate user")
		return
	}
	reader, err := getReader("user")
	if err != nil {
		err = errors.Wrap(err, "failed to get user reader")
		return
	}

	users, err = reader.ListUserBySearch(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to list user by search")
		return
	}

	for _, tmpUser := range users {
		err = sanitizeUser(tmpUser, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize user")
			return
		}
	}

	err = sanitizeUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search user")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateUser will create an user using provided information
func CreateUser(focusUser *model.User, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list user by search without guide+")
		return
	}
	err = prepareUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}

	err = validateUser(focusUser, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate user")
		return
	}
	focusUser.ID = 0
	writer, err := getWriter("user")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for user")
		return
	}
	err = writer.CreateUser(focusUser)
	if err != nil {
		err = errors.Wrap(err, "failed to create user")
		return
	}
	err = sanitizeUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize user")
		return
	}
	return
}

//GetUser gets an user by provided userID
func GetUser(focusUser *model.User, user *model.User) (err error) {
	err = prepareUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}

	err = validateUser(focusUser, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate user")
		return
	}

	reader, err := getReader("user")
	if err != nil {
		err = errors.Wrap(err, "failed to get user reader")
		return
	}

	err = reader.GetUser(focusUser)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	err = sanitizeUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize user")
		return
	}

	return
}

//EditUser edits an existing user
func EditUser(focusUser *model.User, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list user by search without guide+")
		return
	}
	err = prepareUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}

	err = validateUser(focusUser,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"charname",
			"sharedplat",
			"password",
			"status",
			"lsuserID",
			"gmspeed",
			"revoked",
			"karma",
			"miniloginIp",
			"hideme",
			"rulesflag",
			"suspendeduntil",
			"timeCreation",
			"expansion",
			"banReason",
			"suspendReason"},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate user")
		return
	}
	writer, err := getWriter("user")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for user")
		return
	}
	err = writer.EditUser(focusUser)
	if err != nil {
		err = errors.Wrap(err, "failed to edit user")
		return
	}
	err = sanitizeUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize user")
		return
	}
	return
}

//DeleteUser deletes an user by provided userID
func DeleteUser(focusUser *model.User, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete user without admin+")
		return
	}
	err = prepareUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare user")
		return
	}

	err = validateUser(focusUser, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate user")
		return
	}
	writer, err := getWriter("user")
	if err != nil {
		err = errors.Wrap(err, "failed to get user writer")
		return
	}
	err = writer.DeleteUser(focusUser)
	if err != nil {
		err = errors.Wrap(err, "failed to delete user")
		return
	}
	return
}

func prepareUser(focusUser *model.User, user *model.User) (err error) {
	if focusUser == nil {
		err = fmt.Errorf("empty focus user")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateUser(user *model.User, required []string, optional []string) (err error) {
	schema, err := newSchemaUser(required, optional)
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
	return
}

func validateOrderByUserField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "id"
	}

	validNames := []string{
		"id",
		"display_name",
		"email",
	}

	possibleNames := ""
	for _, name := range validNames {
		if page.OrderBy == name {
			return
		}
		possibleNames += name + ", "
	}
	if len(possibleNames) > 0 {
		possibleNames = possibleNames[0 : len(possibleNames)-2]
	}

	err = &model.ErrValidation{
		Message: "orderBy is invalid. Possible fields: " + possibleNames,
		Reasons: map[string]string{
			"orderBy": "field is not valid",
		},
	}
	return
}

func sanitizeUser(focusUser *model.User, user *model.User) (err error) {
	focusUser.Password = ""
	err = user.IsGuide()
	if err != nil {

		err = nil
	}
	return
}

func newSchemaUser(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyUser(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyUser(field); err != nil {
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

func getSchemaPropertyUser(field string) (prop model.Schema, err error) {
	switch field {

	case "ID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "displayName":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "primaryAccountID": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "primaryCharacterID": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "email": //string `json:"
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
		prop.Pattern = "^[a-zA-Z]*$"
	case "password": //string `json:"
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	case "googleToken": //int64 `json:"

	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
