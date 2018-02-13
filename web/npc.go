package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func npcRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListNpc",
			"GET",
			"/npc",
			listNpc,
		},
		{
			"ListNpcZone",
			"GET",
			"/npc/zone",
			listNpcZone,
		},
		{
			"ListNpcZone",
			"GET",
			"/npc/zone",
			listNpcZone,
		},
		{
			"ListNpcByZone",
			"GET",
			"/npc/zone/{zoneID:[0-9]+}",
			listNpcByZone,
		},
		{
			"GetNpc",
			"GET",
			"/npc/{npcID:[0-9]+}",
			getNpc,
		},
	}
	return
}

func listNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Page *model.Page
		Npcs []*model.Npc
	}

	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}
	npcs, err := cases.ListNpc(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site: site,
		Npcs: npcs,
		Page: page,
	}

	tmp, err = loadTemplate(nil, "body", "npc/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func listNpcZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Page  *model.Page
		Zones []*model.Zone
	}

	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}
	zones, err := cases.ListZone(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:  site,
		Zones: zones,
		Page:  page,
	}

	tmp, err = loadTemplate(nil, "body", "npc/list_zone.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func listNpcByZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Page *model.Page
		Zone *model.Zone
		Npcs []*model.Npc
	}

	zone := &model.Zone{
		ID: getIntVar(r, "zoneID"),
	}

	err = cases.GetZone(zone, user)
	if err != nil {
		return
	}
	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}

	npcs, err := cases.ListNpcByZone(page, zone, user)
	if err != nil {
		return
	}

	content = Content{
		Site: site,
		Zone: zone,
		Page: page,
		Npcs: npcs,
	}

	tmp, err = loadTemplate(nil, "body", "npc/list_by_zone.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Npc   *model.Npc
		Loot  *model.Loot
		Items []*model.Item
	}

	npcID := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		return
	}
	npc := &model.Npc{
		ID: npcID,
	}

	err = cases.GetNpc(npc, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	loot := &model.Loot{
		ID: npc.LootID,
	}
	err = cases.GetLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	site := newSite(r, user)
	site.Page = "npc"
	site.Title = "Npc"
	site.Section = "npc"

	content = Content{
		Site: site,
		Npc:  npc,
		Loot: loot,
	}

	//loot.Entries[0].DropEntrys[0].Item

	tmp, err = loadTemplate(nil, "body", "npc/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
