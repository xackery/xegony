package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listSpell(w http.ResponseWriter, r *http.Request) {
	var err error

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

	spells, err := a.spellRepo.List(spellPage.PageSize, spellPage.PageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	spellPage.Total, err = a.spellRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:      site,
		Spells:    spells,
		SpellPage: spellPage,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spell/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("spell", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getSpell(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Spell *model.Spell
		Npcs  []*model.Npc
		Items []*model.Item
	}

	if strings.ToLower(getVar(r, "spellID")) == "search" {
		a.searchSpell(w, r)
		return
	}

	spellID, err := getIntVar(r, "spellID")
	if err != nil {
		err = errors.Wrap(err, "spellID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spell, err := a.spellRepo.Get(spellID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcs, err := a.npcRepo.ListBySpell(spellID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	items, err := a.itemRepo.ListBySpell(spellID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "spell"
	site.Title = "Spell"

	content := Content{
		Site:  site,
		Spell: spell,
		Items: items,
		Npcs:  npcs,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spell/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("spell", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) searchSpell(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Spells []*model.Spell
		Search string
	}

	search := getParam(r, "search")

	var spells []*model.Spell

	if len(search) > 0 {
		spells, err = a.spellRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "spell"
	site.Title = "Spell"

	content := Content{
		Site:   site,
		Spells: spells,
		Search: search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spell/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("spellsearch", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
