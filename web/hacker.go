package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listHacker(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		Hackers    []*model.Hacker
		HackerPage *model.Page
	}

	site := a.newSite(r)
	site.Page = "hacker"
	site.Title = "Hacker"

	hackerPage := &model.Page{
		Scope: "hacker",
	}

	hackerPage.PageSize = getIntParam(r, "hackerPage.PageSize")
	hackerPage.PageNumber = getIntParam(r, "hackerPage.PageNumber")

	hackers, err := a.hackerRepo.List(hackerPage.PageSize, hackerPage.PageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	hackerPage.Total, err = a.hackerRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, hacker := range hackers {
		if len(hacker.ZoneName.String) > 0 {
			hacker.Zone, err = a.zoneRepo.GetByShortname(hacker.ZoneName.String)
			if err != nil {
				err = errors.Wrap(err, "invalid zone load")
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
		if len(hacker.AccountName) > 0 {
			hacker.Account, err = a.accountRepo.GetByName(hacker.AccountName)
			if err != nil {
				err = errors.Wrap(err, "invalid account name")
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
		if len(hacker.CharacterName) > 0 {
			hacker.Character, err = a.characterRepo.GetByName(hacker.CharacterName)
			if err != nil {
				err = errors.Wrap(err, "invalid account name")
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
	}

	content := Content{
		Site:       site,
		Hackers:    hackers,
		HackerPage: hackerPage,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "hacker/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("hacker", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getHacker(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Hacker *model.Hacker
	}

	if strings.ToLower(getVar(r, "hackerID")) == "search" {
		a.searchHacker(w, r)
		return
	}

	hackerID, err := getIntVar(r, "hackerID")
	if err != nil {
		err = errors.Wrap(err, "hackerID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	hacker, err := a.hackerRepo.Get(hackerID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	if len(hacker.ZoneName.String) > 0 {
		hacker.Zone, err = a.zoneRepo.GetByShortname(hacker.ZoneName.String)
		if err != nil {
			err = errors.Wrap(err, "invalid zone load")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	site := a.newSite(r)
	site.Page = "hacker"
	site.Title = "Hacker"

	content := Content{
		Site:   site,
		Hacker: hacker,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "hacker/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("hacker", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) searchHacker(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site    site
		Hackers []*model.Hacker
		Search  string
	}

	search := getParam(r, "search")

	var hackers []*model.Hacker

	if len(search) > 0 {
		hackers, err = a.hackerRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "hacker"
	site.Title = "Hacker"

	content := Content{
		Site:    site,
		Hackers: hackers,
		Search:  search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "hacker/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("hackersearch", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
