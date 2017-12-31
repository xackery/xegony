package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) getLootTableEntry(w http.ResponseWriter, r *http.Request) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootTableEntry, err := a.lootTableEntryRepo.Get(lootTableID, lootDropID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, lootTableEntry, http.StatusOK)
	return
}

func (a *Api) createLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootTableEntry := &model.LootTableEntry{}
	err = decodeBody(r, lootTableEntry)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.lootTableEntryRepo.Create(lootTableEntry)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, lootTableEntry, http.StatusCreated)
	return
}

func (a *Api) deleteLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.lootTableEntryRepo.Delete(lootTableID, lootDropID)
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

func (a *Api) editLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}
	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTableEntry := &model.LootTableEntry{}
	err = decodeBody(r, lootTableEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.lootTableEntryRepo.Edit(lootTableID, lootDropID, lootTableEntry)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootTableEntry, http.StatusOK)
	return
}

func (a *Api) listLootTableEntry(w http.ResponseWriter, r *http.Request) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTableEntrys, err := a.lootTableEntryRepo.List(lootTableID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootTableEntrys, http.StatusOK)
	return
}
