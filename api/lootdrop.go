package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) lootDropRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *API) getLootDrop(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}
	lootDrop := &model.LootDrop{
		ID: lootDropID,
	}
	err = a.lootDropRepo.Get(lootDrop, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootDrop
	return
}

func (a *API) createLootDrop(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootDrop := &model.LootDrop{}
	err = decodeBody(r, lootDrop)
	if err != nil {
		return
	}
	err = a.lootDropRepo.Create(lootDrop, user)
	if err != nil {
		return
	}
	content = lootDrop
	return
}

func (a *API) deleteLootDrop(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}
	lootDrop := &model.LootDrop{
		ID: lootDropID,
	}
	err = a.lootDropRepo.Delete(lootDrop, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = lootDrop
	return
}

func (a *API) editLootDrop(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	lootDrop := &model.LootDrop{}
	err = decodeBody(r, lootDrop)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	lootDrop.ID = lootDropID

	err = a.lootDropRepo.Edit(lootDrop, user)
	if err != nil {
		return
	}
	content = lootDrop
	return
}

func (a *API) listLootDrop(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	lootDrops, err := a.lootDropRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootDrops
	return
}
