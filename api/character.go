package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) getCharacter(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "characterId")
	if err != nil {
		err = errors.Wrap(err, "characterId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	character, err := a.characterRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, character, http.StatusOK)
	return
}

func (a *Api) createCharacter(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	character := &model.Character{}
	err = decodeBody(r, character)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.characterRepo.Create(character)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, character, http.StatusCreated)
	return
}

func (a *Api) deleteCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "characterId")
	if err != nil {
		err = errors.Wrap(err, "characterId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.characterRepo.Delete(id)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			writeData(w, r, nil, http.StatusNotModified)
			return
		default:
			err = errors.Wrap(err, "Request failed")
			writeError(w, r, err, http.StatusInternalServerError)
		}
		return
	}
	writeData(w, r, nil, http.StatusNoContent)
	return
}

func (a *Api) editCharacter(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "characterId")
	if err != nil {
		err = errors.Wrap(err, "characterId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	character := &model.Character{}
	err = decodeBody(r, character)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.characterRepo.Edit(id, character)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, character, http.StatusOK)
	return
}

func (a *Api) listCharacter(w http.ResponseWriter, r *http.Request) {
	characters, err := a.characterRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, characters, http.StatusOK)
	return
}
