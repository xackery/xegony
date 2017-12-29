package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetLootTableEntry(w http.ResponseWriter, r *http.Request) {

	lootTableId, err := getIntVar(r, "lootTableId")
	if err != nil {
		err = errors.Wrap(err, "lootTableId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootTableEntry, err := a.lootTableEntryRepo.Get(lootTableId, lootDropId)
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

func (a *Api) CreateLootTableEntry(w http.ResponseWriter, r *http.Request) {
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

func (a *Api) DeleteLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootTableId, err := getIntVar(r, "lootTableId")
	if err != nil {
		err = errors.Wrap(err, "lootTableId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.lootTableEntryRepo.Delete(lootTableId, lootDropId)
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

func (a *Api) EditLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}
	lootTableId, err := getIntVar(r, "lootTableId")
	if err != nil {
		err = errors.Wrap(err, "lootTableId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
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

	err = a.lootTableEntryRepo.Edit(lootTableId, lootDropId, lootTableEntry)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootTableEntry, http.StatusOK)
	return
}

func (a *Api) ListLootTableEntry(w http.ResponseWriter, r *http.Request) {

	lootTableId, err := getIntVar(r, "lootTableId")
	if err != nil {
		err = errors.Wrap(err, "lootTableId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTableEntrys, err := a.lootTableEntryRepo.List(lootTableId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootTableEntrys, http.StatusOK)
	return
}
