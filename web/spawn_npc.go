package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) spawnNpcRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListSpawnNpc",
			"GET",
			"/spawn/{spawnID:[0-9]+}/npc",
			a.listSpawnNpc,
		},
		{
			"GetSpawnNpc",
			"GET",
			"/spawn/{spawnID:[0-9]+}/npc/{spawnNpcID:[0-9]+}",
			a.getSpawnNpc,
		},
	}
	return
}

func (a *Web) listSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Out struct {
		Site      site
		Spawn     *model.Spawn
		SpawnNpcs []*model.SpawnNpc
	}

	site := a.newSite(r)
	site.Page = "spawnNpc"
	site.Title = "SpawnNpc"
	site.Section = "spawnNpc"

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

	spawnNpcs, err := a.spawnNpcRepo.ListBySpawn(spawn, user)
	if err != nil {
		return
	}

	out := &Out{
		Site:      site,
		SpawnNpcs: spawnNpcs,
		Spawn:     spawn,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/npc/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spawnNpc", tmp)
	}

	content = out
	return
}

func (a *Web) getSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		SpawnNpc *model.SpawnNpc
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		return
	}
	spawnNpc := &model.SpawnNpc{
		NpcID:   npcID,
		SpawnID: spawnID,
	}
	err = a.spawnNpcRepo.Get(spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "spawnNpc"
	site.Title = "SpawnNpc"
	site.Section = "spawnNpc"

	content = Content{
		Site:     site,
		SpawnNpc: spawnNpc,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "spawn/npc/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("spawnNpc", tmp)
	}

	return
}
