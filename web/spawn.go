package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) spawnRoutes() (routes []*route) {
	routes = []*route{
		//Spawn
		{
			"ListSpawn",
			"GET",
			"/spawn",
			a.listSpawn,
		},
		{
			"GetSpawn",
			"GET",
			"/spawn/{spawnID}/details",
			a.getSpawn,
		},
	}
	return
}

func (a *Web) listSpawn(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Spawns []*model.Spawn
	}

	site := a.newSite(r)
	site.Page = "spawn"
	site.Title = "Spawn"
	site.Section = "spawn"

	spawns, err := a.spawnRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	content := Content{
		Site:   site,
		Spawns: spawns,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("spawn", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getSpawn(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Spawn *model.Spawn
	}

	id, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	spawn, err := a.spawnRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "spawn"
	site.Title = "Spawn"
	site.Section = "spawn"

	content := Content{
		Site:  site,
		Spawn: spawn,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("spawn", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
