package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) spawnEntryRoutes() (routes []*route) {
	routes = []*route{
		//SpawnEntry
		{
			"ListSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry",
			a.listSpawnEntry,
		},
		{
			"GetSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry/{spawnEntryID:[0-9]+}",
			a.getSpawnEntry,
		},
	}
	return
}

func (a *Web) listSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Out struct {
		Site        site
		Spawn       *model.Spawn
		SpawnEntrys []*model.SpawnEntry
	}

	site := a.newSite(r)
	site.Page = "spawnEntry"
	site.Title = "SpawnEntry"
	site.Section = "spawnEntry"

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawn := &model.Spawn{
		ID: spawnID,
	}

	err = a.spawnRepo.Get(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnEntrys, err := a.spawnEntryRepo.ListBySpawn(spawn, user)
	if err != nil {
		return
	}

	out := &Out{
		Site:        site,
		SpawnEntrys: spawnEntrys,
		Spawn:       spawn,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/entry/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spawnEntry", tmp)
	}

	content = out
	return
}

func (a *Web) getSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		SpawnEntry *model.SpawnEntry
	}

	spawnEntryID, err := getIntVar(r, "spawnEntryID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
		return
	}
	spawnEntry := &model.SpawnEntry{
		ID: spawnEntryID,
	}
	err = a.spawnEntryRepo.Get(spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "spawnEntry"
	site.Title = "SpawnEntry"
	site.Section = "spawnEntry"

	content = Content{
		Site:       site,
		SpawnEntry: spawnEntry,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/entry/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spawnEntry", tmp)
	}

	return
}
