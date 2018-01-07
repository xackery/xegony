package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getBazaar(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	bazaar, err := a.bazaarRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, bazaar, http.StatusOK)
	return
}

func (a *API) createBazaar(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	bazaar := &model.Bazaar{}
	err = decodeBody(r, bazaar)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.bazaarRepo.Create(bazaar)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, bazaar, http.StatusCreated)
	return
}

func (a *API) deleteBazaar(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.bazaarRepo.Delete(id)
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

func (a *API) editBazaar(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	bazaar := &model.Bazaar{}
	err = decodeBody(r, bazaar)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.bazaarRepo.Edit(id, bazaar)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, bazaar, http.StatusOK)
	return
}

func (a *API) listBazaar(w http.ResponseWriter, r *http.Request) {
	bazaars, err := a.bazaarRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, bazaars, http.StatusOK)
	return
}
