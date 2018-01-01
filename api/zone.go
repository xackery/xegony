package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getZone(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	zone, err := a.zoneRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, zone, http.StatusOK)
	return
}

func (a *API) createZone(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	zone := &model.Zone{}
	err = decodeBody(r, zone)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.zoneRepo.Create(zone)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, zone, http.StatusCreated)
	return
}

func (a *API) deleteZone(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.zoneRepo.Delete(id)
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

func (a *API) editZone(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	zone := &model.Zone{}
	err = decodeBody(r, zone)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.zoneRepo.Edit(id, zone)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, zone, http.StatusOK)
	return
}

func (a *API) listZone(w http.ResponseWriter, r *http.Request) {
	zones, err := a.zoneRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, zones, http.StatusOK)
	return
}
