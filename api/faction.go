package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getFaction(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	faction, err := a.factionRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, faction, http.StatusOK)
	return
}

func (a *API) createFaction(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	faction := &model.Faction{}
	err = decodeBody(r, faction)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.factionRepo.Create(faction)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, faction, http.StatusCreated)
	return
}

func (a *API) deleteFaction(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.factionRepo.Delete(id)
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

func (a *API) editFaction(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	faction := &model.Faction{}
	err = decodeBody(r, faction)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.factionRepo.Edit(id, faction)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, faction, http.StatusOK)
	return
}

func (a *API) listFaction(w http.ResponseWriter, r *http.Request) {
	factions, err := a.factionRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, factions, http.StatusOK)
	return
}
