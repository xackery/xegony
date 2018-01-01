package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getActivity(w http.ResponseWriter, r *http.Request) {

	activityID, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activity, err := a.activityRepo.Get(taskID, activityID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, activity, http.StatusOK)
	return
}

func (a *API) createActivity(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	activity := &model.Activity{}
	err = decodeBody(r, activity)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.activityRepo.Create(activity)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, activity, http.StatusCreated)
	return
}

func (a *API) deleteActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.activityRepo.Delete(id)
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

func (a *API) editActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	activity := &model.Activity{}
	err = decodeBody(r, activity)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.activityRepo.Edit(id, activity)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, activity, http.StatusOK)
	return
}

func (a *API) listActivity(w http.ResponseWriter, r *http.Request) {
	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activitys, err := a.activityRepo.List(taskID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, activitys, http.StatusOK)
	return
}
