package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

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
		tmp, err = a.loadTemplate(nil, "body", "zone/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("zone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listZoneByHotzone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	zones, err := a.zoneRepo.ListByHotzone()
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
		tmp, err = a.loadTemplate(nil, "body", "zone/listbyhotzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("zone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site site
		Zone *model.Zone
	}

	if strings.ToLower(getVar(r, "zoneID")) == "byexpansion" {
		a.listZoneByExpansion(w, r)
		return
	}

	if strings.ToLower(getVar(r, "zoneID")) == "byhotzone" {
		a.listZoneByHotzone(w, r)
		return
	}

	id, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	zone, err := a.zoneRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	content := Content{
		Site: site,
		Zone: zone,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "zone/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("zone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listZoneByExpansion(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Zones map[string][]*model.Zone
	}

	site := a.newSite(r)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	zones, err := a.zoneRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	zonesByExpansion := map[string][]*model.Zone{}
	for _, zone := range zones {
		zonesByExpansion[zone.ExpansionName()] = append(zonesByExpansion[zone.ExpansionName()], zone)
	}

	content := Content{
		Site:  site,
		Zones: zonesByExpansion,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "zone/listbyexpansion.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("zone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
