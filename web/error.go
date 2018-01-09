package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) errorRoutes() (routes []*route) {
	routes = []*route{
		{
			"SearchError",
			"GET",
			"/error/search/{search}",
			a.searchError,
		},
		{
			"GetError",
			"GET",
			"/error/{errorID}",
			a.getError,
		},
		{
			"ListError",
			"GET",
			"/error",
			a.listError,
		},
	}
	return
}

func (a *Web) listError(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		Errors    []*model.Error
		ErrorPage *model.Page
	}

	site := a.newSite(r)
	site.Page = "error"
	site.Title = "Error"

	errorPage := &model.Page{
		Scope: "error",
	}

	errorPage.PageSize = getIntParam(r, "pageSize")
	errorPage.PageNumber = getIntParam(r, "pageNumber")

	errors, err := a.errorRepo.List(errorPage.PageSize, errorPage.PageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	errorPage.Total, err = a.errorRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:      site,
		Errors:    errors,
		ErrorPage: errorPage,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "error/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("error", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getError(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Error *model.Error
	}

	if strings.ToLower(getVar(r, "errorID")) == "search" {
		a.searchError(w, r)
		return
	}

	errorID, err := getIntVar(r, "errorID")
	if err != nil {
		err = errors.Wrap(err, "errorID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	errorStruct, err := a.errorRepo.Get(errorID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "error"
	site.Title = "Error"

	content := Content{
		Site:  site,
		Error: errorStruct,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "error/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("error", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) searchError(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Errors []*model.Error
		Search string
	}

	search := getParam(r, "search")

	var errors []*model.Error

	if len(search) > 0 {
		errors, err = a.errorRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "error"
	site.Title = "Error"

	content := Content{
		Site:   site,
		Errors: errors,
		Search: search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "error/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("errorsearch", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
