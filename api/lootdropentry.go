package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) lootDropEntryRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *API) getLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		return
	}

	lootDropEntry := &model.LootDropEntry{
		ItemID:     itemID,
		LootdropID: lootDropID,
	}
	err = a.lootDropEntryRepo.Get(lootDropEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootDropEntry
	return
}

func (a *API) createLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropEntry := &model.LootDropEntry{}
	err = decodeBody(r, lootDropEntry)
	if err != nil {
		return
	}
	err = a.lootDropEntryRepo.Create(lootDropEntry, user)
	if err != nil {
		return
	}
	content = lootDropEntry
	return
}

func (a *API) deleteLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		return
	}

	lootDropEntry := &model.LootDropEntry{
		LootdropID: lootDropID,
		ItemID:     itemID,
	}
	err = a.lootDropEntryRepo.Delete(lootDropEntry, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = lootDropEntry
	return
}

func (a *API) editLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	itemID, err := getIntVar(r, "itemID")
	if err != nil {
		err = errors.Wrap(err, "itemID argument is required")
		return
	}

	lootDropEntry := &model.LootDropEntry{}
	err = decodeBody(r, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	lootDropEntry.ItemID = itemID
	lootDropEntry.LootdropID = lootDropID

	err = a.lootDropEntryRepo.Edit(lootDropEntry, user)
	if err != nil {
		return
	}
	content = lootDropEntry
	return
}

func (a *API) listLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	lootDropID, err := getIntVar(r, "lootDropID")
	if err != nil {
		err = errors.Wrap(err, "lootDropID argument is required")
		return
	}

	lootDrop := &model.LootDrop{
		ID: lootDropID,
	}
	lootDropEntrys, err := a.lootDropEntryRepo.ListByLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = lootDropEntrys
	return
}
