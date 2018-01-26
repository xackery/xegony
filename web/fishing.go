package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) fishingRoutes() (routes []*route) {
	routes = []*route{
		//Fishing
		{
			"GetFishing",
			"GET",
			"/fishing/{fishingID:[0-9]+}",
			a.getFishing,
		},
		{
			"ListFishingByZone",
			"GET",
			"/fishing/zone",
			a.listFishingByZone,
		},
		{
			"ListFishing",
			"GET",
			"/fishing",
			a.listFishing,
		},
	}
	return
}

func (a *Web) listFishing(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site        site
		Fishings    []*model.Fishing
		FishingPage *model.Page
	}

	site := a.newSite(r)
	site.Page = "fishing"
	site.Title = "Fishing"

	fishingPage := &model.Page{
		Scope: "fishing",
	}
	fishingPage.PageSize = getIntParam(r, "pageSize")
	fishingPage.PageNumber = getIntParam(r, "pageNumber")

	fishings, err := a.fishingRepo.List(fishingPage.PageSize, fishingPage.PageNumber, user)
	if err != nil {
		return
	}

	for _, fishing := range fishings {
		if fishing.ItemID > 0 {
			item := &model.Item{
				ID: fishing.ItemID,
			}
			err = a.itemRepo.Get(item, user)
			if err != nil {
				return
			}
		}
		if fishing.ZoneID > 0 {
			zone := &model.Zone{
				ZoneIDNumber: fishing.ZoneID,
			}
			err = a.zoneRepo.Get(zone, user)
			if err != nil {
				return
			}
		}
		if fishing.NpcID > 0 {
			npc := &model.Npc{
				ID: fishing.NpcID,
			}
			err = a.npcRepo.Get(npc, user)
			if err != nil {
				return
			}
		}
	}
	fishingPage.Total, err = a.fishingRepo.ListCount(user)
	if err != nil {
		return
	}
	content = Content{
		Site:        site,
		Fishings:    fishings,
		FishingPage: fishingPage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "fishing/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("fishing", tmp)
	}

	return
}

func (a *Web) listFishingByZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "fishing"
	site.Title = "Fishing"

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
		tmp, err = a.loadTemplate(nil, "body", "fishing/listbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("fishinglistbyzone", tmp)
	}

	return
}

func (a *Web) getFishingByZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		Zone     *model.Zone
		NpcLoots []*model.NpcLoot
	}

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		return
	}

	zone := &model.Zone{
		ZoneIDNumber: zoneID,
	}
	err = a.zoneRepo.Get(zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get zone")
	}

	site := a.newSite(r)
	site.Page = "fishing"
	site.Title = "Fishing"

	npcLoots, err := a.npcLootRepo.ListByZone(zone, user)
	if err != nil {
		return
	}

	content = Content{
		Site:     site,
		NpcLoots: npcLoots,
		Zone:     zone,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "fishing/getbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("fishinglistbyzone", tmp)
	}

	return
}

func (a *Web) getFishing(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site    site
		Fishing *model.Fishing
	}

	fishingID, err := getIntVar(r, "fishingID")
	if err != nil {
		err = errors.Wrap(err, "fishingID argument is required")
		return
	}
	fishing := &model.Fishing{
		ID: fishingID,
	}
	err = a.fishingRepo.Get(fishing, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	if fishing.ItemID > 0 {
		item := &model.Item{
			ID: fishing.ItemID,
		}
		err = a.itemRepo.Get(item, user)
		if err != nil {
			return
		}
		fishing.Item = item
	}
	if fishing.ZoneID > 0 {
		zone := &model.Zone{
			ID: fishing.ZoneID,
		}
		err = a.zoneRepo.Get(zone, user)
		if err != nil {
			return
		}
		fishing.Zone = zone
	}
	if fishing.NpcID > 0 {
		npc := &model.Npc{
			ID: fishing.NpcID,
		}
		err = a.npcRepo.Get(npc, user)
		if err != nil {
			return
		}
		fishing.Npc = npc
	}

	site := a.newSite(r)
	site.Page = "fishing"
	site.Title = "Fishing"

	content = Content{
		Site:    site,
		Fishing: fishing,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "fishing/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("fishing", tmp)
	}

	return
}
