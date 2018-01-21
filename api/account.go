package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) accountRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateAccount",
			"POST",
			"/account",
			a.createAccount,
		},
		{
			"DeleteAccount",
			"DELETE",
			"/account/{accountID:[0-9]+}",
			a.deleteAccount,
		},
		{
			"EditAccount",
			"PUT",
			"/account/{accountID}:[0-9]+",
			a.editAccount,
		},
		{
			"GetAccount",
			"GET",
			"/account/{accountID}:[0-9]+",
			a.getAccount,
		},
		{
			"ListAccount",
			"GET",
			"/account",
			a.listAccount,
		},
	}
	return
}

func (a *API) getAccount(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	accountID, err := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}
	account := &model.Account{
		ID: accountID,
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

func (a *API) createAccount(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

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

func (a *API) deleteAccount(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	accountID, err := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}
	account := &model.Account{
		ID: accountID,
	}

	err = a.accountRepo.Delete(account, auth.User)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = account
	return
}

func (a *API) editAccount(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	accountID, err := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}

	account := &model.Account{
		ID: accountID,
	}
	err = decodeBody(r, account)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = a.accountRepo.Edit(account, auth.User)
	if err != nil {
		return
	}
	content = account
	return
}

func (a *API) listAccount(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	accounts, err := a.accountRepo.List(auth.User)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = accounts
	return
}
