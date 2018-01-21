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
			"/task/{taskID:[0-9]+}",
			a.getTask,
		},
		{
			"ListTask",
			"GET",
			"/task",
			a.listTask,
		},
		{
			"PostTask",
			"POST",
			"/task",
			a.createTask,
		},
	}
	return
}

func (a *API) getTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}
	task := &model.Task{
		ID: taskID,
	}
	err = a.taskRepo.Get(task, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = task
	return
}

func (a *API) createTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	task := &model.Task{}
	err = decodeBody(r, task)
	if err != nil {
		return
	}
	err = a.taskRepo.Create(task, user)
	if err != nil {
		return
	}
	content = task
	return
}

func (a *API) deleteTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}

	task := &model.Task{
		ID: taskID,
	}
	err = a.taskRepo.Delete(task, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = task
	return
}

func (a *API) editTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}

	task := &model.Task{}
	err = decodeBody(r, task)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	task.ID = taskID
	err = a.taskRepo.Edit(task, user)
	if err != nil {
		return
	}
	content = task
	return
}

func (a *API) listTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	tasks, err := a.taskRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = tasks
	return
}
