package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) activityRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListActivity",
			"GET",
			"/task/{taskID:[0-9]+}/activity",
			a.listActivity,
		},
		{
			"GetActivity",
			"GET",
			"/task/{taskID:[0-9]+}/activity/{activityID:[0-9]+}",
			a.getActivity,
		},
		{
			"CreateActivity",
			"POST",
			"/task/{taskID:[0-9]+}/activity",
			a.createActivity,
		},
		{
			"DeleteActivity",
			"DELETE",
			"/task/{taskID:[0-9]+}/activity/{activityID:[0-9]+}",
			a.deleteActivity,
		},
	}
	return
}

func (a *API) getActivity(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	activityID, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		return
	}

	activity := &model.Activity{
		ActivityID: activityID,
	}
	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}
	activity.TaskID = taskID
	err = a.activityRepo.Get(activity, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = activity
	return
}

func (a *API) createActivity(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	activity := &model.Activity{}
	err = decodeBody(r, activity)
	if err != nil {
		return
	}
	err = a.activityRepo.Create(activity, user)
	if err != nil {
		return
	}
	content = activity
	return
}

func (a *API) deleteActivity(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	activityID, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		return
	}

	activity := &model.Activity{
		ActivityID: activityID,
	}
	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}
	activity.TaskID = taskID

	err = a.activityRepo.Delete(activity, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = activity
	return
}

func (a *API) editActivity(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	activityID, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		return
	}

	activity := &model.Activity{}
	err = decodeBody(r, activity)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	activity.ActivityID = activityID
	err = a.activityRepo.Edit(activity, user)
	if err != nil {
		return
	}
	content = activity
	return
}

func (a *API) listActivity(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}
	task := &model.Task{
		ID: taskID,
	}
	activitys, err := a.activityRepo.ListByTask(task, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = activitys
	return
}
