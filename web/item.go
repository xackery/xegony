package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListItem(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Items []*model.Item
	}

	site := a.NewSite(r)
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
		tmp, err = a.loadTemplate(nil, "body", "item/list.tpl")
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

func (a *Web) ListItemByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Zones []*model.Zone
	}

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"

	zones, err := a.zoneRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:  site,
		Zones: zones,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/listbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("itemlistbyzone", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) GetItemByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Zone  *model.Zone
		Items []*model.Item
	}

	zoneId, err := getIntVar(r, "zoneId")
	if err != nil {
		err = errors.Wrap(err, "zoneId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"

	items, err := a.itemRepo.GetByZone(zoneId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	zone, err := a.zoneRepo.Get(zoneId)
	content := Content{
		Site:  site,
		Items: items,
		Zone:  zone,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/getbyzone.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("itemlistbyzone", tmp)
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
		Inventory map[int]model.Item
		Character *model.Character
	}

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"
	character, err := a.characterRepo.Get(characterId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	inventory, err := a.itemRepo.ListByCharacter(characterId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemInventory := map[int]model.Item{}

	for i, _ := range inventory {
		itemInventory[int(inventory[i].SlotId)] = *inventory[i]
	}

	content := Content{
		Site:      site,
		Inventory: itemInventory,
		Character: character,
	}

	fmt.Println(characterId, len(inventory))
	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/listbycharacter.tpl")
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

func (a *Web) ListItemBySlot(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Items []*model.Item
	}

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"
	items, err := a.itemRepo.ListBySlot()
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
		tmp, err = a.loadTemplate(nil, "body", "item/listbyslot.tpl")
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

func (a *Web) GetItemBySlot(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   Site
		Items  []*model.Item
		SlotId int64
	}

	slotId, err := getIntVar(r, "slotId")
	if err != nil {
		err = errors.Wrap(err, "slotId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"
	items, err := a.itemRepo.GetBySlot(slotId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	content := Content{
		Site:   site,
		Items:  items,
		SlotId: slotId,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/getbyslot.tpl")
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

	if strings.ToLower(getVar(r, "itemId")) == "byslot" {
		a.ListItemBySlot(w, r)
		return
	}

	if strings.ToLower(getVar(r, "itemId")) == "byzone" {
		a.ListItemByZone(w, r)
		return
	}

	if strings.ToLower(getVar(r, "itemId")) == "search" {
		a.SearchItem(w, r)
		return
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

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"

	content := Content{
		Site: site,
		Item: item,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/get.tpl")
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

func (a *Web) SearchItemByAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site    Site
		Items   []*model.Item
		Account []*model.Account
		Search  string
	}

	claims, err := api.GetAuthClaims(r)
	if err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	search := getParam(r, "search")

	var items []*model.Item

	if len(search) > 0 {
		items, err = a.itemRepo.SearchByAccount(claims.User.AccountId, search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"

	content := Content{
		Site:   site,
		Items:  items,
		Search: search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/searchbyaccount.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("itemsearchbyaccount", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) SearchItem(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   Site
		Items  []*model.Item
		Search string
	}

	search := getParam(r, "search")

	var items []*model.Item

	if len(search) > 0 {
		items, err = a.itemRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.NewSite(r)
	site.Page = "item"
	site.Title = "Item"

	content := Content{
		Site:   site,
		Items:  items,
		Search: search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("itemsearch", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
