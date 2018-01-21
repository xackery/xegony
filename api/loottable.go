package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) lootTableRoutes() (routes []*route) {
	routes = []*route{}
	return
}
func (a *API) getLootTable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}
	lootTable := &model.LootTable{
		ID: lootTableID,
	}
	err = a.lootTableRepo.Get(lootTable, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootTable
	return
}

func (a *API) createLootTable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootTable := &model.LootTable{}
	err = decodeBody(r, lootTable)
	if err != nil {
		return
	}
	err = a.lootTableRepo.Create(lootTable, user)
	if err != nil {
		return
	}
	content = lootTable
	return
}

func (a *API) deleteLootTable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootTable := &model.LootTable{
		ID: lootTableID,
	}
	err = a.lootTableRepo.Delete(lootTable, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = lootTable
	return
}

func (a *API) editLootTable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootTable := &model.LootTable{}
	err = decodeBody(r, lootTable)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	lootTable.ID = lootTableID
	err = a.lootTableRepo.Edit(lootTable, user)
	if err != nil {
		return
	}
	content = lootTable
	return
}

func (a *API) listLootTable(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	lootTables, err := a.lootTableRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootTables
	return
}
