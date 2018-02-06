package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpawnEntryRequest is a list of parameters used for spawnEntry
// swagger:parameters deleteSpawnEntry editSpawnEntry getSpawnEntry
type SpawnEntryRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"spawnID"`
	// EntryID to get information about
	// in: path
	// example: 55091
	EntryID int64 `json:"entryID"`
}

// SpawnEntryResponse is what endpoints respond with
// swagger:response
type SpawnEntryResponse struct {
	Spawn      *model.Spawn      `json:"spawn,omitempty"`
	SpawnEntry *model.SpawnEntry `json:"spawnEntry,omitempty"`
}

// SpawnEntryCreateRequest is the body parameters for creating an spawnEntry
// swagger:parameters createSpawnEntry
type SpawnEntryCreateRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"spawnID"`
	// SpawnEntry details to create
	// in: body
	SpawnEntry *model.SpawnEntry `json:"spawnEntry"`
}

// SpawnEntryEditRequest is the body parameters for creating an spawnEntry
// swagger:parameters editSpawnEntry
type SpawnEntryEditRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"SpawnID"`
	// EntryID to get information about
	// in: path
	// example: 55091
	EntryID int64 `json:"entryID"`
	// SpawnEntry details to edit
	// in: body
	SpawnEntry *model.SpawnEntry `json:"spawnEntry"`
}

// SpawnEntrysRequest is a list of parameters used for spawnEntry
// swagger:parameters listSpawnEntry
type SpawnEntrysRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"spawnID"`
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

// SpawnEntrysResponse is a general response to a request
// swagger:response
type SpawnEntrysResponse struct {
	Page        *model.Page       `json:"page,omitempty"`
	Spawn       *model.Spawn      `json:"spawn"`
	SpawnEntrys model.SpawnEntrys `json:"spawnEntrys,omitempty"`
}

// SpawnEntrysBySearchRequest is a list of parameters used for spawnEntry
// swagger:parameters listSpawnEntryBySearch
type SpawnEntrysBySearchRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"spawnID"`
	// EntryID is which spawnEntry to get information about
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

// SpawnEntrysBySearchResponse is a general response to a request
// swagger:response
type SpawnEntrysBySearchResponse struct {
	Search      *model.SpawnEntry `json:"search,omitempty"`
	Page        *model.Page       `json:"page,omitempty"`
	Spawn       *model.Spawn      `json:"spawn,omitempty"`
	SpawnEntrys model.SpawnEntrys `json:"spawnEntrys,omitempty"`
}

func spawnEntryRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spawn/{spawnID}/entry spawn listSpawnEntry
		//
		// Lists spawnEntrys
		//
		// This will show all available spawnEntrys by default.
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
		//       200: SpawnEntrysResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry",
			listSpawnEntry,
		},
		// swagger:route GET /spawn/{spawnID}/entry/search spawn listSpawnEntryBySearch
		//
		// Search spawnEntrys by entryid
		//
		// This will show all available spawnEntrys by default.
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
		//       200: SpawnEntrysBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawnEntryBySearch",
			"GET",
			"/spawn/entry/search",
			listSpawnEntryBySearch,
		},
		// swagger:route POST /spawn/{spawnID}/entry/{entryID} spawn createSpawnEntry
		//
		// Create an spawnEntry
		//
		// This will create an spawnEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpawnEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpawnEntry",
			"POST",
			"/spawn/{spawnID:[0-9]+}/entry/{entryID:[0-9]+}",
			createSpawnEntry,
		},
		// swagger:route GET /spawn/{spawnID}/entry/{entryID} spawn getSpawnEntry
		//
		// Get an spawnEntry
		//
		// This will get an individual spawnEntry available spawnEntrys by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpawnEntryResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry/{entryID:[0-9]+}",
			getSpawnEntry,
		},
		// swagger:route PUT /spawn/{spawnID}/entry/{entryID} spawn editSpawnEntry
		//
		// Edit an spawnEntry
		//
		// This will edit an spawnEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpawnEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpawnEntry",
			"PUT",
			"/spawn/{spawnID:[0-9]+}/entry/{entryID:[0-9]+}",
			editSpawnEntry,
		},
		// swagger:route DELETE /spawn/{spawnID}/entry/{entryID} spawn deleteSpawnEntry
		//
		// Delete an spawnEntry
		//
		// This will delete an spawnEntry
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
			"DeleteSpawnEntry",
			"DELETE",
			"/spawn/{spawnID:[0-9]+}/entry/{entryID:[0-9]+}",
			deleteSpawnEntry,
		},
	}
	return
}

func getSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnEntryRequest{
		SpawnID: getIntVar(r, "spawnID"),
		EntryID: getIntVar(r, "entryID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnEntry := &model.SpawnEntry{
		SpawnID: request.SpawnID,
		ID:      request.EntryID,
	}

	err = cases.GetSpawnEntry(spawn, spawnEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpawnEntryResponse{
		Spawn:      spawn,
		SpawnEntry: spawnEntry,
	}
	content = response
	return
}

func createSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnEntryCreateRequest{
		SpawnID: getIntVar(r, "spawnID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		return
	}

	err = cases.CreateSpawnEntry(spawn, spawnEntry, user)
	if err != nil {
		return
	}
	response := &SpawnEntryResponse{
		Spawn:      spawn,
		SpawnEntry: spawnEntry,
	}
	content = response
	return
}

func deleteSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnEntryRequest{
		SpawnID: getIntVar(r, "spawnID"),
		EntryID: getIntVar(r, "entryID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnEntry := &model.SpawnEntry{
		SpawnID: request.SpawnID,
		ID:      request.EntryID,
	}

	err = cases.DeleteSpawnEntry(spawnEntry, spawn, user)
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

func editSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnEntryEditRequest{
		SpawnID: getIntVar(r, "spawnID"),
		EntryID: getIntVar(r, "entryID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnEntry := &model.SpawnEntry{}

	err = decodeBody(r, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	spawnEntry.ID = request.EntryID
	spawnEntry.SpawnID = request.SpawnID

	err = cases.EditSpawnEntry(spawn, spawnEntry, user)
	if err != nil {
		return
	}
	response := &SpawnEntryResponse{
		Spawn:      spawn,
		SpawnEntry: spawnEntry,
	}
	content = response
	return
}

func listSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnEntrysRequest{
		SpawnID: getIntVar(r, "spawnID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	spawnEntrys, err := cases.ListSpawnEntry(page, spawn, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpawnEntrysResponse{
		Page:        page,
		Spawn:       spawn,
		SpawnEntrys: spawnEntrys,
	}
	content = response
	return
}

func listSpawnEntryBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnEntrysBySearchRequest{
		SpawnID: getIntVar(r, "spawnID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	spawnEntry := &model.SpawnEntry{
		SpawnID: request.SpawnID,
	}
	spawnEntry.ID = getIntQuery(r, "entryID")

	spawnEntrys, err := cases.ListSpawnEntryBySearch(page, spawn, spawnEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpawnEntrysBySearchResponse{
		Page:        page,
		Spawn:       spawn,
		SpawnEntrys: spawnEntrys,
		Search:      spawnEntry,
	}
	content = response
	return
}
