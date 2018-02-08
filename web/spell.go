package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func spellRoutes() (routes []*route) {
	routes = []*route{
		//Spell
		{
			"ListSpell",
			"GET",
			"/spell",
			listSpell,
		},
		{
			"GetSpell",
			"GET",
			"/spell/{spellID:[0-9]+}",
			getSpell,
		},
	}
	return
}

func listSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Spells []*model.Spell
		Page   *model.Page
	}

	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}
	spells, err := cases.ListSpell(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:   site,
		Spells: spells,
		Page:   page,
	}

	tmp, err = loadTemplate(nil, "body", "spell/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Spell *model.Spell
	}

	spellID := getIntVar(r, "spellID")
	if err != nil {
		err = errors.Wrap(err, "spellID argument is required")
		return
	}
	spell := &model.Spell{
		ID: spellID,
	}

	err = cases.GetSpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := newSite(r, user)
	site.Page = "spell"
	site.Title = "Spell"
	site.Section = "spell"

	content = Content{
		Site:  site,
		Spell: spell,
	}

	tmp, err = loadTemplate(nil, "body", "spell/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
