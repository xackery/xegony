package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       Site
		Characters []*model.Character
	}

	site := a.NewSite(r)
	site.Page = "character"
	site.Title = "Character"

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

func (a *Web) SearchCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       Site
		Characters []*model.Character
		Search     string
	}

	search := getParam(r, "search")

	site := a.NewSite(r)
	site.Page = "character"
	site.Title = "Character"
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

func (a *Web) ListCharacterByRanking(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       Site
		Characters []*model.Character
	}

	site := a.NewSite(r)
	site.Page = "ranking"
	site.Title = "Character"

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

func (a *Web) ListCharacterByOnline(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       Site
		Characters []*model.Character
	}

	site := a.NewSite(r)
	site.Page = "ranking"
	site.Title = "Character"

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

func (a *Web) ListCharacterByAccount(w http.ResponseWriter, r *http.Request) {
	var err error

	accountId, err := getIntVar(r, "accountId")
	if err != nil {
		err = errors.Wrap(err, "accountId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	type Content struct {
		Site       Site
		Characters []*model.Character
	}

	site := a.NewSite(r)
	site.Page = "ranking"
	site.Title = "Character"

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

func (a *Web) GetCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      Site
		Character *model.Character
	}
	if (strings.ToLower(getVar(r, "characterId"))) == "byranking" {
		a.ListCharacterByRanking(w, r)
		return
	}

	if (strings.ToLower(getVar(r, "characterId"))) == "byonline" {
		a.ListCharacterByOnline(w, r)
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

	site := a.NewSite(r)
	site.Page = "character"
	site.Title = "Character"

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
