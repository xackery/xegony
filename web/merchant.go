package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listMerchant(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		Merchants []*model.Merchant
	}

	site := a.newSite(r)
	site.Page = "merchant"
	site.Title = "Merchant"

	pageSize := getIntParam(r, "pageSize")
	pageNumber := getIntParam(r, "pageNumber")

	merchants, err := a.merchantRepo.List(pageSize, pageNumber)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	site.ResultCount, err = a.merchantRepo.ListCount()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, merchant := range merchants {
		merchant.Entrys, _, err = a.merchantEntryRepo.List(merchant.MerchantID)
		if err != nil {
			err = errors.Wrap(err, "merchant entry not found")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
		for _, entry := range merchant.Entrys {
			entry.Item, err = a.itemRepo.Get(entry.ItemID)
			if err != nil {
				err = errors.Wrap(err, "merchant item not found")
				a.writeError(w, r, err, http.StatusBadRequest)
				return
			}
		}
	}
	content := Content{
		Site:      site,
		Merchants: merchants,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchant/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("merchant", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getMerchant(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Merchant *model.Merchant
	}

	if strings.ToLower(getVar(r, "merchantID")) == "search" {
		a.searchMerchant(w, r)
		return
	}

	merchantID, err := getIntVar(r, "merchantID")
	if err != nil {
		err = errors.Wrap(err, "merchantID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	merchant, err := a.merchantRepo.Get(merchantID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	for _, entry := range merchant.Entrys {
		entry.Item, err = a.itemRepo.Get(entry.ItemID)
		if err != nil {
			err = errors.Wrap(err, "merchan item not found")
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "merchant"
	site.Title = "Merchant"

	content := Content{
		Site:     site,
		Merchant: merchant,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchant/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("merchant", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) searchMerchant(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		Merchants []*model.Merchant
		Search    string
	}

	search := getParam(r, "search")

	var merchants []*model.Merchant

	if len(search) > 0 {
		merchants, err = a.merchantRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}

	site := a.newSite(r)
	site.Page = "merchant"
	site.Title = "Merchant"

	content := Content{
		Site:      site,
		Merchants: merchants,
		Search:    search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchant/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("merchantsearch", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
