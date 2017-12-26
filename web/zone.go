package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Zones []*model.Zone
	}

	site := a.NewSite()
	site.Page = "zone"
	site.Title = "Zone"

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
		tmp, err = a.loadTemplate(nil, "body", "listzone.tpl")
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

	site := a.NewSite()
	site.Page = "zone"
	site.Title = "Zone"

	content := Content{
		Site: site,
		Zone: zone,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "getzone.tpl")
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
