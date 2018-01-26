package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) npcRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *API) getNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		return
	}
	npc := &model.Npc{
		ID: npcID,
	}
	err = a.npcRepo.Get(npc, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = npc
	return
}

func (a *API) createNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	npc := &model.Npc{}
	err = decodeBody(r, npc)
	if err != nil {
		return
	}
	err = a.npcRepo.Create(npc, user)
	if err != nil {
		return
	}
	content = npc
	return
}

func (a *API) deleteNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		return
	}

	npc := &model.Npc{
		ID: npcID,
	}
	err = a.npcRepo.Delete(npc, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = npc
	return
}

func (a *API) editNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
		return
	}

	npc := &model.Npc{}
	err = decodeBody(r, npc)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	npc.ID = npcID
	err = a.npcRepo.Edit(npc, user)
	if err != nil {
		return
	}
	content = npc
	return
}

func (a *API) listNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	npcs, err := a.npcRepo.List(20, 0, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = npcs
	return
}
