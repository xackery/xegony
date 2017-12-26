package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetItem(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	item, err := a.itemRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, item, http.StatusOK)
	return
}

func (a *Api) CreateItem(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	item := &model.Item{}
	err = decodeBody(r, item)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.itemRepo.Create(item)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, item, http.StatusCreated)
	return
}

func (a *Api) DeleteItem(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.itemRepo.Delete(id)
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

func (a *Api) EditItem(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	item := &model.Item{}
	err = decodeBody(r, item)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.itemRepo.Edit(id, item)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, item, http.StatusOK)
	return
}

func (a *Api) ListItem(w http.ResponseWriter, r *http.Request) {
	items, err := a.itemRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, items, http.StatusOK)
	return
}
