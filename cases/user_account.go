package cases

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListUserAccount lists all userAccounts accessible by provided user
func ListUserAccount(page *model.Page, focusUser *model.User, user *model.User) (userAccounts []*model.UserAccount, err error) {
	err = validateOrderByUserAccountField(page)
	if err != nil {
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("userAccount")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for userAccount")
		return
	}

	page.Total, err = reader.ListUserAccountTotalCount(focusUser)
	if err != nil {
		err = errors.Wrap(err, "failed to list userAccount total count")
		return
	}

	userAccounts, err = reader.ListUserAccount(page, focusUser)
	if err != nil {
		err = errors.Wrap(err, "failed to list userAccount")
		return
	}
	for i, userAccount := range userAccounts {
		err = sanitizeUserAccount(userAccount, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize userAccount element %d", i)
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

//ListUserAccountBySearch will request any userAccount matching the pattern of name
func ListUserAccountBySearch(page *model.Page, focusUser *model.User, userAccount *model.UserAccount, user *model.User) (userAccounts []*model.UserAccount, err error) {

	err = validateOrderByUserAccountField(page)
	if err != nil {
		return
	}

	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre userAccount")
		return
	}

	err = validateUserAccount(userAccount, nil, []string{ //optional
		"accountID",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate userAccount")
		return
	}
	reader, err := getReader("userAccount")
	if err != nil {
		err = errors.Wrap(err, "failed to get userAccount reader")
		return
	}

	userAccounts, err = reader.ListUserAccountBySearch(page, focusUser, userAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to list userAccount by search")
		return
	}

	page.Total, err = reader.ListUserAccountBySearchTotalCount(focusUser, userAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to list userAccount by search total count")
		return
	}
	for _, userAccount := range userAccounts {
		err = sanitizeUserAccount(userAccount, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize userAccount")
			return
		}
	}

	err = sanitizeUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search userAccount")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateUserAccount will create an userAccount using provided information
func CreateUserAccount(focusUser *model.User, userAccount *model.UserAccount, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list userAccount by search without guide+")
		return
	}
	err = prepareUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare userAccount")
		return
	}

	err = validateUserAccount(userAccount, []string{"accountID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate userAccount")
		return
	}
	userAccount.UserID = focusUser.ID
	//userAccount.TimeCreation = time.Now().Unix()
	writer, err := getWriter("userAccount")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for userAccount")
		return
	}
	err = writer.CreateUserAccount(focusUser, userAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to create userAccount")
		return
	}
	err = sanitizeUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize userAccount")
		return
	}
	return
}

//GetUserAccount gets an userAccount by provided userAccountID
func GetUserAccount(focusUser *model.User, userAccount *model.UserAccount, user *model.User) (err error) {
	err = prepareUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare userAccount")
		return
	}

	err = validateUserAccount(userAccount, []string{"accountID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate userAccount")
		return
	}

	reader, err := getReader("userAccount")
	if err != nil {
		err = errors.Wrap(err, "failed to get userAccount reader")
		return
	}

	err = reader.GetUserAccount(focusUser, userAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to get userAccount")
		return
	}

	err = sanitizeUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize userAccount")
		return
	}

	return
}

//EditUserAccount edits an existing userAccount
func EditUserAccount(focusUser *model.User, userAccount *model.UserAccount, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list userAccount by search without guide+")
		return
	}
	err = prepareUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare userAccount")
		return
	}

	err = validateUserAccount(userAccount,
		[]string{"accountID"}, //required
		[]string{              //optional
		},
	)
	if err != nil {
		err = errors.Wrap(err, "failed to validate userAccount")
		return
	}
	writer, err := getWriter("userAccount")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for userAccount")
		return
	}
	err = writer.EditUserAccount(focusUser, userAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to edit userAccount")
		return
	}
	err = sanitizeUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize userAccount")
		return
	}
	return
}

//DeleteUserAccount deletes an userAccount by provided userAccountID
func DeleteUserAccount(userAccount *model.UserAccount, focusUser *model.User, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete userAccount without admin+")
		return
	}
	err = prepareUserAccount(userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare userAccount")
		return
	}

	err = validateUserAccount(userAccount, []string{"accountID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate userAccount")
		return
	}
	writer, err := getWriter("userAccount")
	if err != nil {
		err = errors.Wrap(err, "failed to get userAccount writer")
		return
	}
	err = writer.DeleteUserAccount(focusUser, userAccount)
	if err != nil {
		err = errors.Wrap(err, "failed to delete userAccount")
		return
	}
	return
}

func prepareUserAccount(userAccount *model.UserAccount, user *model.User) (err error) {
	if userAccount == nil {
		err = fmt.Errorf("empty userAccount")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateUserAccount(userAccount *model.UserAccount, required []string, optional []string) (err error) {
	schema, err := newSchemaUserAccount(required, optional)
	if err != nil {
		return
	}

	result, err := schema.Validate(gojsonschema.NewGoLoader(userAccount))
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

func validateOrderByUserAccountField(page *model.Page) (err error) {
	if len(page.OrderBy) == 0 {
		page.OrderBy = "accountid"
	}

	validNames := []string{
		"accountid",
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

func sanitizeUserAccount(userAccount *model.UserAccount, user *model.User) (err error) {

	userAccount.Account = &model.Account{
		ID: userAccount.AccountID,
	}

	err = GetAccount(userAccount.Account, user)
	if err != nil {
		err = errors.Wrapf(err, "failed to get account %d", userAccount.AccountID)
		return
	}
	return
}

func newSchemaUserAccount(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyUserAccount(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyUserAccount(field); err != nil {
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

func getSchemaPropertyUserAccount(field string) (prop model.Schema, err error) {
	switch field {
	case "accountID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "focusUserID":
		prop.Type = "integer"
		prop.Minimum = 0
	case "chance":
		prop.Type = "integer"
		prop.Minimum = 0
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
