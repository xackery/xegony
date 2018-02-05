package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpawnRequest is a list of parameters used for spawn
// swagger:parameters deleteSpawn editSpawn getSpawn
type SpawnRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
}

// SpawnResponse is what endpoints respond with
// swagger:response
type SpawnResponse struct {
	Spawn *model.Spawn `json:"spawn,omitempty"`
}

// SpawnCreateRequest is the body parameters for creating an spawn
// swagger:parameters createSpawn
type SpawnCreateRequest struct {
	// Spawn details to create
	// in: body
	Spawn *model.Spawn `json:"spawn"`
}

// SpawnEditRequest is the body parameters for creating an spawn
// swagger:parameters editSpawn
type SpawnEditRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
	// Spawn details to edit
	// in: body
	Spawn *model.Spawn `json:"spawn"`
}

// SpawnsRequest is a list of parameters used for spawn
// swagger:parameters listSpawn
type SpawnsRequest struct {
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SpawnsResponse is a general response to a request
// swagger:response
type SpawnsResponse struct {
	Page   *model.Page  `json:"page,omitempty"`
	Spawns model.Spawns `json:"spawns,omitempty"`
}

// SpawnsBySearchRequest is a list of parameters used for spawn
// swagger:parameters listSpawnBySearch
type SpawnsBySearchRequest struct {
	// Name is which spawn to get information about
	// example: heal
	// in: query
	Name string `json:"name"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SpawnsBySearchResponse is a general response to a request
// swagger:response
type SpawnsBySearchResponse struct {
	Search *model.Spawn `json:"search,omitempty"`
	Page   *model.Page  `json:"page,omitempty"`
	Spawns model.Spawns `json:"spawns,omitempty"`
}

func spawnRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spawn spawn listSpawn
		//
		// Lists spawns
		//
		// This will show all available spawns by default.
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
		//       200: SpawnsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawn",
			"GET",
			"/spawn",
			listSpawn,
		},
		// swagger:route GET /spawn/search spawn listSpawnBySearch
		//
		// Search spawns by name
		//
		// This will show all available spawns by default.
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
		//       200: SpawnsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawnBySearch",
			"GET",
			"/spawn/search",
			listSpawnBySearch,
		},
		// swagger:route POST /spawn spawn createSpawn
		//
		// Create an spawn
		//
		// This will create an spawn
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpawnResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpawn",
			"POST",
			"/spawn",
			createSpawn,
		},
		// swagger:route GET /spawn/{ID} spawn getSpawn
		//
		// Get an spawn
		//
		// This will get an individual spawn available spawns by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpawnResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpawn",
			"GET",
			"/spawn/{ID:[0-9]+}",
			getSpawn,
		},
		// swagger:route PUT /spawn/{ID} spawn editSpawn
		//
		// Edit an spawn
		//
		// This will edit an spawn
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpawnResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpawn",
			"PUT",
			"/spawn/{ID:[0-9]+}",
			editSpawn,
		},
		// swagger:route DELETE /spawn/{ID} spawn deleteSpawn
		//
		// Delete an spawn
		//
		// This will delete an spawn
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
			"DeleteSpawn",
			"DELETE",
			"/spawn/{ID:[0-9]+}",
			deleteSpawn,
		},
	}
	return
}

func getSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnRequest{
		ID: getIntVar(r, "ID"),
	}

	spawn := &model.Spawn{
		ID: request.ID,
	}

	err = cases.GetSpawn(spawn, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpawnResponse{
		Spawn: spawn,
	}
	content = response
	return
}

func createSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spawn := &model.Spawn{}
	err = decodeBody(r, spawn)
	if err != nil {
		return
	}
	err = cases.CreateSpawn(spawn, user)
	if err != nil {
		return
	}
	response := &SpawnResponse{
		Spawn: spawn,
	}
	content = response
	return
}

func deleteSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnRequest{
		ID: getIntVar(r, "ID"),
	}

	spawn := &model.Spawn{
		ID: request.ID,
	}

	err = cases.DeleteSpawn(spawn, user)
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

func editSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpawnEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spawn := &model.Spawn{
		ID: request.ID,
	}

	err = decodeBody(r, spawn)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpawn(spawn, user)
	if err != nil {
		return
	}
	response := &SpawnResponse{
		Spawn: spawn,
	}
	content = response
	return
}

func listSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spawns, err := cases.ListSpawn(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpawnsResponse{
		Page:   page,
		Spawns: spawns,
	}
	content = response
	return
}

func listSpawnBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spawn := &model.Spawn{}
	spawn.Name = getQuery(r, "name")
	spawns, err := cases.ListSpawnBySearch(page, spawn, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpawnsBySearchResponse{
		Page:   page,
		Spawns: spawns,
		Search: spawn,
	}
	content = response
	return
}
