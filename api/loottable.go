package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getLootTable(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootTable, err := a.lootTableRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, lootTable, http.StatusOK)
	return
}

func (a *API) createLootTable(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootTable := &model.LootTable{}
	err = decodeBody(r, lootTable)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.lootTableRepo.Create(lootTable)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, lootTable, http.StatusCreated)
	return
}

func (a *API) deleteLootTable(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.lootTableRepo.Delete(id)
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

func (a *API) editLootTable(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTable := &model.LootTable{}
	err = decodeBody(r, lootTable)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.lootTableRepo.Edit(id, lootTable)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootTable, http.StatusOK)
	return
}

func (a *API) listLootTable(w http.ResponseWriter, r *http.Request) {
	lootTables, err := a.lootTableRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootTables, http.StatusOK)
	return
}
