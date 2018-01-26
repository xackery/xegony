package web

import (
	"html/template"
	"net/http"
	"time"

	"github.com/xackery/xegony/model"
)

func (a *Web) getLogout(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Host string
	}

	cookie := &http.Cookie{
		Name:  "apiKey",
		Value: "",

		Expires: time.Now().Add(-10 * time.Minute),
	}

	site := a.newSite(r)
	site.Page = "logout"

	content = Content{
		Site: site,
	}
	http.SetCookie(w, cookie)
	site.User = nil
	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "logout.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("logout", tmp)
	}

	return
}
