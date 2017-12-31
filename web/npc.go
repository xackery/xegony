package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listNpc(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site site
		Npcs []*model.Npc
	}

	site := a.newSite(r)
	site.Page = "npclist"
	site.Title = "Npc"
	site.Section = "npc"

	npcs, err := a.npcRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site: site,
		Npcs: npcs,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("npc", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) searchNpc(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Npcs   []*model.Npc
		Search string
	}

	search := getParam(r, "search")

	site := a.newSite(r)
	site.Page = "npcsearch"
	site.Title = "Npc"
	site.Section = "npc"
	var npcs []*model.Npc

	if len(search) > 0 {
		npcs, err = a.npcRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	content := Content{
		Site:   site,
		Npcs:   npcs,
		Search: search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("npcsearch", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listNpcByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	_, err = getIntVar(r, "zoneId")
	if err == nil {
		a.getNpcByZone(w, r)
		return
	}

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "listnpcbyzone"
	site.Title = "Npc"
	site.Section = "npc"

	zones, err := a.zoneRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:  site,
		Zones: zones,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/listbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("npc", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getNpcByZone(w http.ResponseWriter, r *http.Request) {
	var err error
	zoneId, err := getIntVar(r, "zoneId")
	if err != nil {
		err = errors.Wrap(err, "zoneId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site site
		Npcs []*model.Npc
	}

	site := a.newSite(r)
	site.Page = "listnpcbyzone"
	site.Title = "Npc"
	site.Section = "npc"

	npcs, err := a.npcRepo.ListByZone(zoneId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site: site,
		Npcs: npcs,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/getbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("npc", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listNpcByFaction(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Factions []*model.Faction
	}

	site := a.newSite(r)
	site.Page = "listnpcbyfaction"
	site.Title = "Npc"
	site.Section = "npc"

	factions, err := a.factionRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:     site,
		Factions: factions,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/listbyfaction.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("npc", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getNpcByFaction(w http.ResponseWriter, r *http.Request) {
	var err error
	factionId, err := getIntVar(r, "factionId")
	if err != nil {
		err = errors.Wrap(err, "factionId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site    site
		Npcs    []*model.Npc
		Faction *model.Faction
	}

	site := a.newSite(r)
	site.Page = "listnpcbyfaction"
	site.Title = "Npc"
	site.Section = "npc"

	npcs, err := a.npcRepo.ListByFaction(factionId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	faction, err := a.factionRepo.Get(factionId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:    site,
		Npcs:    npcs,
		Faction: faction,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/getbyfaction.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("npc", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getNpc(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Npc   *model.Npc
		Items []*model.NpcLoot
	}

	if strings.ToLower(getVar(r, "npcId")) == "search" {
		a.searchNpc(w, r)
		return
	}
	if strings.ToLower(getVar(r, "npcId")) == "byzone" {
		a.listNpcByZone(w, r)
		return
	}

	if strings.ToLower(getVar(r, "npcId")) == "byfaction" {
		a.listNpcByFaction(w, r)
		return
	}

	npcId, err := getIntVar(r, "npcId")
	if err != nil {
		err = errors.Wrap(err, "npcId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	npc, err := a.npcRepo.Get(npcId)
	if err != nil {
		err = errors.Wrap(err, "Request error on npc")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	items, err := a.npcLootRepo.List(npcId)
	if err != nil {
		err = errors.Wrap(err, "Request error on items")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "npc"
	site.Title = "Npc"

	content := Content{
		Site:  site,
		Npc:   npc,
		Items: items,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("npc", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
