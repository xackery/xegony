package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) spellRoutes() (routes []*route) {
	routes = []*route{

		//Spell
		{
			"SearchSpell",
			"GET",
			"/spell/search/{search:[0-9]+}",
			a.searchSpell,
		},
		{
			"GetSpell",
			"GET",
			"/spell/{spellID:[0-9]+}",
			a.getSpell,
		},
		{
			"ListSpell",
			"GET",
			"/spell",
			a.listSpell,
		},
	}
	return
}

func (a *Web) listSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		Spells    []*model.Spell
		SpellPage *model.Page
	}

	site := a.newSite(r)
	site.Page = "spell"
	site.Title = "Spell"

	spellPage := &model.Page{
		Scope: "spell",
	}

	spellPage.PageSize = getIntParam(r, "spellPage.PageSize")
	spellPage.PageNumber = getIntParam(r, "spellPage.PageNumber")

	spells, err := a.spellRepo.List(spellPage.PageSize, spellPage.PageNumber, user)
	if err != nil {
		return
	}
	for _, spell := range spells {
		spell.Skill = &model.Skill{
			ID: spell.SkillID,
		}
		err = a.skillRepo.Get(spell.Skill, user)
		if err != nil {
			return
		}
	}
	spellPage.Total, err = a.spellRepo.ListCount(user)
	if err != nil {
		return
	}
	content = Content{
		Site:      site,
		Spells:    spells,
		SpellPage: spellPage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spell/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spell", tmp)
	}

	return
}

func (a *Web) getSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Spell *model.Spell
		Npcs  []*model.Npc
		Items []*model.Item
	}

	spellID, err := getIntVar(r, "spellID")
	if err != nil {
		err = errors.Wrap(err, "spellID argument is required")
		return
	}

	spell := &model.Spell{
		ID: spellID,
	}
	err = a.spellRepo.Get(spell, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	npcs, err := a.npcRepo.ListBySpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	items, err := a.itemRepo.ListBySpell(spell, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "spell"
	site.Title = "Spell"

	content = Content{
		Site:  site,
		Spell: spell,
		Items: items,
		Npcs:  npcs,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spell/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spell", tmp)
	}

	return
}

func (a *Web) searchSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Spells []*model.Spell
		Search string
	}

	search := getParam(r, "search")

	var spells []*model.Spell

	if len(search) > 0 {
		spell := &model.Spell{}
		spell.Name.String = search
		spells, err = a.spellRepo.Search(spell, user)
		if err != nil {
			return
		}
	}

	site := a.newSite(r)
	site.Page = "spell"
	site.Title = "Spell"

	content = Content{
		Site:   site,
		Spells: spells,
		Search: search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spell/search.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spellsearch", tmp)
	}

	return
}
