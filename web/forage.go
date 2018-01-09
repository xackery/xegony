package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) forageRoutes() (routes []*route) {
	routes = []*route{

		//Forage
		{
			"GetForage",
			"GET",
			"/forage/{forageID}",
			a.getForage,
		},
		{
			"ListForage",
			"GET",
			"/forage",
			a.listForage,
		},
	}
	return
}

func (a *Web) listForage(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		Forages    []*model.Forage
		ForagePage *model.Page
	}

	site := a.newSite(r)
	site.Page = "forage"
	site.Title = "Forage"

	foragePage := &model.Page{
		Scope: "forage",
	}
	foragePage.PageSize = getIntParam(r, "pageSize")
	foragePage.PageNumber = getIntParam(r, "pageNumber")

	forages, err := a.forageRepo.List(foragePage.PageSize, foragePage.PageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, forage := range forages {
		if forage.ItemID > 0 {
			forage.Item, err = a.itemRepo.Get(forage.ItemID)
			if err != nil {
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
		if forage.ZoneID > 0 {
			forage.Zone, err = a.zoneRepo.Get(forage.ZoneID)
			if err != nil {
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
	}
	foragePage.Total, err = a.forageRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:       site,
		Forages:    forages,
		ForagePage: foragePage,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forage/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("forage", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listForageByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "forage"
	site.Title = "Forage"

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
		tmp, err = a.loadTemplate(nil, "body", "forage/listbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("foragelistbyzone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getForageByZone(w http.ResponseWriter, r *http.Request) {
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
	site.Page = "forage"
	site.Title = "Forage"

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
		tmp, err = a.loadTemplate(nil, "body", "forage/getbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("foragelistbyzone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getForage(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Forage *model.Forage
	}

	if strings.ToLower(getVar(r, "forageID")) == "byzone" {
		a.listForageByZone(w, r)
		return
	}

	forageID, err := getIntVar(r, "forageID")
	if err != nil {
		err = errors.Wrap(err, "forageID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	forage, err := a.forageRepo.Get(forageID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	if forage.ItemID > 0 {
		forage.Item, err = a.itemRepo.Get(forage.ItemID)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	if forage.ZoneID > 0 {
		forage.Zone, err = a.zoneRepo.Get(forage.ZoneID)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "forage"
	site.Title = "Forage"

	content := Content{
		Site:   site,
		Forage: forage,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forage/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("forage", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
