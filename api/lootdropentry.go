package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetLootDropEntry(w http.ResponseWriter, r *http.Request) {

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemId, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntry, err := a.lootDropEntryRepo.Get(lootDropId, itemId)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, lootDropEntry, http.StatusOK)
	return
}

func (a *Api) CreateLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootDropEntry := &model.LootDropEntry{}
	err = decodeBody(r, lootDropEntry)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.lootDropEntryRepo.Create(lootDropEntry)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, lootDropEntry, http.StatusCreated)
	return
}

func (a *Api) DeleteLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemId, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.lootDropEntryRepo.Delete(lootDropId, itemId)
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

func (a *Api) EditLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemId, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntry := &model.LootDropEntry{}
	err = decodeBody(r, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.lootDropEntryRepo.Edit(lootDropId, itemId, lootDropEntry)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootDropEntry, http.StatusOK)
	return
}

func (a *Api) ListLootDropEntry(w http.ResponseWriter, r *http.Request) {
	itemId, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntrys, err := a.lootDropEntryRepo.List(itemId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, lootDropEntrys, http.StatusOK)
	return
}
