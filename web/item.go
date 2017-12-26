package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListItem(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Items []*model.Item
	}

	site := a.NewSite()
	site.Page = "item"
	site.Title = "Item"

	items, err := a.itemRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:  site,
		Items: items,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "listitem.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("item", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) ListItemByCharacter(w http.ResponseWriter, r *http.Request) {
	var err error
	characterId, err := getIntVar(r, "characterId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site      Site
		Inventory []*model.Item
		Character *model.Character
	}

	site := a.NewSite()
	site.Page = "item"
	site.Title = "Item"
	character, err := a.characterRepo.Get(characterId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	items, err := a.itemRepo.ListByCharacter(characterId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:      site,
		Inventory: items,
		Character: character,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "listitembycharacter.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("item", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
func (a *Web) GetItem(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Item *model.Item
	}

	id, err := getIntVar(r, "itemId")
	if err != nil {
		err = errors.Wrap(err, "itemId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	item, err := a.itemRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite()
	site.Page = "item"
	site.Title = "Item"

	content := Content{
		Site: site,
		Item: item,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "getitem.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("item", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
