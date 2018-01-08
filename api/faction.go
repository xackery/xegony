package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) factionRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListFaction",
			"GET",
			"/faction",
			a.listFaction,
		},
		{
			"CreateFaction",
			"POST",
			"/faction",
			a.createFaction,
		},
		{
			"DeleteFaction",
			"DELETE",
			"/faction/{factionID}",
			a.deleteFaction,
		},
		{
			"EditFaction",
			"PUT",
			"/faction/{factionID}",
			a.editFaction,
		},
		{
			"GetFaction",
			"GET",
			"/faction/{factionID}",
			a.getFaction,
		},
	}
	return
}

func (a *API) getFaction(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	faction, err := a.factionRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, faction, http.StatusOK)
	return
}

func (a *API) createFaction(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	faction := &model.Faction{}
	err = decodeBody(r, faction)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.factionRepo.Create(faction)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, faction, http.StatusCreated)
	return
}

func (a *API) deleteFaction(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.factionRepo.Delete(id)
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

func (a *API) editFaction(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	faction := &model.Faction{}
	err = decodeBody(r, faction)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.factionRepo.Edit(id, faction)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, faction, http.StatusOK)
	return
}

func (a *API) listFaction(w http.ResponseWriter, r *http.Request) {
	factions, err := a.factionRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, factions, http.StatusOK)
	return
}
