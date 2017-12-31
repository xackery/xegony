package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) getTopic(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "topicId")
	if err != nil {
		err = errors.Wrap(err, "topicId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	topic, err := a.topicRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, topic, http.StatusOK)
	return
}

func (a *Api) createTopic(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	topic := &model.Topic{}
	err = decodeBody(r, topic)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.topicRepo.Create(topic)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, topic, http.StatusCreated)
	return
}

func (a *Api) deleteTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "topicId")
	if err != nil {
		err = errors.Wrap(err, "topicId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.topicRepo.Delete(id)
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

func (a *Api) editTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "topicId")
	if err != nil {
		err = errors.Wrap(err, "topicId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	topic := &model.Topic{}
	err = decodeBody(r, topic)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.topicRepo.Edit(id, topic)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, topic, http.StatusOK)
	return
}

func (a *Api) listTopic(w http.ResponseWriter, r *http.Request) {
	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	topics, err := a.topicRepo.List(forumID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, topics, http.StatusOK)
	return
}
