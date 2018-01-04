package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listMerchantEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	merchantID, err := getIntVar(r, "merchantID")
	if err != nil {
		err = errors.Wrap(err, "merchantEntryID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
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

	merchant, err := a.merchantRepo.Get(merchantID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	merchant.Npcs, err = a.npcRepo.ListByMerchant(merchantID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	merchant.Entrys, _, err = a.merchantEntryRepo.List(merchantID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, entry := range merchant.Entrys {
		entry.Item, err = a.itemRepo.Get(entry.ItemID)
		if err != nil {
			err = errors.Wrap(err, "itemID get is required")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	content := Content{
		Site:     site,
		Merchant: merchant,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchantentry/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("merchantentry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getMerchantEntry(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site          site
		MerchantEntry *model.MerchantEntry
	}

	merchantID, err := getIntVar(r, "merchantID")
	if err != nil {
		err = errors.Wrap(err, "merchantEntryID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	if strings.ToLower(getVar(r, "itemID")) == "details" {
		a.getMerchant(w, r)
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	merchantEntry, _, err := a.merchantEntryRepo.Get(merchantID, itemID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "merchantentry"
	site.Title = "merchantentry"
	site.Section = "merchantentry"

	content := Content{
		Site:          site,
		MerchantEntry: merchantEntry,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchantentry/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("merchantentry", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
