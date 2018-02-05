package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// NpcRequest is a list of parameters used for npc
// swagger:parameters deleteNpc editNpc getNpc
type NpcRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
}

// NpcResponse is what endpoints respond with
// swagger:response
type NpcResponse struct {
	Npc *model.Npc `json:"npc,omitempty"`
}

// NpcCreateRequest is the body parameters for creating an npc
// swagger:parameters createNpc
type NpcCreateRequest struct {
	// Npc details to create
	// in: body
	Npc *model.Npc `json:"npc"`
}

// NpcEditRequest is the body parameters for creating an npc
// swagger:parameters editNpc
type NpcEditRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
	// Npc details to edit
	// in: body
	Npc *model.Npc `json:"npc"`
}

// NpcsRequest is a list of parameters used for npc
// swagger:parameters listNpc
type NpcsRequest struct {
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

// NpcsResponse is a general response to a request
// swagger:response
type NpcsResponse struct {
	Page *model.Page `json:"page,omitempty"`
	Npcs model.Npcs  `json:"npcs,omitempty"`
}

// NpcsBySearchRequest is a list of parameters used for npc
// swagger:parameters listNpcBySearch
type NpcsBySearchRequest struct {
	// Name is which npc to get information about
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

// NpcsBySearchResponse is a general response to a request
// swagger:response
type NpcsBySearchResponse struct {
	Search *model.Npc  `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Npcs   model.Npcs  `json:"npcs,omitempty"`
}

func npcRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /npc npc listNpc
		//
		// Lists npcs
		//
		// This will show all available npcs by default.
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
		//       200: NpcsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListNpc",
			"GET",
			"/npc",
			listNpc,
		},
		// swagger:route GET /npc/search npc listNpcBySearch
		//
		// Search npcs by name
		//
		// This will show all available npcs by default.
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
		//       200: NpcsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListNpcBySearch",
			"GET",
			"/npc/search",
			listNpcBySearch,
		},
		// swagger:route POST /npc npc createNpc
		//
		// Create an npc
		//
		// This will create an npc
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: NpcResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateNpc",
			"POST",
			"/npc",
			createNpc,
		},
		// swagger:route GET /npc/{ID} npc getNpc
		//
		// Get an npc
		//
		// This will get an individual npc available npcs by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: NpcResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetNpc",
			"GET",
			"/npc/{ID:[0-9]+}",
			getNpc,
		},
		// swagger:route PUT /npc/{ID} npc editNpc
		//
		// Edit an npc
		//
		// This will edit an npc
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: NpcResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditNpc",
			"PUT",
			"/npc/{ID:[0-9]+}",
			editNpc,
		},
		// swagger:route DELETE /npc/{ID} npc deleteNpc
		//
		// Delete an npc
		//
		// This will delete an npc
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
			"DeleteNpc",
			"DELETE",
			"/npc/{ID:[0-9]+}",
			deleteNpc,
		},
	}
	return
}

func getNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &NpcRequest{
		ID: getIntVar(r, "ID"),
	}

	npc := &model.Npc{
		ID: request.ID,
	}

	err = cases.GetNpc(npc, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &NpcResponse{
		Npc: npc,
	}
	content = response
	return
}

func createNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	npc := &model.Npc{}
	err = decodeBody(r, npc)
	if err != nil {
		return
	}
	err = cases.CreateNpc(npc, user)
	if err != nil {
		return
	}
	response := &NpcResponse{
		Npc: npc,
	}
	content = response
	return
}

func deleteNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &NpcRequest{
		ID: getIntVar(r, "ID"),
	}

	npc := &model.Npc{
		ID: request.ID,
	}

	err = cases.DeleteNpc(npc, user)
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

func editNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &NpcEditRequest{
		ID: getIntVar(r, "ID"),
	}

	npc := &model.Npc{
		ID: request.ID,
	}

	err = decodeBody(r, npc)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditNpc(npc, user)
	if err != nil {
		return
	}
	response := &NpcResponse{
		Npc: npc,
	}
	content = response
	return
}

func listNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	npcs, err := cases.ListNpc(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &NpcsResponse{
		Page: page,
		Npcs: npcs,
	}
	content = response
	return
}

func listNpcBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	npc := &model.Npc{}
	npc.Name = getQuery(r, "name")
	npcs, err := cases.ListNpcBySearch(page, npc, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &NpcsBySearchResponse{
		Page:   page,
		Npcs:   npcs,
		Search: npc,
	}
	content = response
	return
}
