package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) npcRoutes() (routes []*route) {
	routes = []*route{
		//Npc
		{
			"GetNpc",
			"GET",
			"/npc/{npcID}",
			a.getNpc,
		},
		{
			"ListNpc",
			"GET",
			"/npc",
			a.listNpc,
		},
		{
			"ListNpcByZone",
			"GET",
			"/npc/byzone",
			a.listNpcByZone,
		},
		{
			"GetNpcByZone",
			"GET",
			"/npc/byzone/{zoneID}",
			a.getNpcByZone,
		},
		{
			"ListNpcByFaction",
			"GET",
			"/npc/byfaction",
			a.listNpcByFaction,
		},
		{
			"GetNpcByFaction",
			"GET",
			"/npc/byfaction/{factionID}",
			a.getNpcByFaction,
		},
	}
	return
}

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

	_, err = getIntVar(r, "zoneID")
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
	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
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

	npcs, err := a.npcRepo.ListByZone(zoneID)
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
	factionID, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
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

	npcs, err := a.npcRepo.ListByFaction(factionID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	faction, err := a.factionRepo.Get(factionID)
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
		Site   site
		Npc    *model.Npc
		Items  []*model.Item
		Map    *model.ZoneLevel
		Spawns []*model.Spawn
	}

	if strings.ToLower(getVar(r, "npcID")) == "search" {
		a.searchNpc(w, r)
		return
	}
	if strings.ToLower(getVar(r, "npcID")) == "byzone" {
		a.listNpcByZone(w, r)
		return
	}

	if strings.ToLower(getVar(r, "npcID")) == "byfaction" {
		a.listNpcByFaction(w, r)
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	npc, err := a.npcRepo.Get(npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error on npc")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	mapData, err := a.zoneLevelRepo.Get(npc.ZoneID())
	if err != nil {
		//err = errors.Wrap(err, "Request error on zonelevel")
		//a.writeError(w, r, err, http.StatusBadRequest)
		//return
	}

	entrys, _, err := a.spawnEntryRepo.ListByNpc(npc.ID)
	if err != nil {
		err = errors.Wrap(err, "Request error on spawnentry")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	spawns := []*model.Spawn{}

	for _, entry := range entrys {

		spawnGroups := []*model.Spawn{}
		spawnGroups, err = a.spawnRepo.ListBySpawnGroup(entry.SpawngroupID)
		if err != nil {
			err = errors.Wrap(err, "Request error on spawn")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
		for _, spawn := range spawnGroups {
			spawn.XScaled -= mapData.MapXOffset
			spawn.XScaled *= mapData.MapAspect

			spawn.YScaled -= mapData.MapYOffset
			spawn.YScaled *= mapData.MapAspect

			spawns = append(spawns, spawn)
		}
	}

	itemsLoots, err := a.npcLootRepo.List(npcID)
	if err != nil {
		err = errors.Wrap(err, "Request error on items")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "npc"
	site.Title = fmt.Sprintf("%s in %s", npc.CleanName(), npc.ZoneName())

	content := Content{
		Site:   site,
		Npc:    npc,
		Map:    mapData,
		Spawns: spawns,
	}

	var item *model.Item
	for _, itemLoot := range itemsLoots {
		item, err = a.itemRepo.Get(itemLoot.ItemID)
		if err != nil {
			err = errors.Wrap(err, "Request error on item")
			return
		}
		item.Reference = "Drop"
		content.Items = append(content.Items, item)
	}

	var merchentrys []*model.MerchantEntry
	if npc.MerchantID > 0 {
		merchentrys, _, err = a.merchantEntryRepo.List(npc.MerchantID)
		if err != nil {
			err = errors.Wrap(err, "Request error on merchant entries")
			return
		}
		for _, entry := range merchentrys {
			item, err = a.itemRepo.Get(entry.ItemID)
			if err != nil {
				err = errors.Wrap(err, "Request error on item")
				return
			}
			item.Reference = "Merchant"
			content.Items = append(content.Items, item)
		}
	}
	content.Site.Description = fmt.Sprintf("%s is a level %d %s %s found in %s who drops %d items and spawns at %d locations", npc.CleanName(), npc.Level, npc.RaceName(), npc.ClassName(), npc.ZoneName(), len(content.Items), len(spawns))
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
