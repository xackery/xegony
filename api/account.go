package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

// swagger:parameters deleteAccount editAccount getAccount
type AccountParams struct {
	//AccountID to get information about
	// in: path
	AccountID int64 `json:"accountID"`
	//todo: pagination
}

func (a *API) accountRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /account account listAccount
		//
		// Lists accounts
		//
		// This will show all available accounts by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: Accounts
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListAccount",
			"GET",
			"/account",
			a.listAccount,
		},
		// swagger:route POST /account account createAccount
		//
		// Create an account
		//
		// This will create an account
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateAccount",
			"POST",
			"/account",
			a.createAccount,
		},
		// swagger:route GET /account/{accountID} account getAccount
		//
		// Get an account
		//
		// This will get an individual account available accounts by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: Account
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetAccount",
			"GET",
			"/account/{accountID:[0-9]+}",
			a.getAccount,
		},
		// swagger:route PUT /account/{accountID} account editAccount
		//
		// Edit an account
		//
		// This will edit an account
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ErrNoContent
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditAccount",
			"PUT",
			"/account/{accountID:[0-9]+}",
			a.editAccount,
		},
		// swagger:route DELETE /account/{accountID} account deleteAccount
		//
		// Delete an account
		//
		// This will delete an account
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"DeleteAccount",
			"DELETE",
			"/account/{accountID:[0-9]+}",
			a.deleteAccount,
		},
	}
	return
}

func (a *API) getAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	accountReq := &AccountParams{}

	accountReq.AccountID, err = getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}
	account := &model.Account{
		ID: accountReq.AccountID,
	}
	err = a.accountRepo.Get(account, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = account
	return
}

func (a *API) createAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	account := &model.Account{}
	err = decodeBody(r, account)
	if err != nil {
		return
	}
	err = a.accountRepo.Create(account, user)
	if err != nil {
		return
	}
	content = account
	return
}

func (a *API) deleteAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	accountReq := &AccountParams{}
	accountReq.AccountID, err = getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}
	account := &model.Account{
		ID: accountReq.AccountID,
	}

	err = a.accountRepo.Delete(account, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
	}
	err = &model.ErrNoContent{}
	return
}

func (a *API) editAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	accountReq := &AccountParams{}
	accountReq.AccountID, err = getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}

	account := &model.Account{
		ID: accountReq.AccountID,
	}
	err = decodeBody(r, account)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = a.accountRepo.Edit(account, user)
	if err != nil {
		return
	}
	content = account
	return
}

func (a *API) listAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	accounts, err := a.accountRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = accounts
	return
}
