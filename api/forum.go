package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) forumRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateForum",
			"POST",
			"/forum",
			a.createForum,
		},
		{
			"DeleteForum",
			"DELETE",
			"/forum/{forumID:[0-9]+}",
			a.deleteForum,
		},
		{
			"EditForum",
			"PUT",
			"/forum/{forumID:[0-9]+}",
			a.editForum,
		},
		{
			"GetForum",
			"GET",
			"/forum/{forumID:[0-9]+}",
			a.getForum,
		},
		{
			"ListForum",
			"GET",
			"/forum",
			a.listForum,
		},
	}
	return
}

func (a *API) getForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		return
	}
	forum := &model.Forum{
		ID: forumID,
	}
	err = a.forumRepo.Get(forum, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = forum
	return
}

func (a *API) createForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	forum := &model.Forum{}

	err = decodeBody(r, forum)
	if err != nil {
		return
	}

	forum.OwnerID = user.ID
	err = a.forumRepo.Create(forum, user)
	if err != nil {
		return
	}
	content = forum
	return
}

func (a *API) deleteForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		return
	}
	forum := &model.Forum{
		ID: forumID,
	}
	err = a.forumRepo.Delete(forum, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = forum
	return
}

func (a *API) editForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		return
	}

	forum := &model.Forum{}
	err = decodeBody(r, forum)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	forum.ID = forumID

	err = a.forumRepo.Edit(forum, user)
	if err != nil {
		return
	}
	content = forum
	return
}

func (a *API) listForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	forums, err := a.forumRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = forums
	return
}
