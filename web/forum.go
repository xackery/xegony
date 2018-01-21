package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) forumRoutes() (routes []*route) {
	routes = []*route{
		//Forum
		{
			"ListForum",
			"GET",
			"/forum",
			a.listForum,
		},
		{
			"GetForum",
			"GET",
			"/forum/{forumID:[0-9]+}/details",
			a.getForum,
		},
		{
			"CreateForum",
			"GET",
			"/forum/create",
			a.createForum,
		},
	}
	return
}

func (a *Web) listForum(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Forums []*model.Forum
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	forums, err := a.forumRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:   site,
		Forums: forums,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forum/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("forum", tmp)
	}

	return
}

func (a *Web) getForum(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Forum *model.Forum
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
	content = Content{
		Site:  site,
		Forum: forum,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forum/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("forum", tmp)
	}

	return
}

func (a *Web) createForum(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	content = Content{
		Site: site,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forum/create.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("forum", tmp)
	}

	return
}
