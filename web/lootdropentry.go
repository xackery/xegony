package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site           Site
		LootDropEntrys []*model.LootDropEntry
	}

	site := a.NewSite(r)
	site.Page = "lootDropEntry"
	site.Title = "LootDropEntry"
	site.Section = "lootDropEntry"

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntrys, err := a.lootDropEntryRepo.List(lootDropId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:           site,
		LootDropEntrys: lootDropEntrys,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdropentry/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("lootDropEntry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) GetLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site          Site
		LootDropEntry *model.LootDropEntry
	}

	lootDropId, err := getIntVar(r, "lootDropId")
	if err != nil {
		err = errors.Wrap(err, "lootDropId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemId, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntry, err := a.lootDropEntryRepo.Get(lootDropId, itemId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "lootDropEntry"
	site.Title = "LootDropEntry"
	site.Section = "lootDropEntry"

	content := Content{
		Site:          site,
		LootDropEntry: lootDropEntry,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdropentry/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("lootdrop", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
