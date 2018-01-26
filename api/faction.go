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
			"/faction/{factionID:[0-9]+}",
			a.deleteFaction,
		},
		{
			"EditFaction",
			"PUT",
			"/faction/{factionID:[0-9]+}",
			a.editFaction,
		},
		{
			"GetFaction",
			"GET",
			"/faction/{factionID:[0-9]+}",
			a.getFaction,
		},
	}
	return
}

func (a *API) getFaction(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	factionID, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		return
	}
	faction := &model.Faction{
		ID: factionID,
	}
	err = a.factionRepo.Get(faction, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = faction
	return
}

func (a *API) createFaction(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	faction := &model.Faction{}
	err = decodeBody(r, faction)
	if err != nil {
		return
	}
	err = a.factionRepo.Create(faction, user)
	if err != nil {
		return
	}
	content = faction
	return
}

func (a *API) deleteFaction(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	factionID, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		return
	}

	faction := &model.Faction{
		ID: factionID,
	}
	err = a.factionRepo.Delete(faction, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = faction
	return
}

func (a *API) editFaction(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	factionID, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
		return
	}

	faction := &model.Faction{}
	err = decodeBody(r, faction)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	faction.ID = factionID

	err = a.factionRepo.Edit(faction, user)
	if err != nil {
		return
	}
	content = faction
	return
}

func (a *API) listFaction(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	factions, err := a.factionRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = factions
	return
}
