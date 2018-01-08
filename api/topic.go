package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) topicRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateTopic",
			"POST",
			"/topic",
			a.createTopic,
		},
		{
			"DeleteTopic",
			"DELETE",
			"/topic/{topicID}",
			a.deleteTopic,
		},
		{
			"EditTopic",
			"PUT",
			"/topic/{topicID}",
			a.editTopic,
		},
		{
			"GetTopic",
			"GET",
			"/topic/{topicID}",
			a.getTopic,
		},
	}
	return
}

func (a *API) getTopic(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	topic, err := a.topicRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			a.writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	a.writeData(w, r, topic, http.StatusOK)
	return
}

func (a *API) createTopic(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	topic := &model.Topic{}
	err = decodeBody(r, topic)
	if err != nil {
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.topicRepo.Create(topic)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	a.writeData(w, r, topic, http.StatusCreated)
	return
}

func (a *API) deleteTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.topicRepo.Delete(id)
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

func (a *API) editTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	topic := &model.Topic{}
	err = decodeBody(r, topic)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.topicRepo.Edit(id, topic)
	if err != nil {
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, topic, http.StatusOK)
	return
}

func (a *API) listTopic(w http.ResponseWriter, r *http.Request) {
	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	topics, err := a.topicRepo.List(forumID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	a.writeData(w, r, topics, http.StatusOK)
	return
}
