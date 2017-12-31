package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) getNpc(w http.ResponseWriter, r *http.Request) {

	id, err := getIntVar(r, "npcId")
	if err != nil {
		err = errors.Wrap(err, "npcId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	npc, err := a.npcRepo.Get(id)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, npc, http.StatusOK)
	return
}

func (a *Api) createNpc(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	npc := &model.Npc{}
	err = decodeBody(r, npc)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.npcRepo.Create(npc)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, npc, http.StatusCreated)
	return
}

func (a *Api) deleteNpc(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "npcId")
	if err != nil {
		err = errors.Wrap(err, "npcId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.npcRepo.Delete(id)
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

func (a *Api) editNpc(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	id, err := getIntVar(r, "npcId")
	if err != nil {
		err = errors.Wrap(err, "npcId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	npc := &model.Npc{}
	err = decodeBody(r, npc)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.npcRepo.Edit(id, npc)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, npc, http.StatusOK)
	return
}

func (a *Api) listNpc(w http.ResponseWriter, r *http.Request) {
	npcs, err := a.npcRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, npcs, http.StatusOK)
	return
}
