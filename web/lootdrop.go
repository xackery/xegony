package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) lootDropRoutes() (routes []*route) {
	routes = []*route{

		//LootDropEntry
		{
			"GetLootDropEntry",
			"GET",
			"/lootdrop/{lootDropID}/{itemID}",
			a.getLootDropEntry,
		},
		{
			"ListLootDropEntry",
			"GET",
			"/lootdrop/{lootDropID}",
			a.listLootDropEntry,
		},
	}
	return
}

func (a *Web) listLootDrop(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		LootDrops []*model.LootDrop
	}

	site := a.newSite(r)
	site.Page = "lootDrop"
	site.Title = "LootDrop"
	site.Section = "lootDrop"

	lootDrops, err := a.lootDropRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:      site,
		LootDrops: lootDrops,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdrop/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("lootDrop", tmp)
	}

	return
}

func (a *Web) getLootDrop(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		LootDrop *model.LootDrop
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}
	lootDrop := &model.LootDrop{
		ID: lootDropID,
	}

	err = a.lootDropRepo.Get(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "lootDrop"
	site.Title = "LootDrop"
	site.Section = "lootDrop"

	content = Content{
		Site:     site,
		LootDrop: lootDrop,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdrop/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("lootdrop", tmp)
	}

	return
}
