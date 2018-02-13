package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// LootDropRequest is a list of parameters used for lootDrop
// swagger:parameters deleteLootDrop editLootDrop getLootDrop
type LootDropRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
}

// LootDropResponse is what endpoints respond with
// swagger:response
type LootDropResponse struct {
	LootDrop *model.LootDrop `json:"lootDrop,omitempty"`
}

// LootDropCreateRequest is the body parameters for creating an lootDrop
// swagger:parameters createLootDrop
type LootDropCreateRequest struct {
	// LootDrop details to create
	// in: body
	LootDrop *model.LootDrop `json:"lootDrop"`
}

// LootDropEditRequest is the body parameters for creating an lootDrop
// swagger:parameters editLootDrop
type LootDropEditRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
	// LootDrop details to edit
	// in: body
	LootDrop *model.LootDrop `json:"lootDrop"`
}

// LootDropsRequest is a list of parameters used for lootDrop
// swagger:parameters listLootDrop
type LootDropsRequest struct {
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

// LootDropsResponse is a general response to a request
// swagger:response
type LootDropsResponse struct {
	Page      *model.Page     `json:"page,omitempty"`
	LootDrops model.LootDrops `json:"lootDrops,omitempty"`
}

// LootDropsBySearchRequest is a list of parameters used for lootDrop
// swagger:parameters listLootDropBySearch
type LootDropsBySearchRequest struct {
	// Name is which lootDrop to get information about
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

// LootDropsBySearchResponse is a general response to a request
// swagger:response
type LootDropsBySearchResponse struct {
	Search    *model.LootDrop `json:"search,omitempty"`
	Page      *model.Page     `json:"page,omitempty"`
	LootDrops model.LootDrops `json:"lootDrops,omitempty"`
}

func lootDropRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /loot/drop lootDrop listLootDrop
		//
		// Lists lootDrops
		//
		// This will show all available lootDrops by default.
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
		//       200: LootDropsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLootDrop",
			"GET",
			"/loot/drop",
			listLootDrop,
		},
		// swagger:route GET /loot/drop/search lootDrop listLootDropBySearch
		//
		// Search lootDrops by name
		//
		// This will show all available lootDrops by default.
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
		//       200: LootDropsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLootDropBySearch",
			"GET",
			"/loot/drop/search",
			listLootDropBySearch,
		},
		// swagger:route POST /loot/drop lootDrop createLootDrop
		//
		// Create an lootDrop
		//
		// This will create an lootDrop
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootDropResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateLootDrop",
			"POST",
			"/loot/drop",
			createLootDrop,
		},
		// swagger:route GET /loot/drop/{ID} lootDrop getLootDrop
		//
		// Get an lootDrop
		//
		// This will get an individual lootDrop available lootDrops by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: LootDropResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetLootDrop",
			"GET",
			"/loot/drop/{ID:[0-9]+}",
			getLootDrop,
		},
		// swagger:route PUT /loot/drop/{ID} lootDrop editLootDrop
		//
		// Edit an lootDrop
		//
		// This will edit an lootDrop
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootDropResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditLootDrop",
			"PUT",
			"/loot/drop/{ID:[0-9]+}",
			editLootDrop,
		},
		// swagger:route DELETE /loot/drop/{ID} lootDrop deleteLootDrop
		//
		// Delete an lootDrop
		//
		// This will delete an lootDrop
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
			"DeleteLootDrop",
			"DELETE",
			"/loot/drop/{ID:[0-9]+}",
			deleteLootDrop,
		},
	}
	return
}

func getLootDrop(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropRequest{
		ID: getIntVar(r, "ID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.ID,
	}

	err = cases.GetLootDrop(lootDrop, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &LootDropResponse{
		LootDrop: lootDrop,
	}
	content = response
	return
}

func createLootDrop(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	lootDrop := &model.LootDrop{}
	err = decodeBody(r, lootDrop)
	if err != nil {
		return
	}
	err = cases.CreateLootDrop(lootDrop, user)
	if err != nil {
		return
	}
	response := &LootDropResponse{
		LootDrop: lootDrop,
	}
	content = response
	return
}

func deleteLootDrop(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropRequest{
		ID: getIntVar(r, "ID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.ID,
	}

	err = cases.DeleteLootDrop(lootDrop, user)
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

func editLootDrop(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootDropEditRequest{
		ID: getIntVar(r, "ID"),
	}

	lootDrop := &model.LootDrop{
		ID: request.ID,
	}

	err = decodeBody(r, lootDrop)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditLootDrop(lootDrop, user)
	if err != nil {
		return
	}
	response := &LootDropResponse{
		LootDrop: lootDrop,
	}
	content = response
	return
}

func listLootDrop(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	lootDrops, err := cases.ListLootDrop(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &LootDropsResponse{
		Page:      page,
		LootDrops: lootDrops,
	}
	content = response
	return
}

func listLootDropBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	lootDrop := &model.LootDrop{}
	lootDrop.Name = getQuery(r, "name")
	lootDrops, err := cases.ListLootDropBySearch(page, lootDrop, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &LootDropsBySearchResponse{
		Page:      page,
		LootDrops: lootDrops,
		Search:    lootDrop,
	}
	content = response
	return
}
