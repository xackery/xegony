package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) errorRoutes() (routes []*route) {
	routes = []*route{
		{
			"SearchError",
			"GET",
			"/error/{search}",
			a.searchError,
		},
		{
			"GetError",
			"GET",
			"/error/{errorID:[0-9]+}",
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

func (a *Web) listError(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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

	errors, err := a.errorRepo.List(errorPage.PageSize, errorPage.PageNumber, user)
	if err != nil {
		return
	}
	errorPage.Total, err = a.errorRepo.ListCount(user)
	if err != nil {
		return
	}
	content = Content{
		Site:      site,
		Errors:    errors,
		ErrorPage: errorPage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "error/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("error", tmp)
	}

	return
}

func (a *Web) getError(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Error *model.Error
	}

	errorID, err := getIntVar(r, "errorID")
	if err != nil {
		err = errors.Wrap(err, "errorID argument is required")
		return
	}
	errorStruct := &model.Error{
		ID: errorID,
	}

	err = a.errorRepo.Get(errorStruct, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "error"
	site.Title = "Error"

	content = Content{
		Site:  site,
		Error: errorStruct,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "error/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("error", tmp)
	}

	return
}

func (a *Web) searchError(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Errors []*model.Error
		Search string
	}

	search := getParam(r, "search")

	var errors []*model.Error

	if len(search) > 0 {
		errorStruct := &model.Error{
			Message: search,
		}
		errors, err = a.errorRepo.SearchByMessage(errorStruct, user)
		if err != nil {
			return
		}
	}

	site := a.newSite(r)
	site.Page = "error"
	site.Title = "Error"

	content = Content{
		Site:   site,
		Errors: errors,
		Search: search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "error/search.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("errorsearch", tmp)
	}

	return
}
