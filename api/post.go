package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) postRoutes() (routes []*route) {
	routes = []*route{}
	return
}
func (a *API) getPost(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "postID")
	if err != nil {
		err = errors.Wrap(err, "postID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	post, err := a.postRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, post, http.StatusOK)
	return
}

func (a *API) createPost(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	post := &model.Post{}
	err = decodeBody(r, post)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.postRepo.Create(post)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, post, http.StatusCreated)
	return
}

func (a *API) deletePost(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "postID")
	if err != nil {
		err = errors.Wrap(err, "postID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.postRepo.Delete(id)
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

func (a *API) editPost(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "postID")
	if err != nil {
		err = errors.Wrap(err, "postID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	post := &model.Post{}
	err = decodeBody(r, post)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.postRepo.Edit(id, post)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, post, http.StatusOK)
	return
}

func (a *API) listPost(w http.ResponseWriter, r *http.Request) {
	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	posts, err := a.postRepo.List(forumID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, posts, http.StatusOK)
	return
}
