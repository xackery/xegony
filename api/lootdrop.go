package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) lootDropRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *API) getLootDrop(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootDrop, err := a.lootDropRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, lootDrop, http.StatusOK)
	return
}

func (a *API) createLootDrop(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootDrop := &model.LootDrop{}
	err = decodeBody(r, lootDrop)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.lootDropRepo.Create(lootDrop)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, lootDrop, http.StatusCreated)
	return
}

func (a *API) deleteLootDrop(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.lootDropRepo.Delete(id)
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

func (a *API) editLootDrop(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDrop := &model.LootDrop{}
	err = decodeBody(r, lootDrop)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.lootDropRepo.Edit(id, lootDrop)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, lootDrop, http.StatusOK)
	return
}

func (a *API) listLootDrop(w http.ResponseWriter, r *http.Request) {
	lootDrops, err := a.lootDropRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, lootDrops, http.StatusOK)
	return
}
