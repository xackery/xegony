package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) hackerRoutes() (routes []*route) {
	routes = []*route{
		//Hacker
		{
			"SearchHacker",
			"GET",
			"/hacker/{search:[a-Z]+}",
			a.searchHacker,
		},
		{
			"GetHacker",
			"GET",
			"/hacker/{hackerID:[0-9]+}",
			a.getHacker,
		},
		{
			"ListHacker",
			"GET",
			"/hacker",
			a.listHacker,
		},
	}
	return
}

func (a *Web) listHacker(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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

	hackers, err := a.hackerRepo.List(hackerPage.PageSize, hackerPage.PageNumber, user)
	if err != nil {
		return
	}
	hackerPage.Total, err = a.hackerRepo.ListCount(user)
	if err != nil {
		return
	}

	for _, hacker := range hackers {
		if len(hacker.ZoneName.String) > 0 {
			zone := &model.Zone{
				ShortName: hacker.ZoneName,
			}
			err = a.zoneRepo.GetByShortname(zone, user)
			if err != nil {
				err = errors.Wrap(err, "invalid zone load")
				return
			}
		}
		if len(hacker.AccountName) > 0 {
			account := &model.Account{
				Name: hacker.AccountName,
			}
			err = a.accountRepo.GetByName(account, user)
			if err != nil {
				err = errors.Wrap(err, "invalid account name")
				return
			}
		}
		if len(hacker.CharacterName) > 0 {
			character := &model.Character{
				Name: hacker.CharacterName,
			}
			err = a.characterRepo.GetByName(character, user)
			if err != nil {
				err = errors.Wrap(err, "invalid account name")
				return
			}
		}
	}

	content = Content{
		Site:       site,
		Hackers:    hackers,
		HackerPage: hackerPage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "hacker/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("hacker", tmp)
	}

	return
}

func (a *Web) getHacker(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Hacker *model.Hacker
	}

	hackerID, err := getIntVar(r, "hackerID")
	if err != nil {
		err = errors.Wrap(err, "hackerID argument is required")
		return
	}

	hacker := &model.Hacker{
		ID: hackerID,
	}

	err = a.hackerRepo.Get(hacker, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	if len(hacker.ZoneName.String) > 0 {
		zone := &model.Zone{
			ShortName: hacker.ZoneName,
		}
		err = a.zoneRepo.GetByShortname(zone, user)
		if err != nil {
			err = errors.Wrap(err, "invalid zone load")
			return
		}
	}
	site := a.newSite(r)
	site.Page = "hacker"
	site.Title = "Hacker"

	content = Content{
		Site:   site,
		Hacker: hacker,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "hacker/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("hacker", tmp)
	}

	return
}

func (a *Web) searchHacker(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site    site
		Hackers []*model.Hacker
		Search  string
	}

	search := getParam(r, "search")

	var hackers []*model.Hacker

	if len(search) > 0 {
		hacker := &model.Hacker{
			Hacked: search,
		}
		hackers, err = a.hackerRepo.SearchByMessage(hacker, user)
		if err != nil {
			return
		}
	}

	site := a.newSite(r)
	site.Page = "hacker"
	site.Title = "Hacker"

	content = Content{
		Site:    site,
		Hackers: hackers,
		Search:  search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "hacker/search.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("hackersearch", tmp)
	}

	return
}
