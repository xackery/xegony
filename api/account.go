package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getAccount(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	account, err := a.accountRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, account, http.StatusOK)
	return
}

func (a *API) createAccount(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	account := &model.Account{}
	err = decodeBody(r, account)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.accountRepo.Create(account)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, account, http.StatusCreated)
	return
}

func (a *API) deleteAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.accountRepo.Delete(id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			a.writeData(w, r, nil, http.StatusNotModified)
			return
		default:
			err = errors.Wrap(err, "Request failed")
			a.writeError(w, r, err, http.StatusInternalServerError)
		}
		return
	}
	a.writeData(w, r, nil, http.StatusNoContent)
	return
}

func (a *API) editAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	account := &model.Account{}
	err = decodeBody(r, account)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.accountRepo.Edit(id, account)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, account, http.StatusOK)
	return
}

func (a *API) listAccount(w http.ResponseWriter, r *http.Request) {
	accounts, err := a.accountRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, accounts, http.StatusOK)
	return
}
