package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

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
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTableEntrys, err := a.lootTableEntryRepo.List(lootTableID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:            site,
		LootTableEntrys: lootTableEntrys,
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

		a.setTemplate("lootTableEntry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site           site
		LootTableEntry *model.LootTableEntry
	}

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTableEntry, err := a.lootTableEntryRepo.Get(lootTableID, lootDropID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "lootTableEntry"
	site.Title = "LootTableEntry"
	site.Section = "lootTableEntry"

	content := Content{
		Site:           site,
		LootTableEntry: lootTableEntry,
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
