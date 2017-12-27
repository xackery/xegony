package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   Site
		Topics []*model.Topic
		Forum  *model.Forum
	}

	site := a.NewSite(r)
	site.Page = "topic"
	site.Title = "Topic"

	forumId, err := getIntVar(r, "forumId")
	if err != nil {
		err = errors.Wrap(err, "forumId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	topics, err := a.topicRepo.List(forumId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	forum, err := a.forumRepo.Get(forumId)
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
		tmp, err = a.loadTemplate(nil, "body", "listtopic.tpl")
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

func (a *Web) GetTopic(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Topic *model.Topic
	}

	id, err := getIntVar(r, "topicId")
	if err != nil {
		err = errors.Wrap(err, "topicId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	topic, err := a.topicRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "topic"
	site.Title = "Topic"

	content := Content{
		Site:  site,
		Topic: topic,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "gettopic.tpl")
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
