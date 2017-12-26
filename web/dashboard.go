package web

import (
	"net/http"
)

func (a *Web) GetDashboard(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Host string
	}

	site := a.NewSite(r)
	site.Page = "dashboard"

	content := Content{
		Site: site,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "dashboard.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("dashboard", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
