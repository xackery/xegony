package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site       Site
		Characters []*model.Character
	}

	site := a.NewSite()
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
		tmp, err = a.loadTemplate(nil, "body", "listcharacter.tpl")
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

	site := a.NewSite()
	site.Page = "character"
	site.Title = "Character"

	content := Content{
		Site:      site,
		Character: character,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "getcharacter.tpl")
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
