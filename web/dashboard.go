package web

import (
	"html/template"
	"net/http"

	"github.com/xackery/xegony/model"
)

func (a *Web) dashboardRoutes() (routes []*route) {
	routes = []*route{
		{
			"GetDashboard",
			"GET",
			"/dashboard",
			a.getDashboard,
		},
	}
	return
}

func (a *Web) getDashboard(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Host string
	}

	site := a.newSite(r)
	site.Page = "dashboard"

	content = Content{
		Site: site,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "dashboard.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("dashboard", tmp)
	}

	return
}
