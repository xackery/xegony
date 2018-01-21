package web

import (
	"html/template"
	"net/http"

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

func (a *Web) listForage(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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

	forages, err := a.forageRepo.List(foragePage.PageSize, foragePage.PageNumber, user)
	if err != nil {
		return
	}

	for _, forage := range forages {
		if forage.ItemID > 0 {
			item := &model.Item{
				ID: forage.ItemID,
			}
			err = a.itemRepo.Get(item, user)
			if err != nil {
				return
			}
		}
		if forage.ZoneID > 0 {
			zone := &model.Zone{
				ZoneIDNumber: forage.ZoneID,
			}
			err = a.zoneRepo.Get(zone, user)
			if err != nil {
				return
			}
		}
	}
	foragePage.Total, err = a.forageRepo.ListCount(user)
	if err != nil {
		return
	}
	content = Content{
		Site:       site,
		Forages:    forages,
		ForagePage: foragePage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forage/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("forage", tmp)
	}

	return
}

func (a *Web) listForageByZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "forage"
	site.Title = "Forage"

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
		tmp, err = a.loadTemplate(nil, "body", "forage/listbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("foragelistbyzone", tmp)
	}

	return
}

func (a *Web) getForageByZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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

	site := a.newSite(r)
	site.Page = "forage"
	site.Title = "Forage"

	zone := &model.Zone{
		ZoneIDNumber: zoneID,
	}

	err = a.zoneRepo.Get(zone, user)
	if err != nil {
		return
	}

	npcLoots, err := a.npcLootRepo.ListByZone(zone, user)
	if err != nil {
		return
	}

	content = Content{
		Site:     site,
		NpcLoots: npcLoots,
		Zone:     zone,
	}
	if err != nil {
		err = errors.Wrap(err, "failed to get zone")
		return
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forage/getbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("foragelistbyzone", tmp)
	}

	return
}

func (a *Web) getForage(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Forage *model.Forage
	}

	forageID, err := getIntVar(r, "forageID")
	if err != nil {
		err = errors.Wrap(err, "forageID argument is required")
		return
	}
	forage := &model.Forage{
		ID: forageID,
	}
	err = a.forageRepo.Get(forage, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	if forage.ItemID > 0 {
		item := &model.Item{
			ID: forage.ItemID,
		}
		err = a.itemRepo.Get(item, user)
		if err != nil {
			return
		}
	}
	if forage.ZoneID > 0 {
		zone := &model.Zone{
			ZoneIDNumber: forage.ZoneID,
		}
		err = a.zoneRepo.Get(zone, user)
		if err != nil {
			return
		}
	}

	site := a.newSite(r)
	site.Page = "forage"
	site.Title = "Forage"

	content = Content{
		Site:   site,
		Forage: forage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "forage/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("forage", tmp)
	}

	return
}
