package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) getSpawnEntry(w http.ResponseWriter, r *http.Request) {

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	spawnEntry, err := a.spawnEntryRepo.Get(spawnID, npcID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, spawnEntry, http.StatusOK)
	return
}

func (a *Api) createSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.spawnEntryRepo.Create(spawnEntry)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, spawnEntry, http.StatusCreated)
	return
}

func (a *Api) deleteSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.spawnEntryRepo.Delete(spawnID, npcID)
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

func (a *Api) editSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}
	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.spawnEntryRepo.Edit(spawnID, npcID, spawnEntry)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, spawnEntry, http.StatusOK)
	return
}

func (a *Api) listSpawnEntry(w http.ResponseWriter, r *http.Request) {

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spawnEntrys, err := a.spawnEntryRepo.List(spawnID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, spawnEntrys, http.StatusOK)
	return
}
