package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) lootDropEntryRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *API) getLootDropEntry(w http.ResponseWriter, r *http.Request) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntry, err := a.lootDropEntryRepo.Get(lootDropID, itemID)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, lootDropEntry, http.StatusOK)
	return
}

func (a *API) createLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootDropEntry := &model.LootDropEntry{}
	err = decodeBody(r, lootDropEntry)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.lootDropEntryRepo.Create(lootDropEntry)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, lootDropEntry, http.StatusCreated)
	return
}

func (a *API) deleteLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.lootDropEntryRepo.Delete(lootDropID, itemID)
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

func (a *API) editLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntry := &model.LootDropEntry{}
	err = decodeBody(r, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.lootDropEntryRepo.Edit(lootDropID, itemID, lootDropEntry)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, lootDropEntry, http.StatusOK)
	return
}

func (a *API) listLootDropEntry(w http.ResponseWriter, r *http.Request) {
	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntrys, err := a.lootDropEntryRepo.List(itemID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, lootDropEntrys, http.StatusOK)
	return
}
