package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Zones []*model.Zone
	}

	site := a.NewSite(r)
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

func (a *Web) GetZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Zone *model.Zone
	}

	if strings.ToLower(getVar(r, "zoneId")) == "byexpansion" {
		a.ListZoneByExpansion(w, r)
		return
	}

	id, err := getIntVar(r, "zoneId")
	if err != nil {
		err = errors.Wrap(err, "zoneId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	zone, err := a.zoneRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
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

func (a *Web) ListZoneByExpansion(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Zones map[string][]*model.Zone
	}

	site := a.NewSite(r)
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
