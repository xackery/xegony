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
			"/forum/{forumID}",
			a.deleteForum,
		},
		{
			"EditForum",
			"PUT",
			"/forum/{forumID}",
			a.editForum,
		},
		{
			"GetForum",
			"GET",
			"/forum/{forumID}",
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

func (a *API) getForum(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	forum, err := a.forumRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, forum, http.StatusOK)
	return
}

func (a *API) createForum(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	claims, err := GetAuthClaims(r)
	if err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
	}

	forum := &model.Forum{}

	err = decodeBody(r, forum)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	forum.OwnerID = claims.User.ID
	err = a.forumRepo.Create(forum)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, forum, http.StatusCreated)
	return
}

func (a *API) deleteForum(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.forumRepo.Delete(id)
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

func (a *API) editForum(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	forum := &model.Forum{}
	err = decodeBody(r, forum)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.forumRepo.Edit(id, forum)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, forum, http.StatusOK)
	return
}

func (a *API) listForum(w http.ResponseWriter, r *http.Request) {
	forums, err := a.forumRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, forums, http.StatusOK)
	return
}
