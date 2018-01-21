package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) topicRoutes() (routes []*route) {
	routes = []*route{
		//Topic
		{
			"ListTopic",
			"GET",
			"/forum/{forumID:[0-9]+}",
			a.listTopic,
		},
		{
			"GetTopic",
			"GET",
			"/topic/{topicID:[0-9]+}/details",
			a.getTopic,
		},
	}
	return
}

func (a *Web) listTopic(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Topics []*model.Topic
		Forum  *model.Forum
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

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
		return
	}

	topics, err := a.topicRepo.ListByForum(forum, user)
	if err != nil {
		return
	}
	content = Content{
		Site:   site,
		Topics: topics,
		Forum:  forum,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "topic/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("topic", tmp)
	}

	return
}

func (a *Web) getTopic(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Topic *model.Topic
	}

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
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	content = Content{
		Site:  site,
		Topic: topic,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "topic/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("topic", tmp)
	}

	return
}
