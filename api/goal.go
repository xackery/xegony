package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) goalRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateGoal",
			"POST",
			"/goal",
			a.createGoal,
		},
		{
			"DeleteAccount",
			"DELETE",
			"/goal/{goalID:[0-9]+}",
			a.deleteGoal,
		},
		{
			"EditGoal",
			"PUT",
			"/goal/{goalID:[0-9]+}",
			a.editGoal,
		},
		{
			"GetAccount",
			"GET",
			"/goal/{goalID:[0-9]+}",
			a.getGoal,
		},
	}
	return
}

func (a *API) getGoal(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	listID, err := getIntVar(r, "listID")
	if err != nil {
		err = errors.Wrap(err, "listID argument is required")
		return
	}

	entryID, err := getIntVar(r, "entryID")
	if err != nil {
		err = errors.Wrap(err, "entryID argument is required")
		return
	}
	goal := &model.Goal{
		EntryID: entryID,
		ListID:  listID,
	}
	err = a.goalRepo.Get(goal, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = goal
	return
}

func (a *API) createGoal(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	goal := &model.Goal{}
	err = decodeBody(r, goal)
	if err != nil {
		return
	}
	err = a.goalRepo.Create(goal, user)
	if err != nil {
		return
	}
	content = goal
	return
}

func (a *API) deleteGoal(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	listID, err := getIntVar(r, "listID")
	if err != nil {
		err = errors.Wrap(err, "listID argument is required")
		return
	}

	entryID, err := getIntVar(r, "entryID")
	if err != nil {
		err = errors.Wrap(err, "entryID argument is required")
		return
	}
	goal := &model.Goal{
		EntryID: entryID,
		ListID:  listID,
	}

	err = a.goalRepo.Delete(goal, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = goal
	return
}

func (a *API) editGoal(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	listID, err := getIntVar(r, "listID")
	if err != nil {
		err = errors.Wrap(err, "listID argument is required")
		return
	}

	goal := &model.Goal{}
	err = decodeBody(r, goal)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	goal.ListID = listID

	err = a.goalRepo.Edit(goal, user)
	if err != nil {
		return
	}
	content = goal
	return
}

func (a *API) listGoal(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	goals, err := a.goalRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = goals
	return
}
