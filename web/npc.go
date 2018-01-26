package web

import (
	"fmt"
	"html/template"
	"net/http"

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

func (a *Web) listNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site    site
		Npcs    []*model.Npc
		NpcPage *model.Page
	}

	npcPage := &model.Page{
		Scope: "npc",
	}
	npcPage.PageSize = getIntParam(r, "pageSize")
	npcPage.PageNumber = getIntParam(r, "pageNumber")

	site := a.newSite(r)
	site.Page = "npclist"
	site.Title = "Npc"
	site.Section = "npc"

	npcs, err := a.npcRepo.List(npcPage.PageSize, npcPage.PageNumber, user)
	if err != nil {
		return
	}
	npcPage.Total, err = a.npcRepo.ListCount(user)
	if err != nil {
		return
	}

	content = Content{
		Site:    site,
		Npcs:    npcs,
		NpcPage: npcPage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("npc", tmp)
	}

	return
}

func (a *Web) searchNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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
		npc := &model.Npc{
			Name: search,
		}
		npcs, err = a.npcRepo.SearchByName(npc, user)
		if err != nil {
			return
		}
	}
	content = Content{
		Site:   site,
		Npcs:   npcs,
		Search: search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/search.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("npcsearch", tmp)
	}

	return
}

func (a *Web) listNpcByZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "listnpcbyzone"
	site.Title = "Npc"
	site.Section = "npc"

	zones, err := a.zoneRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:  site,
		Zones: zones,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/listbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("npc", tmp)
	}

	return
}

func (a *Web) getNpcByZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
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

	zone := &model.Zone{
		ZoneIDNumber: zoneID,
	}
	npcs, err := a.npcRepo.ListByZone(zone, user)
	if err != nil {
		return
	}
	content = Content{
		Site: site,
		Npcs: npcs,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/getbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("npc", tmp)
	}

	return
}

func (a *Web) listNpcByFaction(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		Factions []*model.Faction
	}

	site := a.newSite(r)
	site.Page = "listnpcbyfaction"
	site.Title = "Npc"
	site.Section = "npc"

	factions, err := a.factionRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:     site,
		Factions: factions,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/listbyfaction.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("npc", tmp)
	}

	return
}

func (a *Web) getNpcByFaction(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	factionID, err := getIntVar(r, "factionID")
	if err != nil {
		err = errors.Wrap(err, "factionID argument is required")
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

	faction := &model.Faction{
		ID: factionID,
	}
	npcs, err := a.npcRepo.ListByFaction(faction, user)
	if err != nil {
		return
	}

	err = a.factionRepo.Get(faction, user)
	if err != nil {
		return
	}
	content = Content{
		Site:    site,
		Npcs:    npcs,
		Faction: faction,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/getbyfaction.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("npc", tmp)
	}

	return
}

func (a *Web) getNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Npc    *model.Npc
		Items  []*model.Item
		Map    *model.ZoneLevel
		Spawns []*model.Spawn
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		return
	}
	npc := &model.Npc{
		ID: npcID,
	}
	err = a.npcRepo.Get(npc, user)
	if err != nil {
		err = errors.Wrap(err, "Request error on npc")
		return
	}

	zoneLevel := &model.ZoneLevel{
		ZoneID: npc.Zone.ID,
	}
	err = a.zoneLevelRepo.Get(zoneLevel, user)
	if err != nil {
		//err = errors.Wrap(err, "Request error on zonelevel")
		//		//return
	}
	/*
		entrys, _, err := a.spawnEntryRepo.ListByNpc(npc.ID)
		if err != nil {
			err = errors.Wrap(err, "Request error on spawnentry")
			return
		}

		spawns := []*model.Spawn{}

		for _, entry := range entrys {

			spawnGroups := []*model.Spawn{}
			spawnGroups, err = a.spawnRepo.ListBySpawnGroup(entry.SpawngroupID)
			if err != nil {
				err = errors.Wrap(err, "Request error on spawn")
				return
			}
			for _, spawn := range spawnGroups {
				spawn.XScaled -= zoneLevel.MapXOffset
				spawn.XScaled *= zoneLevel.MapAspect

				spawn.YScaled -= zoneLevel.MapYOffset
				spawn.YScaled *= zoneLevel.MapAspect

				spawns = append(spawns, spawn)
			}
		}

		itemsLoots, err := a.npcLootRepo.List(npcID)
		if err != nil {
			err = errors.Wrap(err, "Request error on items")
			return
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
		}*/

	site := a.newSite(r)
	site.Page = "npc"
	site.Title = fmt.Sprintf("%s in %s", npc.CleanName, npc.Zone.LongName)

	out := Content{
		Site: site,
		Npc:  npc,
		Map:  zoneLevel,
		//Spawns: spawns,
	}
	//out.Site.Description = fmt.Sprintf("%s is a level %d %s %s found in %s who drops %d items and spawns at %d locations", npc.CleanName, npc.Level, npc.Race.Name, npc.Class.Name, npc.Zone.LongName, len(content.Items), len(spawns))
	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "npc/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("npc", tmp)
	}

	content = out
	return
}
