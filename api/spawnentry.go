package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) spawnEntryRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateSpawnEntry",
			"POST",
			"/spawn",
			a.createSpawnEntry,
		},
		{
			"DeleteSpawnEntry",
			"DELETE",
			"/spawn/{spawnID}",
			a.deleteSpawnEntry,
		},
		{
			"EditSpawnEntry",
			"PUT",
			"/spawn/{spawnID}",
			a.editSpawnEntry,
		},
		{
			"GetSpawnEntry",
			"GET",
			"/spawn/{spawnID}",
			a.getSpawnEntry,
		},
		{
			"ListSpawnEntry",
			"GET",
			"/spawn",
			a.listSpawnEntry,
		},
	}
	return
}

func (a *API) getSpawnEntry(w http.ResponseWriter, r *http.Request) {

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	_, spawnEntry, err := a.spawnEntryRepo.Get(spawnID, npcID)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, spawnEntry, http.StatusOK)
	return
}

func (a *API) createSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	_, err = a.spawnEntryRepo.Create(spawnEntry)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, spawnEntry, http.StatusCreated)
	return
}

func (a *API) deleteSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	_, err = a.spawnEntryRepo.Delete(spawnID, npcID)
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

func (a *API) editSpawnEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}
	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	_, err = a.spawnEntryRepo.Edit(spawnID, npcID, spawnEntry)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, spawnEntry, http.StatusOK)
	return
}

func (a *API) listSpawnEntry(w http.ResponseWriter, r *http.Request) {

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spawnEntrys, _, err := a.spawnEntryRepo.List(spawnID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, spawnEntrys, http.StatusOK)
	return
}
