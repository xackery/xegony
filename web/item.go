package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func itemRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListItem",
			"GET",
			"/item",
			listItem,
		},
		{
			"ListItemZone",
			"GET",
			"/item/zone",
			listItemZone,
		},
		{
			"ListItemZone",
			"GET",
			"/item/zone",
			listItemZone,
		},
		{
			"ListItemByZone",
			"GET",
			"/item/zone/{zoneID:[0-9]+}",
			listItemByZone,
		},
		{
			"GetItem",
			"GET",
			"/item/{itemID:[0-9]+}",
			getItem,
		},
	}
	return
}

func listItem(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Page  *model.Page
		Items []*model.Item
	}

	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}
	items, err := cases.ListItem(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:  site,
		Items: items,
		Page:  page,
	}

	tmp, err = loadTemplate(nil, "body", "item/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func listItemZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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

	tmp, err = loadTemplate(nil, "body", "item/list_zone.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func listItemByZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Page  *model.Page
		Zone  *model.Zone
		Items []*model.Item
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

	npcPage := &model.Page{
		Limit: 500,
	}
	//Get all npcs of a zone
	npcs, err := cases.ListNpcByZone(npcPage, zone, user)
	if err != nil {
		err = errors.Wrap(err, "failed to list npcs by zone")
		return
	}

	items := []*model.Item{}
	for _, npc := range npcs {

		loot := &model.Loot{
			ID: npc.LootID,
		}
		//Now we need all items of an npc, this is gross.
		err = cases.GetLoot(loot, user)
		if err != nil {
			err = errors.Wrap(err, "failed to list loot by npc")
		}
		for _, entry := range loot.Entrys {
			for _, drop := range entry.DropEntrys {
				items = append(items, drop.Item)
			}
		}
	}

	content = Content{
		Site:  site,
		Zone:  zone,
		Page:  page,
		Items: items,
	}

	tmp, err = loadTemplate(nil, "body", "item/list_by_zone.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getItem(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Item *model.Item
		Npcs []*model.Npc
	}

	itemID := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		return
	}

	item := &model.Item{
		ID: itemID,
	}

	err = cases.GetItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := newSite(r, user)
	site.Page = "item"
	site.Title = "Item"
	site.Section = "item"

	content = Content{
		Site: site,
		Item: item,
	}

	//loot.Entries[0].DropEntrys[0].Item

	tmp, err = loadTemplate(nil, "body", "item/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
