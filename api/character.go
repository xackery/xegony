package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) characterRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateCharacter",
			"POST",
			"/character",
			a.createCharacter,
		},
		{
			"DeleteCharacter",
			"DELETE",
			"/character/{characterID:[0-9]+}",
			a.deleteCharacter,
		},
		{
			"EditCharacter",
			"PUT",
			"/character/{characterID:[0-9]+}",
			a.editCharacter,
		},
		{
			"GetCharacter",
			"GET",
			"/character/{characterID:[0-9]+}",
			a.getCharacter,
		},
		{
			"GetCharacterByName",
			"GET",
			"/character/byname/{name:[a-zA-Z]+}",
			a.getCharacterByName,
		},
		{
			"ListCharacter",
			"GET",
			"/character",
			a.listCharacter,
		},
	}
	return
}

func (a *API) getCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

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
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = character
	return
}

func (a *API) getCharacterByName(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	name := getVar(r, "name")

	character := &model.Character{
		Name: name,
	}

	err = a.characterRepo.GetByName(character, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = character
	return
}

func (a *API) createCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	character := &model.Character{}
	err = decodeBody(r, character)
	if err != nil {
		return
	}
	err = a.characterRepo.Create(character, user)
	if err != nil {
		return
	}
	content = character
	return
}

func (a *API) deleteCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	characterID, err := getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "characterID argument is required")
		return
	}

	character := &model.Character{
		ID: characterID,
	}

	err = a.characterRepo.Delete(character, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = character
	return
}

func (a *API) editCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	characterID, err := getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "characterID argument is required")
		return
	}

	character := &model.Character{}
	err = decodeBody(r, character)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	character.ID = characterID

	err = a.characterRepo.Edit(character, user)
	if err != nil {
		return
	}
	content = character
	return
}

func (a *API) listCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	characters, err := a.characterRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = characters
	return
}
