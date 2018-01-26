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
			"/topic/{topicID:[0-9]+}",
			a.deleteTopic,
		},
		{
			"EditTopic",
			"PUT",
			"/topic/{topicID:[0-9]+}",
			a.editTopic,
		},
		{
			"GetTopic",
			"GET",
			"/topic/{topicID:[0-9]+}",
			a.getTopic,
		},
	}
	return
}

func (a *API) getTopic(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	topicID, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		return
	}
	topic := &model.Topic{
		ID: topicID,
	}
	err = a.topicRepo.Get(topic, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = topic
	return
}

func (a *API) createTopic(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	topic := &model.Topic{}
	err = decodeBody(r, topic)
	if err != nil {
		return
	}
	err = a.topicRepo.Create(topic, user)
	if err != nil {
		return
	}
	content = topic
	return
}

func (a *API) deleteTopic(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	topicID, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		return
	}

	topic := &model.Topic{
		ID: topicID,
	}

	err = a.topicRepo.Delete(topic, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = topic
	return
}

func (a *API) editTopic(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	topicID, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		return
	}

	topic := &model.Topic{}
	err = decodeBody(r, topic)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	topic.ID = topicID
	err = a.topicRepo.Edit(topic, user)
	if err != nil {
		return
	}
	content = topic
	return
}

func (a *API) listTopic(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		return
	}
	forum := &model.Forum{
		ID: forumID,
	}
	topics, err := a.topicRepo.ListByForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = topics
	return
}
