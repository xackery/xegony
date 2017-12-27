package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetTask(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
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

func (a *Api) CreateTask(w http.ResponseWriter, r *http.Request) {
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

func (a *Api) DeleteTask(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
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

func (a *Api) EditTask(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
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

func (a *Api) ListTask(w http.ResponseWriter, r *http.Request) {
	tasks, err := a.taskRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, tasks, http.StatusOK)
	return
}
