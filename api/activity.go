package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetActivity(w http.ResponseWriter, r *http.Request) {

	activityId, err := getIntVar(r, "activityId")
	if err != nil {
		err = errors.Wrap(err, "activityId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	taskId, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activity, err := a.activityRepo.Get(taskId, activityId)
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

func (a *Api) CreateActivity(w http.ResponseWriter, r *http.Request) {
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

func (a *Api) DeleteActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "activityId")
	if err != nil {
		err = errors.Wrap(err, "activityId argument is required")
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

func (a *Api) EditActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "activityId")
	if err != nil {
		err = errors.Wrap(err, "activityId argument is required")
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

func (a *Api) ListActivity(w http.ResponseWriter, r *http.Request) {
	taskId, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activitys, err := a.activityRepo.List(taskId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, activitys, http.StatusOK)
	return
}
