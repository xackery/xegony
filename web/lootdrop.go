package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listLootDrop(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		LootDrops []*model.LootDrop
	}

	site := a.newSite(r)
	site.Page = "lootDrop"
	site.Title = "LootDrop"
	site.Section = "lootDrop"

	lootDrops, err := a.lootDropRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:      site,
		LootDrops: lootDrops,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdrop/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("lootDrop", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getLootDrop(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		LootDrop *model.LootDrop
	}

	id, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	lootDrop, err := a.lootDropRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "lootDrop"
	site.Title = "LootDrop"
	site.Section = "lootDrop"

	content := Content{
		Site:     site,
		LootDrop: lootDrop,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdrop/get.tpl")
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
