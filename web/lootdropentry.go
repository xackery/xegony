package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) lootDropEntryRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *Web) listLootDropEntry(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

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
		return
	}

	lootDrop := &model.LootDrop{
		ID: lootDropID,
	}
	lootDropEntrys, err := a.lootDropEntryRepo.ListByLootDrop(lootDrop, user)
	if err != nil {
		return
	}
	content = Content{
		Site:           site,
		LootDropEntrys: lootDropEntrys,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdropentry/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("lootDropEntry", tmp)
	}

	return
}

func (a *Web) getLootDropEntry(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site          site
		LootDropEntry *model.LootDropEntry
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		return
	}

	lootDropEntry := &model.LootDropEntry{
		ItemID:     itemID,
		LootdropID: lootDropID,
	}

	err = a.lootDropEntryRepo.Get(lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "lootDropEntry"
	site.Title = "LootDropEntry"
	site.Section = "lootDropEntry"

	content = Content{
		Site:          site,
		LootDropEntry: lootDropEntry,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "lootdropentry/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("lootdrop", tmp)
	}

	return
}
