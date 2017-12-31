package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) getAccount(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "accountId")
	if err != nil {
		err = errors.Wrap(err, "accountId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	account, err := a.accountRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, account, http.StatusOK)
	return
}

func (a *Api) createAccount(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	account := &model.Account{}
	err = decodeBody(r, account)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.accountRepo.Create(account)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, account, http.StatusCreated)
	return
}

func (a *Api) deleteAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "accountId")
	if err != nil {
		err = errors.Wrap(err, "accountId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.accountRepo.Delete(id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			writeData(w, r, nil, http.StatusNotModified)
			return
		default:
			err = errors.Wrap(err, "Request failed")
			writeError(w, r, err, http.StatusInternalServerError)
		}
		return
	}
	writeData(w, r, nil, http.StatusNoContent)
	return
}

func (a *Api) editAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "accountId")
	if err != nil {
		err = errors.Wrap(err, "accountId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	account := &model.Account{}
	err = decodeBody(r, account)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.accountRepo.Edit(id, account)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, account, http.StatusOK)
	return
}

func (a *Api) listAccount(w http.ResponseWriter, r *http.Request) {
	accounts, err := a.accountRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, accounts, http.StatusOK)
	return
}
