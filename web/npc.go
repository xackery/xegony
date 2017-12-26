package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListNpc(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Npcs []*model.Npc
	}

	site := a.NewSite(r)
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
		tmp, err = a.loadTemplate(nil, "body", "listnpc.tpl")
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

func (a *Web) ListNpcByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	_, err = getIntVar(r, "zoneId")
	if err == nil {
		a.GetNpcByZone(w, r)
		return
	}

	type Content struct {
		Site  Site
		Zones []*model.Zone
	}

	site := a.NewSite(r)
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
		tmp, err = a.loadTemplate(nil, "body", "listnpcbyzone.tpl")
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

func (a *Web) GetNpcByZone(w http.ResponseWriter, r *http.Request) {
	var err error
	zoneId, err := getIntVar(r, "zoneId")
	if err != nil {
		err = errors.Wrap(err, "zoneId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site Site
		Npcs []*model.Npc
	}

	site := a.NewSite(r)
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
		tmp, err = a.loadTemplate(nil, "body", "getnpcbyzone.tpl")
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

func (a *Web) ListNpcByFaction(w http.ResponseWriter, r *http.Request) {
	var err error

	_, err = getIntVar(r, "factionId")
	if err == nil {
		a.GetNpcByFaction(w, r)
		return
	}

	type Content struct {
		Site     Site
		Factions []*model.Faction
	}

	site := a.NewSite(r)
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
		tmp, err = a.loadTemplate(nil, "body", "listnpcbyfaction.tpl")
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

func (a *Web) GetNpcByFaction(w http.ResponseWriter, r *http.Request) {
	var err error
	factionId, err := getIntVar(r, "factionId")
	if err != nil {
		err = errors.Wrap(err, "factionId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site Site
		Npcs []*model.Npc
	}

	site := a.NewSite(r)
	site.Page = "listnpcbyfaction"
	site.Title = "Npc"
	site.Section = "npc"

	npcs, err := a.npcRepo.ListByFaction(factionId)
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
		tmp, err = a.loadTemplate(nil, "body", "getnpcbyfaction.tpl")
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

func (a *Web) GetNpc(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Npc  *model.Npc
	}

	if strings.ToLower(getVar(r, "npcId")) == "byzone" {
		a.ListNpcByZone(w, r)
		return
	}

	if strings.ToLower(getVar(r, "npcId")) == "byfaction" {
		a.ListNpcByFaction(w, r)
		return
	}

	id, err := getIntVar(r, "npcId")
	if err != nil {
		err = errors.Wrap(err, "npcId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	npc, err := a.npcRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "npc"
	site.Title = "Npc"

	content := Content{
		Site: site,
		Npc:  npc,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "getnpc.tpl")
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
