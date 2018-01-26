package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) merchantRoutes() (routes []*route) {
	routes = []*route{
		//Merchant
		{
			"GetMerchant",
			"GET",
			"/merchant/{merchantID}/details",
			a.getMerchant,
		},
		{
			"ListMerchant",
			"GET",
			"/merchant",
			a.listMerchant,
		},
	}
	return
}

func (a *Web) listMerchant(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site         site
		Merchants    []*model.Merchant
		MerchantPage *model.Page
	}

	site := a.newSite(r)
	site.Page = "merchant"
	site.Title = "Merchant"

	merchantPage := &model.Page{
		Scope: "merchant",
	}
	merchantPage.PageSize = getIntParam(r, "pageSize")
	merchantPage.PageNumber = getIntParam(r, "pageNumber")

	merchants, err := a.merchantRepo.List(merchantPage.PageSize, merchantPage.PageNumber, user)
	if err != nil {
		return
	}
	merchantPage.Total, err = a.merchantRepo.ListCount(user)
	if err != nil {
		return
	}

	for _, merchant := range merchants {

		merchant.Entrys, err = a.merchantEntryRepo.ListByMerchant(merchant, user)
		if err != nil {
			err = errors.Wrap(err, "merchant entry not found")
			return
		}
		for _, entry := range merchant.Entrys {
			entry.Item = &model.Item{
				ID: entry.ItemID,
			}
			err = a.itemRepo.Get(entry.Item, user)
			if err != nil {
				err = errors.Wrap(err, "merchant item not found")
				return
			}
		}
	}
	content = Content{
		Site:      site,
		Merchants: merchants,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchant/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("merchant", tmp)
	}

	return
}

func (a *Web) getMerchant(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		Merchant *model.Merchant
	}

	merchantID, err := getIntVar(r, "merchantID")
	if err != nil {
		err = errors.Wrap(err, "merchantID argument is required")
		return
	}

	merchant := &model.Merchant{
		ID: merchantID,
	}
	err = a.merchantRepo.Get(merchant, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	for _, entry := range merchant.Entrys {
		entry.Item = &model.Item{
			ID: entry.ItemID,
		}
		err = a.itemRepo.Get(entry.Item, user)
		if err != nil {
			err = errors.Wrap(err, "merchan item not found")
			return
		}
	}

	site := a.newSite(r)
	site.Page = "merchant"
	site.Title = "Merchant"

	content = Content{
		Site:     site,
		Merchant: merchant,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "merchant/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("merchant", tmp)
	}

	return
}
