package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func forumRoutes() (routes []*route) {
	routes = []*route{
		//Forum
		{
			"ListForum",
			"GET",
			"/forum",
			listForum,
		},
		{
			"GetForum",
			"GET",
			"/forum/{forumID:[0-9]+}",
			getForum,
		},
	}
	return
}

func listForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Forums []*model.Forum
	}

	site := newSite(r)
	page := &model.Page{}
	forums, err := cases.ListForum(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:   site,
		Forums: forums,
	}

	tmp, err = loadTemplate(nil, "body", "forum/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Forum *model.Forum
	}

	forumID, err := getIntVar(r, "forumID")
	if err != nil {
		err = errors.Wrap(err, "forumID argument is required")
		return
	}
	forum := &model.Forum{
		ID: forumID,
	}

	err = cases.GetForum(forum, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := newSite(r)
	site.Page = "forum"
	site.Title = "Forum"
	site.Section = "forum"

	content = Content{
		Site:  site,
		Forum: forum,
	}

	tmp, err = loadTemplate(nil, "body", "forum/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
