package api

import (
	"database/sql"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) activityRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListActivity",
			"GET",
			"/task/{taskID}",
			a.listActivity,
		},
		{
			"GetActivity",
			"GET",
			"/task/{taskID}/{activityID}",
			a.getActivity,
		},
		{
			"CreateActivity",
			"POST",
			"/task/{taskID}",
			a.createActivity,
		},
	}
	return
}

func (a *API) getActivity(w http.ResponseWriter, r *http.Request) {

	if strings.ToLower(getVar(r, "activityID")) == "details" {
		a.getTask(w, r)
		return
	}

	activityID, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activity, err := a.activityRepo.Get(taskID, activityID)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, activity, http.StatusOK)
	return
}

func (a *API) createActivity(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	activity := &model.Activity{}
	err = decodeBody(r, activity)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.activityRepo.Create(activity)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, activity, http.StatusCreated)
	return
}

func (a *API) deleteActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.activityRepo.Delete(id)
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

func (a *API) editActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	activity := &model.Activity{}
	err = decodeBody(r, activity)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.activityRepo.Edit(id, activity)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, activity, http.StatusOK)
	return
}

func (a *API) listActivity(w http.ResponseWriter, r *http.Request) {
	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activitys, err := a.activityRepo.List(taskID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, activitys, http.StatusOK)
	return
}
