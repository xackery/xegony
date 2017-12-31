package web

import (
	"net/http"
	"time"
)

func (a *Web) getLogout(w http.ResponseWriter, r *http.Request) {
	var err error

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

	content := Content{
		Site: site,
	}
	http.SetCookie(w, cookie)
	site.User = nil
	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "logout.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("logout", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
