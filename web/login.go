package web

import (
	"net/http"
)

func (a *Web) getLogin(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site site
		Host string
	}

	site := a.newSite(r)
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
