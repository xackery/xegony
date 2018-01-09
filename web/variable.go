package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) variableRoutes() (routes []*route) {
	routes = []*route{

		//Variable
		{
			"ListVariable",
			"GET",
			"/variable",
			a.listVariable,
		},
		{
			"GetVariable",
			"GET",
			"/variable/{variableName}",
			a.getVariable,
		},
	}
	return
}

func (a *Web) listVariable(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		Variables []*model.Variable
	}

	site := a.newSite(r)
	site.Page = "variable"
	site.Title = "Variable"
	site.Section = "variable"

	variables, err := a.variableRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:      site,
		Variables: variables,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "variable/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("variable", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getVariable(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Variable *model.Variable
	}

	variableName := getVar(r, "variableName")

	variable, err := a.variableRepo.Get(variableName)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "variable"
	site.Title = "Variable"
	site.Section = "variable"

	content := Content{
		Site:     site,
		Variable: variable,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "variable/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("variable", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
