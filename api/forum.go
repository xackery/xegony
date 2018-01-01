package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getForum(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	forum, err := a.forumRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, forum, http.StatusOK)
	return
}

func (a *API) createForum(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	claims, err := GetAuthClaims(r)
	if err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
	}

	forum := &model.Forum{}

	err = decodeBody(r, forum)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	forum.OwnerID = claims.User.ID
	err = a.forumRepo.Create(forum)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, forum, http.StatusCreated)
	return
}

func (a *API) deleteForum(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.forumRepo.Delete(id)
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

func (a *API) editForum(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	forum := &model.Forum{}
	err = decodeBody(r, forum)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.forumRepo.Edit(id, forum)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, forum, http.StatusOK)
	return
}

func (a *API) listForum(w http.ResponseWriter, r *http.Request) {
	forums, err := a.forumRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, forums, http.StatusOK)
	return
}
