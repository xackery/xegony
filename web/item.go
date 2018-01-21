package web

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) itemRoutes() (routes []*route) {
	routes = []*route{
		//Item
		{
			"GetItem",
			"GET",
			"/item/{itemID}",
			a.getItem,
		},
		{
			"SearchItem",
			"GET",
			"/item/search",
			a.searchItem,
		},
		{
			"SearchItemByAccount",
			"GET",
			"/item/search/byaccount",
			a.searchItemByAccount,
		},
		{
			"ListItemBySlot",
			"GET",
			"/item/byslot",
			a.listItemBySlot,
		},
		{
			"ListItemByZone",
			"GET",
			"/item/byzone",
			a.listItemByZone,
		},
		{
			"GetItemByZone",
			"GET",
			"/item/byzone/{zoneID}",
			a.getItemByZone,
		},
		{
			"GetItemBySlot",
			"GET",
			"/item/byslot/{slotID}",
			a.getItemBySlot,
		},
		{
			"ListItem",
			"GET",
			"/item",
			a.listItem,
		},
	}
	return
}

func (a *Web) listItem(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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

	items, err := a.itemRepo.List(itemPage.PageSize, itemPage.PageNumber, user)
	if err != nil {
		return
	}

	itemPage.Total, err = a.itemRepo.ListCount(user)
	if err != nil {
		return
	}
	content = Content{
		Site:     site,
		Items:    items,
		ItemPage: itemPage,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("item", tmp)
	}

	return
}

func (a *Web) listItemByZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Zones []*model.Zone
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	zones, err := a.zoneRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:  site,
		Zones: zones,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/listbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("itemlistbyzone", tmp)
	}

	return
}

func (a *Web) getItemByZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		Zone     *model.Zone
		NpcLoots []*model.NpcLoot
	}

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		return
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	zone := &model.Zone{
		ID: zoneID,
	}
	npcLoots, err := a.npcLootRepo.ListByZone(zone, user)
	if err != nil {
		return
	}
	err = a.zoneRepo.Get(zone, user)
	content = Content{
		Site:     site,
		NpcLoots: npcLoots,
		Zone:     zone,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/getbyzone.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("itemlistbyzone", tmp)
	}

	return
}

func (a *Web) listItemByCharacter(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	characterID, err := getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
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
	character := &model.Character{
		ID: characterID,
	}
	err = a.characterRepo.Get(character, user)
	if err != nil {
		return
	}

	inventory, err := a.itemRepo.ListByCharacter(character, user)
	if err != nil {
		return
	}

	itemInventory := map[int]model.Item{}

	for i := range inventory {
		itemInventory[int(inventory[i].SlotID)] = *inventory[i]
	}

	content = Content{
		Site:      site,
		Inventory: itemInventory,
		Character: character,
	}

	fmt.Println(characterID, len(inventory))
	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/listbycharacter.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("item", tmp)
	}

	return
}

func (a *Web) listItemBySlot(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Items []*model.Item
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	//itemCategory := &model.ItemCategory{}

	//	err = a.itemRepo.GetByItemCategory(itemCategory, user)
	//	if err != nil {
	//		return
	//	}

	content = Content{
		Site: site,
		//	Items: items,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/listbyslot.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("item", tmp)
	}

	return
}

func (a *Web) getItemBySlot(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Items  []*model.Item
		SlotID int64
	}

	slotID, err := getIntVar(r, "slotID")
	if err != nil {
		err = errors.Wrap(err, "slotID argument is required")
		return
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	itemCategory := &model.ItemCategory{
		ID: slotID,
	}
	items, err := a.itemRepo.GetByItemCategory(itemCategory, user)
	if err != nil {
		return
	}

	content = Content{
		Site:   site,
		Items:  items,
		SlotID: slotID,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/getbyslot.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("item", tmp)
	}

	return
}

func (a *Web) getItem(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		Item      *model.Item
		Npcs      []*model.Npc
		Merchants []*model.Merchant
		Fishings  []*model.Fishing
		Recipes   []*model.Recipe
	}

	prepContent := &Content{}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		return
	}
	item := &model.Item{
		ID: itemID,
	}
	err = a.itemRepo.Get(item, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	prepContent.Item = item

	npcs, err := a.npcRepo.ListByItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get NPCs based on item")
		return
	}
	prepContent.Npcs = npcs

	merchantEntrys, err := a.merchantEntryRepo.ListByItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get merchants based on item")
		return
	}

	for _, merchantEntry := range merchantEntrys {
		if merchantEntry.MerchantID > 0 {
			merchant := &model.Merchant{
				ID: merchantEntry.MerchantID,
			}

			err = a.merchantRepo.Get(merchant, user)
			if err != nil {
				fmt.Println("Failed to get", merchant.ID, err.Error())
				continue
			}

			merchant.Npcs, err = a.npcRepo.ListByMerchant(merchant, user)
			if err != nil {
				fmt.Println("failed to get npcs for", merchant.ID, err.Error())
				continue
			}

			prepContent.Merchants = append(prepContent.Merchants, merchant)
		}
	}

	fishings, err := a.fishingRepo.GetByItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get fishings based on item")
		return
	}
	for _, fishing := range fishings {
		if fishing.ZoneID > 0 {
			zone := &model.Zone{
				ZoneIDNumber: fishing.ZoneID,
			}
			err = a.zoneRepo.Get(zone, user)
			if err != nil {
				err = errors.Wrap(err, "Failed to get zone based on fishing zoneid")
				return
			}
		}
	}
	prepContent.Fishings = fishings

	recipeEntrys, err := a.recipeEntryRepo.ListByItem(item, user)
	if err != nil {
		err = errors.Wrap(err, "Failed to get recipes based on item")
		return
	}

	recipes := []*model.Recipe{}
	for _, recipeEntry := range recipeEntrys {
		if recipeEntry.RecipeID > 0 {
			recipe := &model.Recipe{
				ID: recipeEntry.RecipeID,
			}
			err = a.recipeRepo.Get(recipe, user)
			if err != nil {
				continue
				//err = errors.Wrap(err, "Failed to get recipe")
				//				//return
			}
			recipes = append(recipes, recipe)
		}
	}
	prepContent.Recipes = recipes

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"
	prepContent.Site = site

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("item", tmp)
	}

	content = prepContent
	return
}

func (a *Web) searchItemByAccount(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site    site
		Items   []*model.Item
		Account []*model.Account
		Search  string
	}

	search := getParam(r, "search")

	var items []*model.Item

	if len(search) > 0 {
		account := &model.Account{
			ID: user.AccountID,
		}
		item := &model.Item{
			Name: search,
		}

		items, err = a.itemRepo.SearchByAccount(item, account, user)
		if err != nil {
			return
		}
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	content = Content{
		Site:   site,
		Items:  items,
		Search: search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/searchbyaccount.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("itemsearchbyaccount", tmp)
	}

	return
}

func (a *Web) searchItem(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site   site
		Items  []*model.Item
		Search string
	}

	search := getParam(r, "search")

	var items []*model.Item

	if len(search) > 0 {
		item := &model.Item{
			Name: search,
		}
		items, err = a.itemRepo.SearchByName(item, user)
		if err != nil {
			return
		}
	}

	site := a.newSite(r)
	site.Page = "item"
	site.Title = "Item"

	content = Content{
		Site:   site,
		Items:  items,
		Search: search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "item/search.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("itemsearch", tmp)
	}

	return
}
