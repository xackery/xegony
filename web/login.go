package web

import (
	"net/http"
)

func (a *Web) GetLogin(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Host string
	}

	site := a.NewSite(r)
	site.Page = "login"

	content := Content{
		Site: site,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "login.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("login", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
