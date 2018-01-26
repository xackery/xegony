package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) lootTableEntryRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *Web) listLootTableEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site            site
		LootTableEntrys []*model.LootTableEntry
	}

	site := a.newSite(r)
	site.Page = "lootTableEntry"
	site.Title = "LootTableEntry"
	site.Section = "lootTableEntry"

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootTable := &model.LootTable{
		ID: lootTableID,
	}
	lootTableEntrys, err := a.lootTableEntryRepo.ListByLootTable(lootTable, user)
	if err != nil {
		return
	}
	content = Content{
		Site:            site,
		LootTableEntrys: lootTableEntrys,
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

		a.setTemplate("lootTableEntry", tmp)
	}

	return
}

func (a *Web) getLootTableEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site           site
		LootTableEntry *model.LootTableEntry
	}

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	lootTableEntry := &model.LootTableEntry{
		LoottableID: lootTableID,
		LootdropID:  lootDropID,
	}
	err = a.lootTableEntryRepo.Get(lootTableEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "lootTableEntry"
	site.Title = "LootTableEntry"
	site.Section = "lootTableEntry"

	content = Content{
		Site:           site,
		LootTableEntry: lootTableEntry,
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
