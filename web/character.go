package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) characterRoutes() (routes []*route) {
	routes = []*route{
		{
			"SearchCharacter",
			"GET",
			"/character/search",
			a.searchCharacter,
		},
		{
			"SearchCharacter",
			"GET",
			"/character/search/{search:[a-zA-Z]+}",
			a.searchCharacter,
		},
		{
			"GetCharacter",
			"GET",
			"/character/{characterID:[0-9]+}",
			a.getCharacter,
		},
		{
			"ListCharacter",
			"GET",
			"/character",
			a.listCharacter,
		},
		{
			"ListCharacter",
			"GET",
			"/character/{characterID:[0-9]+}/inventory",
			a.listItemByCharacter,
		},
		{
			"ListCharacterByRanking",
			"GET",
			"/character/byranking",
			a.listCharacterByRanking,
		},
		{
			"ListCharacterByOnline",
			"GET",
			"/character/byonline",
			a.listCharacterByOnline,
		},
		{
			"ListCharacterByAccount",
			"GET",
			"/character/byaccount/{accountID:[0-9]+}",
			a.listCharacterByAccount,
		},
		{
			"ListCharacterByRanking",
			"GET",
			"/ranking",
			a.listCharacterByRanking,
		},
	}
	return
}

func (a *Web) listCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		Characters []*model.Character
	}

	site := a.newSite(r)
	site.Page = "charactersearch"
	site.Title = "Character"
	site.Section = "character"

	characters, err := a.characterRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:       site,
		Characters: characters,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("character", tmp)
	}

	return
}

func (a *Web) searchCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		Characters []*model.Character
		Search     string
	}

	search := getParam(r, "search")

	site := a.newSite(r)
	site.Page = "charactersearch"
	site.Title = "Character"
	site.Section = "character"
	var characters []*model.Character

	if len(search) > 0 {
		character := &model.Character{
			Name: search,
		}
		characters, err = a.characterRepo.SearchByName(character, user)
		if err != nil {
			return
		}
	}
	content = Content{
		Site:       site,
		Characters: characters,
		Search:     search,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/search.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("character", tmp)
	}

	return
}

func (a *Web) listCharacterByRanking(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		Characters []*model.Character
	}

	site := a.newSite(r)
	site.Page = "characterbyranking"
	site.Title = "Character"
	site.Section = "character"

	characters, err := a.characterRepo.ListByRanking(user)
	if err != nil {
		return
	}
	content = Content{
		Site:       site,
		Characters: characters,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/listbyranking.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("character", tmp)
	}

	return
}

func (a *Web) listCharacterByOnline(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		Characters []*model.Character
	}

	site := a.newSite(r)
	site.Page = "characterbyonline"
	site.Title = "Character"
	site.Section = "character"

	characters, err := a.characterRepo.ListByOnline(user)
	if err != nil {
		return
	}
	content = Content{
		Site:       site,
		Characters: characters,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/listbyonline.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("character", tmp)
	}

	return
}

func (a *Web) listCharacterByAccount(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	accountID, err := getIntVar(r, "accountID")
	if err != nil {
		err = errors.Wrap(err, "accountID argument is required")
		return
	}

	type Content struct {
		Site       site
		Characters []*model.Character
	}

	site := a.newSite(r)
	site.Page = "charactersearch"
	site.Title = "Character"
	site.Section = "character"

	account := &model.Account{
		ID: accountID,
	}
	characters, err := a.characterRepo.ListByAccount(account, user)
	if err != nil {
		return
	}
	content = Content{
		Site:       site,
		Characters: characters,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("character", tmp)
	}

	return
}

func (a *Web) getCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		Character *model.Character
	}

	characterID, err := getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "characterID argument is required")
		return
	}

	character := &model.Character{
		ID: characterID,
	}
	err = a.characterRepo.Get(character, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "charactersearch"
	site.Title = "Character"
	site.Section = "character"

	content = Content{
		Site:      site,
		Character: character,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("character", tmp)
	}

	return
}
