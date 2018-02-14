package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func zoneRoutes() (routes []*route) {
	routes = []*route{
		//Zone
		{
			"ListZone",
			"GET",
			"/zone",
			listZone,
		},
		{
			"GetZone",
			"GET",
			"/zone/{zoneID:[0-9]+}",
			getZone,
		},
		{
			"GetZoneEditor",
			"GET",
			"/zone/{zoneID:[0-9]+}/editor",
			getZoneEditor,
		},
	}
	return
}

func listZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := newSite(r, user)
	page := &model.Page{}
	zones, err := cases.ListZone(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:  site,
		Zones: zones,
	}

	tmp, err = loadTemplate(nil, "body", "zone/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Zone *model.Zone
	}

	zoneID := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		return
	}
	zone := &model.Zone{
		ID: zoneID,
	}

	err = cases.GetZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := newSite(r, user)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	content = Content{
		Site: site,
		Zone: zone,
	}

	tmp, err = loadTemplate(nil, "body", "zone/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getZoneEditor(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Zone   *model.Zone
		Spawns map[int64]*model.Spawn
	}

	zoneID := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		return
	}
	zone := &model.Zone{
		ID: zoneID,
	}

	err = cases.GetZone(zone, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	page := &model.Page{
		Limit: 500,
	}

	spawns := map[int64]*model.Spawn{}
	spawn := &model.Spawn{}
	spawnEntry := &model.SpawnEntry{
		ZoneShortName: zone.ShortName,
	}

	//start by entrys, so we can search by zone
	spawnEntrys, err := cases.ListSpawnEntryBySearch(page, spawn, spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "failed to list spawnEntrys by zone")
		return
	}

	//get details of each spawn
	for _, spawnEntry := range spawnEntrys {
		spawn = &model.Spawn{
			ID: spawnEntry.SpawnID,
		}
		err = cases.GetSpawn(spawn, user)
		if err != nil {
			log.Println("Spawn", spawn.ID, "doesn't exist")
		}
		npcPage := &model.Page{
			Limit: 500,
		}
		spawn.Entrys = append(spawn.Entrys, spawnEntry)
		spawn.Npcs, err = cases.ListSpawnNpc(npcPage, spawn, user)
		if err != nil {
			err = errors.Wrap(err, "failed to get spawn npc")
			return
		}
		spawns[spawn.ID] = spawn
	}

	site := newSite(r, user)
	site.Page = "zone"
	site.Title = "Zone"
	site.Section = "zone"

	content = Content{
		Site:   site,
		Zone:   zone,
		Spawns: spawns,
	}

	tmp, err = loadTemplate(nil, "body", "zone/editor.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
