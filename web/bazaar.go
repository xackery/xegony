package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) bazaarRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *Web) listBazaar(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site    site
		Bazaars []*model.Bazaar
	}

	site := a.newSite(r)
	site.Page = "bazaar"
	site.Title = "Bazaar"

	bazaars, err := a.bazaarRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:    site,
		Bazaars: bazaars,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "bazaar/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("bazaar", tmp)
	}

	return
}

func (a *Web) getBazaar(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Bazaar *model.Bazaar
	}

	bazaarID, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		return
	}
	bazaar := &model.Bazaar{
		ID: bazaarID,
	}
	err = a.bazaarRepo.Get(bazaar, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "bazaar"
	site.Title = "Bazaar"

	content = Content{
		Site:   site,
		Bazaar: bazaar,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "bazaar/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("bazaar", tmp)
	}

	return
}
