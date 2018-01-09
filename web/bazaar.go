package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) bazaarRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *Web) listBazaar(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site    site
		Bazaars []*model.Bazaar
	}

	site := a.newSite(r)
	site.Page = "bazaar"
	site.Title = "Bazaar"

	bazaars, err := a.bazaarRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:    site,
		Bazaars: bazaars,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "bazaar/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("bazaar", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getBazaar(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Bazaar *model.Bazaar
	}

	id, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	bazaar, err := a.bazaarRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "bazaar"
	site.Title = "Bazaar"

	content := Content{
		Site:   site,
		Bazaar: bazaar,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "bazaar/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("bazaar", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
