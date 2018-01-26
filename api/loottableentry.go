package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) lootTableEntryRoutes() (routes []*route) {
	routes = []*route{}
	return
}
func (a *API) getLootTableEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}
	lootTableEntry := &model.LootTableEntry{
		LootdropID:  lootDropID,
		LoottableID: lootTableID,
	}
	err = a.lootTableEntryRepo.Get(lootTableEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootTableEntry
	return
}

func (a *API) createLootTableEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableEntry := &model.LootTableEntry{}
	err = decodeBody(r, lootTableEntry)
	if err != nil {
		return
	}
	err = a.lootTableEntryRepo.Create(lootTableEntry, user)
	if err != nil {
		return
	}
	content = lootTableEntry
	return
}

func (a *API) deleteLootTableEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	lootTableEntry := &model.LootTableEntry{
		LoottableID: lootTableID,
		LootdropID:  lootDropID,
	}
	err = a.lootTableEntryRepo.Delete(lootTableEntry, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = lootTableEntry
	return
}

func (a *API) editLootTableEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	lootTableEntry := &model.LootTableEntry{}
	err = decodeBody(r, lootTableEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	lootTableEntry.LoottableID = lootTableID
	lootTableEntry.LootdropID = lootDropID
	err = a.lootTableEntryRepo.Edit(lootTableEntry, user)
	if err != nil {
		return
	}
	content = lootTableEntry
	return
}

func (a *API) listLootTableEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootTableID, err := getIntVar(r, "lootTableID")
	if err != nil {
		err = errors.Wrap(err, "lootTableID argument is required")
		return
	}

	lootTable := &model.LootTable{
		ID: lootTableID,
	}
	lootTableEntrys, err := a.lootTableEntryRepo.ListByLootTable(lootTable, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootTableEntrys
	return
}
