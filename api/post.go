package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetPost(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "postId")
	if err != nil {
		err = errors.Wrap(err, "postId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	post, err := a.postRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, post, http.StatusOK)
	return
}

func (a *Api) CreatePost(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	post := &model.Post{}
	err = decodeBody(r, post)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.postRepo.Create(post)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, post, http.StatusCreated)
	return
}

func (a *Api) DeletePost(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "postId")
	if err != nil {
		err = errors.Wrap(err, "postId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.postRepo.Delete(id)
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

func (a *Api) EditPost(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "postId")
	if err != nil {
		err = errors.Wrap(err, "postId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	post := &model.Post{}
	err = decodeBody(r, post)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.postRepo.Edit(id, post)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, post, http.StatusOK)
	return
}

func (a *Api) ListPost(w http.ResponseWriter, r *http.Request) {
	forumId, err := getIntVar(r, "forumId")
	if err != nil {
		err = errors.Wrap(err, "forumId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	posts, err := a.postRepo.List(forumId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, posts, http.StatusOK)
	return
}
