package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/model"
)

func (a *Web) listItem(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Items    []*model.Item
		ItemPage *model.Page
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	itemPage := &model.Page{
		Scope: "item",
	}
	itemPage.PageSize = getIntParam(r, "pageSize")
	itemPage.PageNumber = getIntParam(r, "pageNumber")

	items, err := a.itemRepo.List(itemPage.PageSize, itemPage.PageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	itemPage.Total, err = a.itemRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:     site,
		Items:    items,
		ItemPage: itemPage,
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

func (a *Web) listItemByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
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

func (a *Web) getItemByZone(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Zone     *model.Zone
		NpcLoots []*model.NpcLoot
	}

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	npcLoots, err := a.npcLootRepo.ListByZone(zoneID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	zone, err := a.zoneRepo.Get(zoneID)
	content := Content{
		Site:     site,
		NpcLoots: npcLoots,
		Zone:     zone,
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

func (a *Web) listItemByCharacter(w http.ResponseWriter, r *http.Request) {
	var err error
	characterID, err := getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site      site
		Inventory map[int]model.Item
		Character *model.Character
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"
	character, err := a.characterRepo.Get(characterID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	inventory, err := a.itemRepo.ListByCharacter(characterID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	itemInventory := map[int]model.Item{}

	for i := range inventory {
		itemInventory[int(inventory[i].SlotID)] = *inventory[i]
	}

	content := Content{
		Site:      site,
		Inventory: itemInventory,
		Character: character,
	}

	fmt.Println(characterID, len(inventory))
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

func (a *Web) listItemBySlot(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  site
		Items []*model.Item
	}

	site := a.newSite(r)
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

func (a *Web) getItemBySlot(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
		Items  []*model.Item
		SlotID int64
	}

	slotID, err := getIntVar(r, "slotID")
	if err != nil {
		err = errors.Wrap(err, "slotID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"
	items, err := a.itemRepo.GetBySlot(slotID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	content := Content{
		Site:   site,
		Items:  items,
		SlotID: slotID,
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

func (a *Web) getItem(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Item     *model.Item
		Npcs     []*model.Npc
		Fishings []*model.Fishing
	}

	if strings.ToLower(getVar(r, "itemID")) == "byslot" {
		a.listItemBySlot(w, r)
		return
	}

	if strings.ToLower(getVar(r, "itemID")) == "byzone" {
		a.listItemByZone(w, r)
		return
	}

	if strings.ToLower(getVar(r, "itemID")) == "search" {
		a.searchItem(w, r)
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	item, err := a.itemRepo.Get(itemID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npcs, err := a.npcRepo.ListByItem(itemID)
	if err != nil {
		err = errors.Wrap(err, "Failed to get NPCs based on item")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	fishings, err := a.fishingRepo.GetByItem(itemID)
	if err != nil {
		err = errors.Wrap(err, "Failed to get fishings based on item")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, fishing := range fishings {
		if fishing.ZoneID > 0 {
			fishing.Zone, err = a.zoneRepo.Get(fishing.ZoneID)
			if err != nil {
				err = errors.Wrap(err, "Failed to get zone based on fishing zoneid")
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	content := Content{
		Site:     site,
		Item:     item,
		Npcs:     npcs,
		Fishings: fishings,
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

func (a *Web) searchItemByAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site    site
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
		items, err = a.itemRepo.SearchByAccount(claims.User.AccountID, search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
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

func (a *Web) searchItem(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site   site
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

	site := a.newSite(r)
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
