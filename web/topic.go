package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Topics []*model.Topic
		Forum  *model.Forum
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	if strings.ToLower(getVar(r, "forumID")) == "create" {
		a.createForum(w, r)
		return
	}

	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	topics, err := a.topicRepo.List(forumID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	forum, err := a.forumRepo.Get(forumID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:   site,
		Topics: topics,
		Forum:  forum,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "topic/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("topic", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Topic *model.Topic
	}

	id, err := getIntVar(r, "topicID")
	if err != nil {
		err = errors.Wrap(err, "topicID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	topic, err := a.topicRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	content := Content{
		Site:  site,
		Topic: topic,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "topic/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("topic", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
