package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listFishing(w http.ResponseWriter, r *http.Request) {
	var err error

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

	fishings, err := a.fishingRepo.List(fishingPage.PageSize, fishingPage.PageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, fishing := range fishings {
		if fishing.ItemID > 0 {
			fishing.Item, err = a.itemRepo.Get(fishing.ItemID)
			if err != nil {
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
		if fishing.ZoneID > 0 {
			fishing.Zone, err = a.zoneRepo.Get(fishing.ZoneID)
			if err != nil {
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
		if fishing.NpcID > 0 {
			fishing.Npc, err = a.npcRepo.Get(fishing.NpcID)
			if err != nil {
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
	}
	fishingPage.Total, err = a.fishingRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:        site,
		Fishings:    fishings,
		FishingPage: fishingPage,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "fishing/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("fishing", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listFishingByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "fishing"
	site.Title = "Fishing"

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
		tmp, err = a.loadTemplate(nil, "body", "fishing/listbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("fishinglistbyzone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getFishingByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Zone     *model.Zone
		NpcLoots []*model.NpcLoot
	}

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "fishing"
	site.Title = "Fishing"

	npcLoots, err := a.npcLootRepo.ListByZone(zoneID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	zone, err := a.zoneRepo.Get(zoneID)
	content := Content{
		Site:     site,
		NpcLoots: npcLoots,
		Zone:     zone,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "fishing/getbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("fishinglistbyzone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getFishing(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site    site
		Fishing *model.Fishing
	}

	if strings.ToLower(getVar(r, "fishingID")) == "byzone" {
		a.listFishingByZone(w, r)
		return
	}

	fishingID, err := getIntVar(r, "fishingID")
	if err != nil {
		err = errors.Wrap(err, "fishingID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	fishing, err := a.fishingRepo.Get(fishingID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	if fishing.ItemID > 0 {
		fishing.Item, err = a.itemRepo.Get(fishing.ItemID)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	if fishing.ZoneID > 0 {
		fishing.Zone, err = a.zoneRepo.Get(fishing.ZoneID)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	if fishing.NpcID > 0 {
		fishing.Npc, err = a.npcRepo.Get(fishing.NpcID)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "fishing"
	site.Title = "Fishing"

	content := Content{
		Site:    site,
		Fishing: fishing,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "fishing/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("fishing", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
