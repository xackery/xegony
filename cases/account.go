package cases

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
	"github.com/xeipuuv/gojsonschema"
)

//ListAccount lists all accounts accessible by provided user
func ListAccount(page *model.Page, user *model.User) (accounts []*model.Account, err error) {
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	reader, err := getReader("account")
	if err != nil {
		err = errors.Wrap(err, "failed to prepare reader for account")
		return
	}

	page.Total, err = reader.ListAccountTotalCount()
	if err != nil {
		err = errors.Wrap(err, "failed to list account toal count")
		return
	}

	accounts, err = reader.ListAccount(page)
	if err != nil {
		err = errors.Wrap(err, "failed to list account")
		return
	}
	for i, account := range accounts {
		err = sanitizeAccount(account, user)
		if err != nil {
			err = errors.Wrapf(err, "failed to sanitize account element %d", i)
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

//ListAccountByName will request any account matching the pattern of name
func ListAccountBySearch(page *model.Page, account *model.Account, user *model.User) (accounts []*model.Account, err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list account by search without guide+")
		return
	}
	err = preparePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare page")
		return
	}

	err = prepareAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepre account")
		return
	}

	err = validateAccount(account, nil, []string{ //optional
		"name",
	})
	if err != nil {
		err = errors.Wrap(err, "failed to validate account")
		return
	}
	reader, err := getReader("account")
	if err != nil {
		err = errors.Wrap(err, "failed to get account reader")
		return
	}

	accounts, err = reader.ListAccountBySearch(page, account)
	if err != nil {
		err = errors.Wrap(err, "failed to list account by search")
		return
	}

	for _, account := range accounts {
		err = sanitizeAccount(account, user)
		if err != nil {
			err = errors.Wrap(err, "failed to sanitize account")
			return
		}
	}

	err = sanitizeAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize search account")
		return
	}

	err = sanitizePage(page, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize page")
		return
	}
	return
}

//CreateAccount will create an account using provided information
func CreateAccount(account *model.Account, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list account by search without guide+")
		return
	}
	err = prepareAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare account")
		return
	}

	err = validateAccount(account, []string{"name"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate account")
		return
	}
	account.ID = 0
	account.TimeCreation = time.Now().Unix()
	writer, err := getWriter("account")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for account")
		return
	}
	err = writer.CreateAccount(account)
	if err != nil {
		err = errors.Wrap(err, "failed to create account")
		return
	}
	err = sanitizeAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize account")
		return
	}
	return
}

//GetAccount gets an account by provided accountID
func GetAccount(account *model.Account, user *model.User) (err error) {
	err = prepareAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare account")
		return
	}

	err = validateAccount(account, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate account")
		return
	}

	reader, err := getReader("account")
	if err != nil {
		err = errors.Wrap(err, "failed to get account reader")
		return
	}

	err = reader.GetAccount(account)
	if err != nil {
		err = errors.Wrap(err, "failed to get account")
		return
	}

	err = sanitizeAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize account")
		return
	}

	return
}

//EditAccount edits an existing account
func EditAccount(account *model.Account, user *model.User) (err error) {
	err = user.IsGuide()
	if err != nil {
		err = errors.Wrap(err, "can't list account by search without guide+")
		return
	}
	err = prepareAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare account")
		return
	}

	err = validateAccount(account,
		[]string{"ID"}, //required
		[]string{ //optional
			"name",
			"charname",
			"sharedplat",
			"password",
			"status",
			"lsaccountID",
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
		err = errors.Wrap(err, "failed to validate account")
		return
	}
	writer, err := getWriter("account")
	if err != nil {
		err = errors.Wrap(err, "failed to get writer for account")
		return
	}
	err = writer.EditAccount(account)
	if err != nil {
		err = errors.Wrap(err, "failed to edit account")
		return
	}
	err = sanitizeAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to sanitize account")
		return
	}
	return
}

//DeleteAccount deletes an account by provided accountID
func DeleteAccount(account *model.Account, user *model.User) (err error) {
	err = user.IsAdmin()
	if err != nil {
		err = errors.Wrap(err, "can't delete account without admin+")
		return
	}
	err = prepareAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "failed to prepare account")
		return
	}

	err = validateAccount(account, []string{"ID"}, nil)
	if err != nil {
		err = errors.Wrap(err, "failed to validate account")
		return
	}
	writer, err := getWriter("account")
	if err != nil {
		err = errors.Wrap(err, "failed to get account writer")
		return
	}
	err = writer.DeleteAccount(account)
	if err != nil {
		err = errors.Wrap(err, "failed to delete account")
		return
	}
	return
}

func prepareAccount(account *model.Account, user *model.User) (err error) {
	if account == nil {
		err = fmt.Errorf("empty account")
		return
	}
	if user == nil {
		err = fmt.Errorf("empty user")
		return
	}
	return
}

func validateAccount(account *model.Account, required []string, optional []string) (err error) {
	schema, err := newSchemaAccount(required, optional)
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
	return
}

func sanitizeAccount(account *model.Account, user *model.User) (err error) {
	account.Password = ""
	err = user.IsGuide()
	if err != nil {
		account.Name = ""
		//account.Sharedplat = ""
		account.Status = 0
		account.Gmspeed = 0
		//account.Revoked = 0
		//account.Karma = 0
		account.MiniloginIP = ""
		account.Hideme = 0
		//account.Rulesflag = 0
		//account.Suspendeduntil =
		//account.TimeCreation = ""
		//account.Expansion = ""
		//account.BanReason = ""
		//account.SuspendReason = ""
		err = nil
	}
	return
}

func newSchemaAccount(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := model.Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]model.Schema)
	var field string
	var prop model.Schema
	for _, field = range requiredFields {
		if prop, err = getSchemaPropertyAccount(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = getSchemaPropertyAccount(field); err != nil {
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

func getSchemaPropertyAccount(field string) (prop model.Schema, err error) {
	switch field {
	case "ID":
		prop.Type = "integer"
		prop.Minimum = 1
	case "name":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "charname": //string `json:"
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 64
		prop.Pattern = "^[a-zA-Z]*$"
	case "sharedplat": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "password": //string `json:"
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
	case "status": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "lsaccountID": //sql.NullInt64 `
		prop.Type = "integer"
		prop.Minimum = 1
	case "gmspeed": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "revoked": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "karma": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "miniloginIp": //string `json:"
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "hideme": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "rulesflag": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "suspendeduntil": //time.Time `

	case "timeCreation": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "expansion": //int64 `json:"
		prop.Type = "integer"
		prop.Minimum = 1
	case "banReason": //sql.NullString `
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	case "suspendReason": //sql.NullString `
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 30
		prop.Pattern = "^[a-zA-Z]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
