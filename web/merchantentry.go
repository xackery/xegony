package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) merchantEntryRoutes() (routes []*route) {
	routes = []*route{
		//MerchantEntry
		{
			"GetMerchantEntry",
			"GET",
			"/merchant/{merchantID}/{itemID}",
			a.getMerchantEntry,
		},
		{
			"ListMerchantEntry",
			"GET",
			"/merchant/{merchantID}",
			a.listMerchantEntry,
		},
	}
	return
}

func (a *Web) listMerchantEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	merchantID, err := getIntVar(r, "merchantID")
	if err != nil {
		err = errors.Wrap(err, "merchantEntryID argument is required")
		return
	}

	type Content struct {
		Site     site
		Merchant *model.Merchant
	}

	site := a.newSite(r)
	site.Page = "merchantentry"
	site.Title = "merchantentry"
	site.Section = "merchantentry"

	merchant := &model.Merchant{
		ID: merchantID,
	}
	err = a.merchantRepo.Get(merchant, user)
	if err != nil {
		return
	}

	merchant.Npcs, err = a.npcRepo.ListByMerchant(merchant, user)
	if err != nil {
		return
	}

	merchant.Entrys, err = a.merchantEntryRepo.ListByMerchant(merchant, user)
	if err != nil {
		return
	}

	for _, entry := range merchant.Entrys {
		entry.Item = &model.Item{
			ID: entry.ItemID,
		}
		err = a.itemRepo.Get(entry.Item, user)
		if err != nil {
			err = errors.Wrap(err, "itemID get is required")
			return
		}
	}

	content = Content{
		Site:     site,
		Merchant: merchant,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchantentry/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("merchantentry", tmp)
	}

	return
}

func (a *Web) getMerchantEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site          site
		MerchantEntry *model.MerchantEntry
	}

	merchantID, err := getIntVar(r, "merchantID")
	if err != nil {
		err = errors.Wrap(err, "merchantEntryID argument is required")
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		return
	}
	merchantEntry := &model.MerchantEntry{
		MerchantID: merchantID,
		ItemID:     itemID,
	}
	err = a.merchantEntryRepo.Get(merchantEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "merchantentry"
	site.Title = "merchantentry"
	site.Section = "merchantentry"

	content = Content{
		Site:          site,
		MerchantEntry: merchantEntry,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchantentry/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("merchantentry", tmp)
	}

	return
}
