package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// AccountRequest is a list of parameters used for account
// swagger:parameters deleteAccount editAccount getAccount
type AccountRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// AccountResponse is what endpoints respond with
// swagger:response
type AccountResponse struct {
	Account *model.Account `json:"account"`
}

// AccountCreateRequest is the body parameters for creating an account
// swagger:parameters createAccount
type AccountCreateRequest struct {
	// Account details to create
	// in: body
	Account *model.Account `json:"account"`
}

// AccountEditRequest is the body parameters for creating an account
// swagger:parameters editAccount
type AccountEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Account details to edit
	// in: body
	Account *model.Account `json:"account"`
}

// AccountsRequest is a list of parameters used for account
// swagger:parameters listAccount
type AccountsRequest struct {
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
}

// AccountsResponse is a general response to a request
// swagger:response
type AccountsResponse struct {
	Page     *model.Page    `json:"page"`
	Accounts model.Accounts `json:"accounts"`
}

// AccountsBySearchRequest is a list of parameters used for account
// swagger:parameters listAccountBySearch
type AccountsBySearchRequest struct {
	// Name is which account to get information about
	// example: xackery
	// in: query
	Name string `json:"name"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
}

// AccountsBySearchResponse is a general response to a request
// swagger:response listAccountBySearch
type AccountsBySearchResponse struct {
	Search   *model.Account `json:"search"`
	Page     *model.Page    `json:"page"`
	Accounts model.Accounts `json:"accounts"`
}

func accountRoutes() (routes []*route) {

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
		//       200: AccountsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListAccount",
			"GET",
			"/account",
			listAccount,
		},
		// swagger:route GET /account/search account listAccountBySearch
		//
		// Search accounts by name
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
		//       200: AccountsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListAccountBySearch",
			"GET",
			"/account/search",
			listAccountBySearch,
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
		//		 200: AccountResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateAccount",
			"POST",
			"/account",
			createAccount,
		},
		// swagger:route GET /account/{ID} account getAccount
		//
		// Get an account
		//
		// This will get an individual account available accounts by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: AccountResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetAccount",
			"GET",
			"/account/{ID:[0-9]+}",
			getAccount,
		},
		// swagger:route PUT /account/{ID} account editAccount
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
		//		 200: AccountResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditAccount",
			"PUT",
			"/account/{ID:[0-9]+}",
			editAccount,
		},
		// swagger:route DELETE /account/{ID} account deleteAccount
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
			"/account/{ID:[0-9]+}",
			deleteAccount,
		},
	}
	return
}

func getAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &AccountRequest{
		ID: getIntVar(r, "ID"),
	}

	account := &model.Account{
		ID: request.ID,
	}

	err = cases.GetAccount(account, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &AccountResponse{
		Account: account,
	}
	content = response
	return
}

func createAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	account := &model.Account{}
	err = decodeBody(r, account)
	if err != nil {
		return
	}
	err = cases.CreateAccount(account, user)
	if err != nil {
		return
	}
	response := &AccountResponse{
		Account: account,
	}
	content = response
	return
}

func deleteAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &AccountRequest{
		ID: getIntVar(r, "ID"),
	}

	account := &model.Account{
		ID: request.ID,
	}

	err = cases.DeleteAccount(account, user)
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

func editAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &AccountEditRequest{
		ID: getIntVar(r, "ID"),
	}

	account := &model.Account{
		ID: request.ID,
	}

	err = decodeBody(r, account)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditAccount(account, user)
	if err != nil {
		return
	}
	response := &AccountResponse{
		Account: account,
	}
	content = response
	return
}

func listAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset: getIntQuery(r, "offset"),
		Limit:  getIntQuery(r, "limit"),
	}
	accounts, err := cases.ListAccount(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &AccountsResponse{
		Page:     page,
		Accounts: accounts,
	}
	content = response
	return
}

func listAccountBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset: getIntQuery(r, "offset"),
		Limit:  getIntQuery(r, "limit"),
	}
	account := &model.Account{
		Name: getQuery(r, "name"),
	}
	accounts, err := cases.ListAccountBySearch(page, account, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	log.Println(accounts)
	response := &AccountsBySearchResponse{
		Page:     page,
		Accounts: accounts,
		Search:   account,
	}
	content = response
	return
}
