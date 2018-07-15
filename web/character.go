package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

func characterRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListCharacter",
			"GET",
			"/character",
			listCharacter,
		},
		{
			"ListCharacterZone",
			"GET",
			"/character/zone",
			listCharacterZone,
		},
		{
			"ListCharacterZone",
			"GET",
			"/character/zone",
			listCharacterZone,
		},
		{
			"GetCharacter",
			"GET",
			"/character/{characterID:[0-9]+}",
			getCharacter,
		},
	}
	return
}

func listCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site       site
		Page       *model.Page
		Characters []*model.Character
	}

	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}
	characters, err := cases.ListCharacter(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:       site,
		Characters: characters,
		Page:       page,
	}

	tmp, err = loadTemplate(nil, "body", "character/list.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func listCharacterZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Page  *model.Page
		Zones []*model.Zone
	}

	site := newSite(r, user)
	page := &model.Page{
		Limit: 10,
	}
	zones, err := cases.ListZone(page, user)
	if err != nil {
		return
	}

	content = Content{
		Site:  site,
		Zones: zones,
		Page:  page,
	}

	tmp, err = loadTemplate(nil, "body", "character/list_zone.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}

func getCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		Character *model.Character
		Loot      *model.Loot
		Items     []*model.Item
		ItemPage  *model.Page
	}

	characterID := getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "characterID argument is required")
		return
	}
	character := &model.Character{
		ID: characterID,
	}

	err = cases.GetCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := newSite(r, user)
	site.Page = "character"
	site.Title = "Character"
	site.Section = "character"

	content = Content{
		Site:      site,
		Character: character,
	}

	//loot.Entries[0].DropEntrys[0].Item

	tmp, err = loadTemplate(nil, "body", "character/get.tpl")
	if err != nil {
		return
	}
	tmp, err = loadStandardTemplate(tmp)
	if err != nil {
		return
	}

	return
}
