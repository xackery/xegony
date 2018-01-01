package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getSpawn(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	spawn, err := a.spawnRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, spawn, http.StatusOK)
	return
}

func (a *API) createSpawn(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	spawn := &model.Spawn{}
	err = decodeBody(r, spawn)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.spawnRepo.Create(spawn)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, spawn, http.StatusCreated)
	return
}

func (a *API) deleteSpawn(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.spawnRepo.Delete(id)
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

func (a *API) editSpawn(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spawn := &model.Spawn{}
	err = decodeBody(r, spawn)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.spawnRepo.Edit(id, spawn)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, spawn, http.StatusOK)
	return
}

func (a *API) listSpawn(w http.ResponseWriter, r *http.Request) {
	spawns, err := a.spawnRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, spawns, http.StatusOK)
	return
}
