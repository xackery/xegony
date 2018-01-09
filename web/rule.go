package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ruleRoutes() (routes []*route) {
	routes = []*route{
		//Rule
		{
			"ListRule",
			"GET",
			"/rule",
			a.listRule,
		},
		{
			"GetRule",
			"GET",
			"/rule/{ruleName}",
			a.getRule,
		},
	}
	return
}

func (a *Web) listRule(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Rules []*model.Rule
	}

	site := a.newSite(r)
	site.Page = "rule"
	site.Title = "Rule"
	site.Section = "rule"

	rules, err := a.ruleRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:  site,
		Rules: rules,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "rule/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("rule", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getRule(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site site
		Rule *model.Rule
	}

	ruleName := getVar(r, "ruleName")

	rule, err := a.ruleRepo.Get(ruleName)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "rule"
	site.Title = "Rule"
	site.Section = "rule"

	content := Content{
		Site: site,
		Rule: rule,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "rule/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("rule", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
