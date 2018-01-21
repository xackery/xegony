package web

import (
	"html/template"
	"net/http"

	"github.com/xackery/xegony/model"
)

func (a *Web) loginRoutes() (routes []*route) {
	routes = []*route{
		{
			"Login",
			"GET",
			"/login",
			a.getLogin,
		},
		{
			"Logout",
			"GET",
			"/logout",
			a.getLogout,
		},
	}
	return
}

func (a *Web) getLogin(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Host string
	}

	site := a.newSite(r)
	site.Page = "login"

	content = Content{
		Site: site,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "login.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("login", tmp)
	}

	return
}
