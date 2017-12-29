package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site            Site
		LootTableEntrys []*model.LootTableEntry
	}

	site := a.NewSite(r)
	site.Page = "lootTableEntry"
	site.Title = "LootTableEntry"
	site.Section = "lootTableEntry"

	lootTableId, err := getIntVar(r, "lootTableId")
	if err != nil {
		err = errors.Wrap(err, "lootTableId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTableEntrys, err := a.lootTableEntryRepo.List(lootTableId)
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

func (a *Web) GetLootTableEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site           Site
		LootTableEntry *model.LootTableEntry
	}

	lootTableId, err := getIntVar(r, "lootTableId")
	if err != nil {
		err = errors.Wrap(err, "lootTableId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootTableEntry, err := a.lootTableEntryRepo.Get(lootTableId, lootDropId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
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
