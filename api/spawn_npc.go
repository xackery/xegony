package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpawnNpcRequest is a list of parameters used for spawnNpc
// swagger:parameters deleteSpawnNpc editSpawnNpc getSpawnNpc
type SpawnNpcRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"spawnID"`
	// NpcID to get information about
	// in: path
	// example: 55091
	NpcID int64 `json:"npcID"`
}

// SpawnNpcResponse is what endpoints respond with
// swagger:response
type SpawnNpcResponse struct {
	Spawn    *model.Spawn    `json:"spawn,omitempty"`
	SpawnNpc *model.SpawnNpc `json:"spawnNpc,omitempty"`
}

// SpawnNpcCreateRequest is the body parameters for creating an spawnNpc
// swagger:parameters createSpawnNpc
type SpawnNpcCreateRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"spawnID"`
	// SpawnNpc details to create
	// in: body
	SpawnNpc *model.SpawnNpc `json:"spawnNpc"`
}

// SpawnNpcEditRequest is the body parameters for creating an spawnNpc
// swagger:parameters editSpawnNpc
type SpawnNpcEditRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"SpawnID"`
	// NpcID to get information about
	// in: path
	// example: 55091
	NpcID int64 `json:"npcID"`
	// SpawnNpc details to edit
	// in: body
	SpawnNpc *model.SpawnNpc `json:"spawnNpc"`
}

// SpawnNpcsRequest is a list of parameters used for spawnNpc
// swagger:parameters listSpawnNpc
type SpawnNpcsRequest struct {
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
	// example: npcid
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SpawnNpcsResponse is a general response to a request
// swagger:response
type SpawnNpcsResponse struct {
	Page      *model.Page     `json:"page,omitempty"`
	Spawn     *model.Spawn    `json:"spawn"`
	SpawnNpcs model.SpawnNpcs `json:"spawnNpcs,omitempty"`
}

// SpawnNpcsBySearchRequest is a list of parameters used for spawnNpc
// swagger:parameters listSpawnNpcBySearch
type SpawnNpcsBySearchRequest struct {
	// SpawnID to get information about
	// in: path
	// example: 1
	SpawnID int64 `json:"spawnID"`
	// NpcID is which spawnNpc to get information about
	// example: 55091
	// in: query
	NpcID int64 `json:"npcID"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: npcid
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SpawnNpcsBySearchResponse is a general response to a request
// swagger:response
type SpawnNpcsBySearchResponse struct {
	Search    *model.SpawnNpc `json:"search,omitempty"`
	Page      *model.Page     `json:"page,omitempty"`
	Spawn     *model.Spawn    `json:"spawn,omitempty"`
	SpawnNpcs model.SpawnNpcs `json:"spawnNpcs,omitempty"`
}

func spawnNpcRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spawn/{spawnID}/npc spawn listSpawnNpc
		//
		// Lists spawnNpcs
		//
		// This will show all available spawnNpcs by default.
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
		//       200: SpawnNpcsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawnNpc",
			"GET",
			"/spawn/{spawnID:[0-9]+}/npc",
			listSpawnNpc,
		},
		// swagger:route GET /spawn/{spawnID}/npc/search spawn listSpawnNpcBySearch
		//
		// Search spawnNpcs by npcid
		//
		// This will show all available spawnNpcs by default.
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
		//       200: SpawnNpcsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawnNpcBySearch",
			"GET",
			"/spawn/npc/search",
			listSpawnNpcBySearch,
		},
		// swagger:route POST /spawn/{spawnID}/npc/{npcID} spawn createSpawnNpc
		//
		// Create an spawnNpc
		//
		// This will create an spawnNpc
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpawnNpcResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpawnNpc",
			"POST",
			"/spawn/{spawnID:[0-9]+}/npc/{npcID:[0-9]+}",
			createSpawnNpc,
		},
		// swagger:route GET /spawn/{spawnID}/npc/{npcID} spawn getSpawnNpc
		//
		// Get an spawnNpc
		//
		// This will get an individual spawnNpc available spawnNpcs by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpawnNpcResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpawnNpc",
			"GET",
			"/spawn/{spawnID:[0-9]+}/npc/{npcID:[0-9]+}",
			getSpawnNpc,
		},
		// swagger:route PUT /spawn/{spawnID}/npc/{npcID} spawn editSpawnNpc
		//
		// Edit an spawnNpc
		//
		// This will edit an spawnNpc
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpawnNpcResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpawnNpc",
			"PUT",
			"/spawn/{spawnID:[0-9]+}/npc/{npcID:[0-9]+}",
			editSpawnNpc,
		},
		// swagger:route DELETE /spawn/{spawnID}/npc/{npcID} spawn deleteSpawnNpc
		//
		// Delete an spawnNpc
		//
		// This will delete an spawnNpc
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
			"DeleteSpawnNpc",
			"DELETE",
			"/spawn/{spawnID:[0-9]+}/npc/{npcID:[0-9]+}",
			deleteSpawnNpc,
		},
	}
	return
}

func getSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnNpcRequest{
		SpawnID: getIntVar(r, "spawnID"),
		NpcID:   getIntVar(r, "npcID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnNpc := &model.SpawnNpc{
		SpawnID: request.SpawnID,
		NpcID:   request.NpcID,
	}

	err = cases.GetSpawnNpc(spawn, spawnNpc, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpawnNpcResponse{
		Spawn:    spawn,
		SpawnNpc: spawnNpc,
	}
	content = response
	return
}

func createSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnNpcCreateRequest{
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

	spawnNpc := &model.SpawnNpc{}
	err = decodeBody(r, spawnNpc)
	if err != nil {
		return
	}

	err = cases.CreateSpawnNpc(spawn, spawnNpc, user)
	if err != nil {
		return
	}
	response := &SpawnNpcResponse{
		Spawn:    spawn,
		SpawnNpc: spawnNpc,
	}
	content = response
	return
}

func deleteSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnNpcRequest{
		SpawnID: getIntVar(r, "spawnID"),
		NpcID:   getIntVar(r, "npcID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnNpc := &model.SpawnNpc{
		SpawnID: request.SpawnID,
		NpcID:   request.NpcID,
	}

	err = cases.DeleteSpawnNpc(spawnNpc, spawn, user)
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

func editSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnNpcEditRequest{
		SpawnID: getIntVar(r, "spawnID"),
		NpcID:   getIntVar(r, "npcID"),
	}

	spawn := &model.Spawn{
		ID: request.SpawnID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get spawn")
		return
	}

	spawnNpc := &model.SpawnNpc{}

	err = decodeBody(r, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	spawnNpc.NpcID = request.NpcID
	spawnNpc.SpawnID = request.SpawnID

	err = cases.EditSpawnNpc(spawn, spawnNpc, user)
	if err != nil {
		return
	}
	response := &SpawnNpcResponse{
		Spawn:    spawn,
		SpawnNpc: spawnNpc,
	}
	content = response
	return
}

func listSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnNpcsRequest{
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

	spawnNpcs, err := cases.ListSpawnNpc(page, spawn, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpawnNpcsResponse{
		Page:      page,
		Spawn:     spawn,
		SpawnNpcs: spawnNpcs,
	}
	content = response
	return
}

func listSpawnNpcBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnNpcsBySearchRequest{
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

	spawnNpc := &model.SpawnNpc{
		SpawnID: request.SpawnID,
	}
	spawnNpc.NpcID = getIntQuery(r, "npcID")

	spawnNpcs, err := cases.ListSpawnNpcBySearch(page, spawn, spawnNpc, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpawnNpcsBySearchResponse{
		Page:      page,
		Spawn:     spawn,
		SpawnNpcs: spawnNpcs,
		Search:    spawnNpc,
	}
	content = response
	return
}
