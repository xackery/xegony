package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func accountRoutes() (routes []*route) {
	routes = []*route{
		//Account
		{
			"ListAccount",
			"GET",
			"/account",
			listAccount,
		},
		{
			"GetAccount",
			"GET",
			"/account/{accountID:[0-9]+}",
			getAccount,
		},
	}
	return
}

func listAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		Accounts []*model.Account
	}

	site := newSite(r, user)
	site.Page = "account"
	site.Title = "Account"
	site.Section = "account"
	page := &model.Page{}
	accounts, err := cases.ListAccount(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:     site,
		Accounts: accounts,
	}

	tmp, err = loadTemplate(nil, "body", "account/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}
	return
}

func getAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site    site
		Account *model.Account
	}

	accountID := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}
	account := &model.Account{
		ID: accountID,
	}

	err = cases.GetAccount(account, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := newSite(r, user)
	site.Page = "account"
	site.Title = "Account"
	site.Section = "account"

	content = Content{
		Site:    site,
		Account: account,
	}

	tmp, err = loadTemplate(nil, "body", "account/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
