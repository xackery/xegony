package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// LootEntryRequest is a list of parameters used for lootEntry
// swagger:parameters deleteLootEntry editLootEntry getLootEntry
type LootEntryRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"lootID"`
	// EntryID to get information about
	// in: path
	// example: 55091
	EntryID int64 `json:"entryID"`
}

// LootEntryResponse is what endpoints respond with
// swagger:response
type LootEntryResponse struct {
	Loot      *model.Loot      `json:"loot,omitempty"`
	LootEntry *model.LootEntry `json:"lootEntry,omitempty"`
}

// LootEntryCreateRequest is the body parameters for creating an lootEntry
// swagger:parameters createLootEntry
type LootEntryCreateRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"lootID"`
	// LootEntry details to create
	// in: body
	LootEntry *model.LootEntry `json:"lootEntry"`
}

// LootEntryEditRequest is the body parameters for creating an lootEntry
// swagger:parameters editLootEntry
type LootEntryEditRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"LootID"`
	// EntryID to get information about
	// in: path
	// example: 55091
	EntryID int64 `json:"entryID"`
	// LootEntry details to edit
	// in: body
	LootEntry *model.LootEntry `json:"lootEntry"`
}

// LootEntrysRequest is a list of parameters used for lootEntry
// swagger:parameters listLootEntry
type LootEntrysRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"lootID"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// LootEntrysResponse is a general response to a request
// swagger:response
type LootEntrysResponse struct {
	Page       *model.Page      `json:"page,omitempty"`
	Loot       *model.Loot      `json:"loot"`
	LootEntrys model.LootEntrys `json:"lootEntrys,omitempty"`
}

// LootEntrysBySearchRequest is a list of parameters used for lootEntry
// swagger:parameters listLootEntryBySearch
type LootEntrysBySearchRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"lootID"`
	// EntryID is which lootEntry to get information about
	// example: 55091
	// in: query
	EntryID int64 `json:"entryID"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// LootEntrysBySearchResponse is a general response to a request
// swagger:response
type LootEntrysBySearchResponse struct {
	Search     *model.LootEntry `json:"search,omitempty"`
	Page       *model.Page      `json:"page,omitempty"`
	Loot       *model.Loot      `json:"loot,omitempty"`
	LootEntrys model.LootEntrys `json:"lootEntrys,omitempty"`
}

func lootEntryRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /loot/{lootID}/entry loot listLootEntry
		//
		// Lists lootEntrys
		//
		// This will show all available lootEntrys by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: LootEntrysResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLootEntry",
			"GET",
			"/loot/{lootID:[0-9]+}/entry",
			listLootEntry,
		},
		// swagger:route GET /loot/{lootID}/entry/search loot listLootEntryBySearch
		//
		// Search lootEntrys by entryid
		//
		// This will show all available lootEntrys by default.
		//
		//     Consumes:
		//     - application/json
		//
		//     Produces:
		//     - application/json
		//     - application/xml
		//     - application/yaml
		//
		//
		//     Responses:
		//       default: ErrInternal
		//       200: LootEntrysBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLootEntryBySearch",
			"GET",
			"/loot/entry/search",
			listLootEntryBySearch,
		},
		// swagger:route POST /loot/{lootID}/entry/{entryID} loot createLootEntry
		//
		// Create an lootEntry
		//
		// This will create an lootEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateLootEntry",
			"POST",
			"/loot/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			createLootEntry,
		},
		// swagger:route GET /loot/{lootID}/entry/{entryID} loot getLootEntry
		//
		// Get an lootEntry
		//
		// This will get an individual lootEntry available lootEntrys by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: LootEntryResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetLootEntry",
			"GET",
			"/loot/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			getLootEntry,
		},
		// swagger:route PUT /loot/{lootID}/entry/{entryID} loot editLootEntry
		//
		// Edit an lootEntry
		//
		// This will edit an lootEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditLootEntry",
			"PUT",
			"/loot/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			editLootEntry,
		},
		// swagger:route DELETE /loot/{lootID}/entry/{entryID} loot deleteLootEntry
		//
		// Delete an lootEntry
		//
		// This will delete an lootEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"DeleteLootEntry",
			"DELETE",
			"/loot/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			deleteLootEntry,
		},
	}
	return
}

func getLootEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootEntryRequest{
		LootID:  getIntVar(r, "lootID"),
		EntryID: getIntVar(r, "entryID"),
	}

	loot := &model.Loot{
		ID: request.LootID,
	}

	err = cases.GetLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	lootEntry := &model.LootEntry{
		LootID:     request.LootID,
		LootDropID: request.EntryID,
	}

	err = cases.GetLootEntry(loot, lootEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &LootEntryResponse{
		Loot:      loot,
		LootEntry: lootEntry,
	}
	content = response
	return
}

func createLootEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootEntryCreateRequest{
		LootID: getIntVar(r, "lootID"),
	}

	loot := &model.Loot{
		ID: request.LootID,
	}

	err = cases.GetLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	lootEntry := &model.LootEntry{}
	err = decodeBody(r, lootEntry)
	if err != nil {
		return
	}

	err = cases.CreateLootEntry(loot, lootEntry, user)
	if err != nil {
		return
	}
	response := &LootEntryResponse{
		Loot:      loot,
		LootEntry: lootEntry,
	}
	content = response
	return
}

func deleteLootEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootEntryRequest{
		LootID:  getIntVar(r, "lootID"),
		EntryID: getIntVar(r, "entryID"),
	}

	loot := &model.Loot{
		ID: request.LootID,
	}

	err = cases.GetLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	lootEntry := &model.LootEntry{
		LootID:     request.LootID,
		LootDropID: request.EntryID,
	}

	err = cases.DeleteLootEntry(lootEntry, loot, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
	}
	err = &model.ErrNoContent{}
	return
}

func editLootEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootEntryEditRequest{
		LootID:  getIntVar(r, "lootID"),
		EntryID: getIntVar(r, "entryID"),
	}

	loot := &model.Loot{
		ID: request.LootID,
	}

	err = cases.GetLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	lootEntry := &model.LootEntry{}

	err = decodeBody(r, lootEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	lootEntry.LootDropID = request.EntryID
	lootEntry.LootID = request.LootID

	err = cases.EditLootEntry(loot, lootEntry, user)
	if err != nil {
		return
	}
	response := &LootEntryResponse{
		Loot:      loot,
		LootEntry: lootEntry,
	}
	content = response
	return
}

func listLootEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootEntrysRequest{
		LootID: getIntVar(r, "lootID"),
	}

	loot := &model.Loot{
		ID: request.LootID,
	}
	err = cases.GetLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	lootEntrys, err := cases.ListLootEntry(page, loot, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &LootEntrysResponse{
		Page:       page,
		Loot:       loot,
		LootEntrys: lootEntrys,
	}
	content = response
	return
}

func listLootEntryBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootEntrysBySearchRequest{
		LootID: getIntVar(r, "lootID"),
	}

	loot := &model.Loot{
		ID: request.LootID,
	}

	err = cases.GetLoot(loot, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	lootEntry := &model.LootEntry{
		LootID: request.LootID,
	}
	lootEntry.LootDropID = getIntQuery(r, "entryID")

	lootEntrys, err := cases.ListLootEntryBySearch(page, loot, lootEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &LootEntrysBySearchResponse{
		Page:       page,
		Loot:       loot,
		LootEntrys: lootEntrys,
		Search:     lootEntry,
	}
	content = response
	return
}
