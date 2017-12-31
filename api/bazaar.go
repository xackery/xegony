package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) getBazaar(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "bazaarId")
	if err != nil {
		err = errors.Wrap(err, "bazaarId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	bazaar, err := a.bazaarRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, bazaar, http.StatusOK)
	return
}

func (a *Api) createBazaar(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	bazaar := &model.Bazaar{}
	err = decodeBody(r, bazaar)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.bazaarRepo.Create(bazaar)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, bazaar, http.StatusCreated)
	return
}

func (a *Api) deleteBazaar(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "bazaarId")
	if err != nil {
		err = errors.Wrap(err, "bazaarId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.bazaarRepo.Delete(id)
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

func (a *Api) editBazaar(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "bazaarId")
	if err != nil {
		err = errors.Wrap(err, "bazaarId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	bazaar := &model.Bazaar{}
	err = decodeBody(r, bazaar)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.bazaarRepo.Edit(id, bazaar)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, bazaar, http.StatusOK)
	return
}

func (a *Api) listBazaar(w http.ResponseWriter, r *http.Request) {
	bazaars, err := a.bazaarRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, bazaars, http.StatusOK)
	return
}
