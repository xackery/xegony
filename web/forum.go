package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListForum(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   Site
		Forums []*model.Forum
	}

	site := a.NewSite(r)
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
		tmp, err = a.loadTemplate(nil, "body", "listforum.tpl")
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

func (a *Web) GetForum(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Forum *model.Forum
	}

	site := a.NewSite(r)
	site.Page = "forum"
	site.Title = "Forum"

	forumId, err := getIntVar(r, "forumId")
	if err != nil {
		err = errors.Wrap(err, "forumId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	forum, err := a.forumRepo.Get(forumId)
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
		tmp, err = a.loadTemplate(nil, "body", "getforum.tpl")
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
