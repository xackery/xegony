package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getTask(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	task, err := a.taskRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, task, http.StatusOK)
	return
}

func (a *API) createTask(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	task := &model.Task{}
	err = decodeBody(r, task)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.taskRepo.Create(task)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, task, http.StatusCreated)
	return
}

func (a *API) deleteTask(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.taskRepo.Delete(id)
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

func (a *API) editTask(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	task := &model.Task{}
	err = decodeBody(r, task)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.taskRepo.Edit(id, task)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, task, http.StatusOK)
	return
}

func (a *API) listTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := a.taskRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, tasks, http.StatusOK)
	return
}
