package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func spawnRoutes() (routes []*route) {
	routes = []*route{
		//Spawn
		{
			"ListSpawn",
			"GET",
			"/spawn",
			listSpawn,
		},
		{
			"GetSpawn",
			"GET",
			"/spawn/{spawnID:[0-9]+}",
			getSpawn,
		},
	}
	return
}

func listSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Spawns []*model.Spawn
	}

	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}
	spawns, err := cases.ListSpawn(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:   site,
		Spawns: spawns,
	}

	tmp, err = loadTemplate(nil, "body", "spawn/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Spawn *model.Spawn
	}

	spawnID := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}
	spawn := &model.Spawn{
		ID: spawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := newSite(r, user)
	site.Page = "spawn"
	site.Title = "Spawn"
	site.Section = "spawn"

	content = Content{
		Site:  site,
		Spawn: spawn,
	}

	tmp, err = loadTemplate(nil, "body", "spawn/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
