package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

// swagger:parameters deleteActivity editActivity getActivity
type ActivityParams struct {
	//ActivityID to get information about
	// in: path
	ActivityID int64 `json:"activityID"`
	//todo: pagination
}

func (a *API) activityRoutes() (routes []*route) {
	routes = []*route{
		// swagger:route GET /activity activity listActivity
		//
		// Lists activities
		//
		// This will show all available activities by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: Activities
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListActivity",
			"GET",
			"/task/{taskID:[0-9]+}",
			a.listActivity,
		},
		// swagger:route POST /activity activity createActivity
		//
		// Create an activity
		//
		// This will create an activity
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateActivity",
			"POST",
			"/task/{taskID:[0-9]+}/activity",
			a.createActivity,
		},
		// swagger:route GET /activity/{activityID} activity getActivity
		//
		// Get an activity
		//
		// This will get an individual activity available activitys by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: Activity
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetActivity",
			"GET",
			"/task/{taskID:[0-9]+}/activity/{activityID:[0-9]+}",
			a.getActivity,
		},
		// swagger:route PUT /activity/{activityID} activity editActivity
		//
		// Edit an activity
		//
		// This will edit an activity
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ErrNoContent
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditActivity",
			"PUT",
			"task/{taskID:[0-9]+}/activity/{activityID:[0-9]+}",
			a.editActivity,
		},
		// swagger:route DELETE /activity/{activityID} activity deleteActivity
		//
		// Delete an activity
		//
		// This will delete an activity
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"DeleteActivity",
			"DELETE",
			"task/{taskID:[0-9]+}/activity/{activityID:[0-9]+}",
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
