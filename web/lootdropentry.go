package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) lootDropEntryRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *Web) listLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site           site
		LootDropEntrys []*model.LootDropEntry
	}

	site := a.newSite(r)
	site.Page = "lootDropEntry"
	site.Title = "LootDropEntry"
	site.Section = "lootDropEntry"

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntrys, err := a.lootDropEntryRepo.List(lootDropID)
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

func (a *Web) getLootDropEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site          site
		LootDropEntry *model.LootDropEntry
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	lootDropEntry, err := a.lootDropEntryRepo.Get(lootDropID, itemID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
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
