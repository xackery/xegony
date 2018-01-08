package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) zoneRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *API) getZone(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	zone, err := a.zoneRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, zone, http.StatusOK)
	return
}

func (a *API) createZone(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	zone := &model.Zone{}
	err = decodeBody(r, zone)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.zoneRepo.Create(zone)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, zone, http.StatusCreated)
	return
}

func (a *API) deleteZone(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.zoneRepo.Delete(id)
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

func (a *API) editZone(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	zone := &model.Zone{}
	err = decodeBody(r, zone)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.zoneRepo.Edit(id, zone)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, zone, http.StatusOK)
	return
}

func (a *API) listZone(w http.ResponseWriter, r *http.Request) {
	zones, err := a.zoneRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, zones, http.StatusOK)
	return
}
