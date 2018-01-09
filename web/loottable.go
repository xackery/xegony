package web

import (
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

func (a *Web) listLootTable(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		LootTables []*model.LootTable
	}

	site := a.newSite(r)
	site.Page = "lootTable"
	site.Title = "LootTable"
	site.Section = "lootTable"

	lootTables, err := a.lootTableRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:       site,
		LootTables: lootTables,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "loottable/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("lootTable", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getLootTable(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		LootTable *model.LootTable
	}

	id, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootTable, err := a.lootTableRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "lootTable"
	site.Title = "LootTable"
	site.Section = "lootTable"

	content := Content{
		Site:      site,
		LootTable: lootTable,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "loottable/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("loottable", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
