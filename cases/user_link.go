package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//GetUserLink gets an userLink by provided userLinkID
func GetUserLink(userLink *model.UserLink, user *model.User) (err error) {
	//Make sure logged in
	err = user.IsLoggedIn()
	if err != nil {
		return
	}

	err = prepareUserLink(userLink, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare userLink")
		return
	}

	err = validateUserLink(userLink, []string{"link"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate userLink")
		return
	}

	reader, err := getReader("userLink")
	if err != nil {
		err = errors.Wrap(err, "failed to get userLink reader")
		return
	}

	err = reader.GetUserLink(userLink)
	if err != nil {
		err = errors.Wrap(err, "failed to get userLink")
		return
	}

	err = sanitizeUserLink(userLink, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize userLink")
		return
	}
	//verified valid link, now let's get user

	userReader, err := getReader("user")
	if err != nil {
		err = errors.Wrap(err, "failed to get user reader")
		return
	}
	focusUser := &model.User{
		ID: user.ID,
	}
	err = userReader.GetUser(focusUser)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	//prep new user account creation
	userAccount := &model.UserAccount{
		UserID:      user.ID,
		AccountID:   userLink.AccountID,
		CharacterID: userLink.CharacterID,
	}
	userAccountWriter, err := getWriter("userAccount")
	if err != nil {
		err = errors.Wrap(err, "failed to get userAccount reader")
		return
	}

	err = userAccountWriter.CreateUserAccount(focusUser, userAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to create user account")
		return
	}

	userWriter, err := getWriter("user")
	if err != nil {
		err = errors.Wrap(err, "failed to get user writer")
		return
	}
	if focusUser.PrimaryAccountID == 0 {
		focusUser.PrimaryAccountID = userAccount.AccountID
		focusUser.PrimaryCharacterID = userAccount.CharacterID
		err = userWriter.EditUser(focusUser)
		if err != nil {
			err = errors.Wrap(err, "failed to edit user")
			return
		}
	}
	//now let's delete user link
	userLinkWriter, err := getWriter("userLink")
	err = userLinkWriter.DeleteUserLinkByAccount(&model.Account{ID: userLink.AccountID})
	if err != nil {
		err = errors.Wrap(err, "failed delete user link")
		return
	}

	return
}

//DeleteUserLinkByAccount deletes an userLink by provided accountID
func DeleteUserLinkByAccount(account *model.Account, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete userLink without admin+")
		return
	}
	writer, err := getWriter("userLink")
	if err != nil {
		err = errors.Wrap(err, "failed to get userLink writer")
		return
	}
	err = writer.DeleteUserLinkByAccount(account)
	if err != nil {
		err = errors.Wrap(err, "failed to delete userLink")
		return
	}

	return
}

func prepareUserLink(userLink *model.UserLink, user *model.User) (err error) {
	if userLink == nil {
		err = fmt.Errorf("empty userLink")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateUserLink(userLink *model.UserLink, required []string, optional []string) (err error) {
	schema, err := newSchemaUserLink(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(userLink))
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

func validateOrderByUserLinkField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "link"
	}

	validNames := []string{
		"link",
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

func sanitizeUserLink(userLink *model.UserLink, user *model.User) (err error) {
	userLink.Account = &model.Account{
		ID: userLink.AccountID,
	}
	err = GetAccount(userLink.Account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get account")
		return
	}

	userLink.Character = &model.Character{
		ID: userLink.CharacterID,
	}
	err = GetCharacter(userLink.Character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get character")
		return
	}
	return
}

func newSchemaUserLink(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyUserLink(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyUserLink(field); err != nil {
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

func getSchemaPropertyUserLink(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "link":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
