package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) lootTableRoutes() (routes []*route) {
	routes = []*route{

		//LootTable
		{
			"GetLootTable",
			"GET",
			"/loottable/{lootTableID}",
			a.getLootTable,
		},
		{
			"ListLootTable",
			"GET",
			"/loottable",
			a.listLootTable,
		},
	}
	return
}

func (a *Web) listLootTable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		LootTables []*model.LootTable
	}

	site := a.newSite(r)
	site.Page = "lootTable"
	site.Title = "LootTable"
	site.Section = "lootTable"

	lootTables, err := a.lootTableRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:       site,
		LootTables: lootTables,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "loottable/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("lootTable", tmp)
	}

	return
}

func (a *Web) getLootTable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		LootTable *model.LootTable
	}

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}
	lootTable := &model.LootTable{
		ID: lootTableID,
	}
	err = a.lootTableRepo.Get(lootTable, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "lootTable"
	site.Title = "LootTable"
	site.Section = "lootTable"

	content = Content{
		Site:      site,
		LootTable: lootTable,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "loottable/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("loottable", tmp)
	}

	return
}
