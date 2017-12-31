package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		Characters []*model.Character
	}

	site := a.newSite(r)
	site.Page = "charactersearch"
	site.Title = "Character"
	site.Section = "character"

	characters, err := a.characterRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:       site,
		Characters: characters,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("character", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) searchCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

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
		characters, err = a.characterRepo.Search(search)
		if err != nil {
			a.writeError(w, r, err, http.StatusBadRequest)
			return
		}
	}
	content := Content{
		Site:       site,
		Characters: characters,
		Search:     search,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/search.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("character", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listCharacterByRanking(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		Characters []*model.Character
	}

	site := a.newSite(r)
	site.Page = "characterbyranking"
	site.Title = "Character"
	site.Section = "character"

	characters, err := a.characterRepo.ListByRanking()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:       site,
		Characters: characters,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/listbyranking.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("character", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listCharacterByOnline(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       site
		Characters []*model.Character
	}

	site := a.newSite(r)
	site.Page = "characterbyonline"
	site.Title = "Character"
	site.Section = "character"

	characters, err := a.characterRepo.ListByOnline()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:       site,
		Characters: characters,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/listbyonline.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("character", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) listCharacterByAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	accountId, err := getIntVar(r, "accountId")
	if err != nil {
		err = errors.Wrap(err, "accountId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
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

	characters, err := a.characterRepo.ListByAccount(accountId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:       site,
		Characters: characters,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("character", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		Character *model.Character
	}
	if (strings.ToLower(getVar(r, "characterId"))) == "byranking" {
		a.listCharacterByRanking(w, r)
		return
	}

	if (strings.ToLower(getVar(r, "characterId"))) == "byonline" {
		a.listCharacterByOnline(w, r)
		return
	}

	id, err := getIntVar(r, "characterId")
	if err != nil {
		err = errors.Wrap(err, "characterId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	character, err := a.characterRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "charactersearch"
	site.Title = "Character"
	site.Section = "character"

	content := Content{
		Site:      site,
		Character: character,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "character/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("character", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
