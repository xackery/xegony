package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/api"
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
			"/forum/{forumID}/details",
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

func (a *Web) listForum(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Forums []*model.Forum
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	forums, err := a.forumRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:   site,
		Forums: forums,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forum/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("forum", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getForum(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Forum *model.Forum
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

	forum, err := a.forumRepo.Get(forumID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:  site,
		Forum: forum,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forum/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("forum", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) createForum(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site site
	}

	site := a.newSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	if err = api.IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	content := Content{
		Site: site,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forum/create.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("forum", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
