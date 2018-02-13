package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// LootDropEntryRequest is a list of parameters used for lootDropEntry
// swagger:parameters deleteLootDropEntry editLootDropEntry getLootDropEntry
type LootDropEntryRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"lootID"`
	// EntryID to get information about
	// in: path
	// example: 55091
	EntryID int64 `json:"entryID"`
}

// LootDropEntryResponse is what endpoints respond with
// swagger:response
type LootDropEntryResponse struct {
	LootDrop      *model.LootDrop      `json:"lootDrop,omitempty"`
	LootDropEntry *model.LootDropEntry `json:"lootDropEntry,omitempty"`
}

// LootDropEntryCreateRequest is the body parameters for creating an lootDropEntry
// swagger:parameters createLootDropEntry
type LootDropEntryCreateRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"lootID"`
	// LootDropEntry details to create
	// in: body
	LootDropEntry *model.LootDropEntry `json:"lootDropEntry"`
}

// LootDropEntryEditRequest is the body parameters for creating an lootDropEntry
// swagger:parameters editLootDropEntry
type LootDropEntryEditRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"LootID"`
	// EntryID to get information about
	// in: path
	// example: 55091
	EntryID int64 `json:"entryID"`
	// LootDropEntry details to edit
	// in: body
	LootDropEntry *model.LootDropEntry `json:"lootDropEntry"`
}

// LootDropEntrysRequest is a list of parameters used for lootDropEntry
// swagger:parameters listLootDropEntry
type LootDropEntrysRequest struct {
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

// LootDropEntrysResponse is a general response to a request
// swagger:response
type LootDropEntrysResponse struct {
	Page           *model.Page          `json:"page,omitempty"`
	LootDrop       *model.LootDrop      `json:"lootDrop"`
	LootDropEntrys model.LootDropEntrys `json:"lootDropEntrys,omitempty"`
}

// LootDropEntrysBySearchRequest is a list of parameters used for lootDropEntry
// swagger:parameters listLootDropEntryBySearch
type LootDropEntrysBySearchRequest struct {
	// LootID to get information about
	// in: path
	// example: 1
	LootID int64 `json:"lootID"`
	// EntryID is which lootDropEntry to get information about
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

// LootDropEntrysBySearchResponse is a general response to a request
// swagger:response
type LootDropEntrysBySearchResponse struct {
	Search         *model.LootDropEntry `json:"search,omitempty"`
	Page           *model.Page          `json:"page,omitempty"`
	LootDrop       *model.LootDrop      `json:"lootDrop,omitempty"`
	LootDropEntrys model.LootDropEntrys `json:"lootDropEntrys,omitempty"`
}

func lootDropEntryRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /loot/drop/{lootID}/entry loot listLootDropEntry
		//
		// Lists lootDropEntrys
		//
		// This will show all available lootDropEntrys by default.
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
		//       200: LootDropEntrysResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLootDropEntry",
			"GET",
			"/loot/drop/{lootID:[0-9]+}/entry",
			listLootDropEntry,
		},
		// swagger:route GET /loot/drop/{lootID}/entry/search loot listLootDropEntryBySearch
		//
		// Search lootDropEntrys by entryid
		//
		// This will show all available lootDropEntrys by default.
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
		//       200: LootDropEntrysBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLootDropEntryBySearch",
			"GET",
			"/loot/drop/entry/search",
			listLootDropEntryBySearch,
		},
		// swagger:route POST /loot/drop/{lootID}/entry/{entryID} loot createLootDropEntry
		//
		// Create an lootDropEntry
		//
		// This will create an lootDropEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootDropEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateLootDropEntry",
			"POST",
			"/loot/drop/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			createLootDropEntry,
		},
		// swagger:route GET /loot/drop/{lootID}/entry/{entryID} loot getLootDropEntry
		//
		// Get an lootDropEntry
		//
		// This will get an individual lootDropEntry available lootDropEntrys by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: LootDropEntryResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetLootDropEntry",
			"GET",
			"/loot/drop/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			getLootDropEntry,
		},
		// swagger:route PUT /loot/drop/{lootID}/entry/{entryID} loot editLootDropEntry
		//
		// Edit an lootDropEntry
		//
		// This will edit an lootDropEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootDropEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditLootDropEntry",
			"PUT",
			"/loot/drop/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			editLootDropEntry,
		},
		// swagger:route DELETE /loot/drop/{lootID}/entry/{entryID} loot deleteLootDropEntry
		//
		// Delete an lootDropEntry
		//
		// This will delete an lootDropEntry
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
			"DeleteLootDropEntry",
			"DELETE",
			"/loot/drop/{lootID:[0-9]+}/entry/{entryID:[0-9]+}",
			deleteLootDropEntry,
		},
	}
	return
}

func getLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropEntryRequest{
		LootID:  getIntVar(r, "lootID"),
		EntryID: getIntVar(r, "entryID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.LootID,
	}

	err = cases.GetLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDrop")
		return
	}

	lootDropEntry := &model.LootDropEntry{
		ItemID:     request.EntryID,
		LootDropID: request.LootID,
	}

	err = cases.GetLootDropEntry(lootDrop, lootDropEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &LootDropEntryResponse{
		LootDrop:      lootDrop,
		LootDropEntry: lootDropEntry,
	}
	content = response
	return
}

func createLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropEntryCreateRequest{
		LootID: getIntVar(r, "lootID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.LootID,
	}

	err = cases.GetLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	lootDropEntry := &model.LootDropEntry{}
	err = decodeBody(r, lootDropEntry)
	if err != nil {
		return
	}

	err = cases.CreateLootDropEntry(lootDrop, lootDropEntry, user)
	if err != nil {
		return
	}
	response := &LootDropEntryResponse{
		LootDrop:      lootDrop,
		LootDropEntry: lootDropEntry,
	}
	content = response
	return
}

func deleteLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropEntryRequest{
		LootID:  getIntVar(r, "lootID"),
		EntryID: getIntVar(r, "entryID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.LootID,
	}

	err = cases.GetLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	lootDropEntry := &model.LootDropEntry{
		ItemID:     request.EntryID,
		LootDropID: request.LootID,
	}

	err = cases.DeleteLootDropEntry(lootDropEntry, lootDrop, user)
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

func editLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropEntryEditRequest{
		LootID:  getIntVar(r, "lootID"),
		EntryID: getIntVar(r, "entryID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.LootID,
	}

	err = cases.GetLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get loot")
		return
	}

	lootDropEntry := &model.LootDropEntry{}

	err = decodeBody(r, lootDropEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	lootDropEntry.ItemID = request.EntryID
	lootDropEntry.LootDropID = request.LootID

	err = cases.EditLootDropEntry(lootDrop, lootDropEntry, user)
	if err != nil {
		return
	}
	response := &LootDropEntryResponse{
		LootDrop:      lootDrop,
		LootDropEntry: lootDropEntry,
	}
	content = response
	return
}

func listLootDropEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropEntrysRequest{
		LootID: getIntVar(r, "lootID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.LootID,
	}
	err = cases.GetLootDrop(lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get lootDrop")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	lootDropEntrys, err := cases.ListLootDropEntry(page, lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &LootDropEntrysResponse{
		Page:           page,
		LootDrop:       lootDrop,
		LootDropEntrys: lootDropEntrys,
	}
	content = response
	return
}

func listLootDropEntryBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropEntrysBySearchRequest{
		LootID: getIntVar(r, "lootID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.LootID,
	}

	err = cases.GetLootDrop(lootDrop, user)
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

	lootDropEntry := &model.LootDropEntry{
		LootDropID: request.LootID,
	}
	lootDropEntry.LootDropID = getIntQuery(r, "entryID")

	lootDropEntrys, err := cases.ListLootDropEntryBySearch(page, lootDrop, lootDropEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &LootDropEntrysBySearchResponse{
		Page:           page,
		LootDrop:       lootDrop,
		LootDropEntrys: lootDropEntrys,
		Search:         lootDropEntry,
	}
	content = response
	return
}
