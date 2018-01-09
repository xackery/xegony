package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) taskRoutes() (routes []*route) {
	routes = []*route{
		//Task
		{
			"GetTask",
			"GET",
			"/task/{taskID}/details",
			a.getTask,
		},
		{
			"ListTask",
			"GET",
			"/task",
			a.listTask,
		},
		{
			"ListTask",
			"POST",
			"/task",
			a.createTask,
		},
	}
	return
}

func (a *API) getTask(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	task, err := a.taskRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, task, http.StatusOK)
	return
}

func (a *API) createTask(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	task := &model.Task{}
	err = decodeBody(r, task)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.taskRepo.Create(task)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, task, http.StatusCreated)
	return
}

func (a *API) deleteTask(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.taskRepo.Delete(id)
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

func (a *API) editTask(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	task := &model.Task{}
	err = decodeBody(r, task)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.taskRepo.Edit(id, task)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, task, http.StatusOK)
	return
}

func (a *API) listTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := a.taskRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, tasks, http.StatusOK)
	return
}
