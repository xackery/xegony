package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// UserAccountRequest is a list of parameters used for userAccount
// swagger:parameters deleteUserAccount editUserAccount getUserAccount
type UserAccountRequest struct {
	// UserID to get information about
	// in: path
	// example: 1
	UserID int64 `json:"userID"`
	// AccountID to get information about
	// in: path
	// example: 55091
	AccountID int64 `json:"accountID"`
}

// UserAccountResponse is what endpoints respond with
// swagger:response
type UserAccountResponse struct {
	User        *model.User        `json:"user,omitempty"`
	UserAccount *model.UserAccount `json:"userAccount,omitempty"`
}

// UserAccountCreateRequest is the body parameters for creating an userAccount
// swagger:parameters createUserAccount
type UserAccountCreateRequest struct {
	// UserID to get information about
	// in: path
	// example: 1
	UserID int64 `json:"userID"`
	// UserAccount details to create
	// in: body
	UserAccount *model.UserAccount `json:"userAccount"`
}

// UserAccountEditRequest is the body parameters for creating an userAccount
// swagger:parameters editUserAccount
type UserAccountEditRequest struct {
	// UserID to get information about
	// in: path
	// example: 1
	UserID int64 `json:"UserID"`
	// AccountID to get information about
	// in: path
	// example: 55091
	AccountID int64 `json:"accountID"`
	// UserAccount details to edit
	// in: body
	UserAccount *model.UserAccount `json:"userAccount"`
}

// UserAccountsRequest is a list of parameters used for userAccount
// swagger:parameters listUserAccount
type UserAccountsRequest struct {
	// UserID to get information about
	// in: path
	// example: 1
	UserID int64 `json:"userID"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: account_id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// UserAccountsResponse is a general response to a request
// swagger:response
type UserAccountsResponse struct {
	Page         *model.Page        `json:"page,omitempty"`
	User         *model.User        `json:"user"`
	UserAccounts model.UserAccounts `json:"userAccounts,omitempty"`
}

// UserAccountsBySearchRequest is a list of parameters used for userAccount
// swagger:parameters listUserAccountBySearch
type UserAccountsBySearchRequest struct {
	// UserID to get information about
	// in: path
	// example: 1
	UserID int64 `json:"userID"`
	// AccountID is which userAccount to get information about
	// example: 55091
	// in: query
	AccountID int64 `json:"accountID"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: account_id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// UserAccountsBySearchResponse is a general response to a request
// swagger:response
type UserAccountsBySearchResponse struct {
	Search       *model.UserAccount `json:"search,omitempty"`
	Page         *model.Page        `json:"page,omitempty"`
	User         *model.User        `json:"user,omitempty"`
	UserAccounts model.UserAccounts `json:"userAccounts,omitempty"`
}

func userAccountRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /user/{userID}/account user listUserAccount
		//
		// Lists userAccounts
		//
		// This will show all available userAccounts by default.
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
		//       200: UserAccountsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListUserAccount",
			"GET",
			"/user/{userID:[0-9]+}/account",
			listUserAccount,
		},
		// swagger:route GET /user/{userID}/account/search user listUserAccountBySearch
		//
		// Search userAccounts by accountid
		//
		// This will show all available userAccounts by default.
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
		//       200: UserAccountsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListUserAccountBySearch",
			"GET",
			"/user/account/search",
			listUserAccountBySearch,
		},
		// swagger:route POST /user/{userID}/account/{accountID} user createUserAccount
		//
		// Create an userAccount
		//
		// This will create an userAccount
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: UserAccountResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateUserAccount",
			"POST",
			"/user/{userID:[0-9]+}/account/{accountID:[0-9]+}",
			createUserAccount,
		},
		// swagger:route GET /user/{userID}/account/{accountID} user getUserAccount
		//
		// Get an userAccount
		//
		// This will get an individual userAccount available userAccounts by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: UserAccountResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetUserAccount",
			"GET",
			"/user/{userID:[0-9]+}/account/{accountID:[0-9]+}",
			getUserAccount,
		},
		// swagger:route PUT /user/{userID}/account/{accountID} user editUserAccount
		//
		// Edit an userAccount
		//
		// This will edit an userAccount
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: UserAccountResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditUserAccount",
			"PUT",
			"/user/{userID:[0-9]+}/account/{accountID:[0-9]+}",
			editUserAccount,
		},
		// swagger:route DELETE /user/{userID}/account/{accountID} user deleteUserAccount
		//
		// Delete an userAccount
		//
		// This will delete an userAccount
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
			"DeleteUserAccount",
			"DELETE",
			"/user/{userID:[0-9]+}/account/{accountID:[0-9]+}",
			deleteUserAccount,
		},
	}
	return
}

func getUserAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserAccountRequest{
		UserID:    getIntVar(r, "userID"),
		AccountID: getIntVar(r, "accountID"),
	}

	focusUser := &model.User{
		ID: request.UserID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	userAccount := &model.UserAccount{
		UserID:    request.UserID,
		AccountID: request.AccountID,
	}

	err = cases.GetUserAccount(focusUser, userAccount, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &UserAccountResponse{
		User:        user,
		UserAccount: userAccount,
	}
	content = response
	return
}

func createUserAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserAccountCreateRequest{
		UserID: getIntVar(r, "userID"),
	}

	focusUser := &model.User{
		ID: request.UserID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	userAccount := &model.UserAccount{}
	err = decodeBody(r, userAccount)
	if err != nil {
		return
	}

	err = cases.CreateUserAccount(focusUser, userAccount, user)
	if err != nil {
		return
	}
	response := &UserAccountResponse{
		User:        focusUser,
		UserAccount: userAccount,
	}
	content = response
	return
}

func deleteUserAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserAccountRequest{
		UserID:    getIntVar(r, "userID"),
		AccountID: getIntVar(r, "accountID"),
	}

	focusUser := &model.User{
		ID: request.UserID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	userAccount := &model.UserAccount{
		UserID:    request.UserID,
		AccountID: request.AccountID,
	}

	err = cases.DeleteUserAccount(userAccount, focusUser, user)
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

func editUserAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserAccountEditRequest{
		UserID:    getIntVar(r, "userID"),
		AccountID: getIntVar(r, "accountID"),
	}

	focusUser := &model.User{
		ID: request.UserID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	userAccount := &model.UserAccount{}

	err = decodeBody(r, userAccount)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	userAccount.AccountID = request.AccountID
	userAccount.UserID = request.UserID

	err = cases.EditUserAccount(focusUser, userAccount, user)
	if err != nil {
		return
	}
	response := &UserAccountResponse{
		User:        focusUser,
		UserAccount: userAccount,
	}
	content = response
	return
}

func listUserAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserAccountsRequest{
		UserID: getIntVar(r, "userID"),
	}

	focusUser := &model.User{
		ID: request.UserID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	userAccounts, err := cases.ListUserAccount(page, focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &UserAccountsResponse{
		Page:         page,
		User:         focusUser,
		UserAccounts: userAccounts,
	}
	content = response
	return
}

func listUserAccountBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserAccountsBySearchRequest{
		UserID: getIntVar(r, "userID"),
	}

	focusUser := &model.User{
		ID: request.UserID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get user")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	userAccount := &model.UserAccount{
		UserID: request.UserID,
	}
	userAccount.AccountID = getIntQuery(r, "accountID")

	userAccounts, err := cases.ListUserAccountBySearch(page, focusUser, userAccount, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &UserAccountsBySearchResponse{
		Page:         page,
		User:         focusUser,
		UserAccounts: userAccounts,
		Search:       userAccount,
	}
	content = response
	return
}
