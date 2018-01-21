package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) zoneRoutes() (routes []*route) {
	routes = []*route{
		//Zone
		{
			"GetZone",
			"GET",
			"/zone/{zoneID:[0-9]+}",
			a.getZone,
		},
		{
			"ListZone",
			"GET",
			"/zone",
			a.listZone,
		},
		{
			"ListZoneByLevels",
			"GET",
			"/zone/bylevels",
			a.listZoneByLevels,
		},
		{
			"ListZoneByHotzone",
			"GET",
			"/zone/byhotzone",
			a.listZoneByHotzone,
		},
	}
	return
}

func (a *Web) listZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

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
		tmp, err = a.loadTemplate(nil, "body", "zone/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("zone", tmp)
	}

	return
}

func (a *Web) listZoneByLevels(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

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
		tmp, err = a.loadTemplate(nil, "body", "zone/listbylevels.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("zone", tmp)
	}

	return
}

func (a *Web) listZoneByHotzone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []model.Zone
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	zones, err := a.zoneRepo.ListByHotzone(user)
	if err != nil {
		return
	}
	content = Content{
		Site:  site,
		Zones: zones,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "zone/listbyhotzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("zone", tmp)
	}

	return
}

func (a *Web) getZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Zone *model.Zone
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
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	content = Content{
		Site: site,
		Zone: zone,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "zone/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("zone", tmp)
	}

	return
}

func (a *Web) listZoneByExpansion(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones map[string][]*model.Zone
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	zones, err := a.zoneRepo.List(user)
	if err != nil {
		return
	}

	zonesByExpansion := map[string][]*model.Zone{}
	for _, zone := range zones {
		zonesByExpansion[zone.ExpansionName()] = append(zonesByExpansion[zone.ExpansionName()], zone)
	}

	content = Content{
		Site:  site,
		Zones: zonesByExpansion,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "zone/listbyexpansion.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("zone", tmp)
	}

	return
}
