package web

import (
	"html/template"
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

func (a *Web) listRule(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Rules []*model.Rule
	}

	site := a.newSite(r)
	site.Page = "rule"
	site.Title = "Rule"
	site.Section = "rule"

	rules, err := a.ruleRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:  site,
		Rules: rules,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "rule/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("rule", tmp)
	}

	return
}

func (a *Web) getRule(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Rule *model.Rule
	}

	ruleName := getVar(r, "ruleName")

	rule := &model.Rule{
		Name: ruleName,
	}
	err = a.ruleRepo.Get(rule, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "rule"
	site.Title = "Rule"
	site.Section = "rule"

	content = Content{
		Site: site,
		Rule: rule,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "rule/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("rule", tmp)
	}

	return
}
