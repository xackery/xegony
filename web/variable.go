package web

import (
	"html/template"
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
			"/variable/{name}",
			a.getVariable,
		},
	}
	return
}

func (a *Web) listVariable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		Variables []*model.Variable
	}

	site := a.newSite(r)
	site.Page = "variable"
	site.Title = "Variable"
	site.Section = "variable"

	variables, err := a.variableRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:      site,
		Variables: variables,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "variable/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("variable", tmp)
	}

	return
}

func (a *Web) getVariable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		Variable *model.Variable
	}

	name := getVar(r, "name")

	variable := &model.Variable{
		Name: name,
	}
	err = a.variableRepo.GetByName(variable, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "variable"
	site.Title = "Variable"
	site.Section = "variable"

	content = Content{
		Site:     site,
		Variable: variable,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "variable/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("variable", tmp)
	}

	return
}
