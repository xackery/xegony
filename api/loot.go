package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// LootRequest is a list of parameters used for loot
// swagger:parameters deleteLoot editLoot getLoot
type LootRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
}

// LootResponse is what endpoints respond with
// swagger:response
type LootResponse struct {
	Loot *model.Loot `json:"loot,omitempty"`
}

// LootCreateRequest is the body parameters for creating an loot
// swagger:parameters createLoot
type LootCreateRequest struct {
	// Loot details to create
	// in: body
	Loot *model.Loot `json:"loot"`
}

// LootEditRequest is the body parameters for creating an loot
// swagger:parameters editLoot
type LootEditRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
	// Loot details to edit
	// in: body
	Loot *model.Loot `json:"loot"`
}

// LootsRequest is a list of parameters used for loot
// swagger:parameters listLoot
type LootsRequest struct {
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

// LootsResponse is a general response to a request
// swagger:response
type LootsResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Loots model.Loots `json:"loots,omitempty"`
}

// LootsBySearchRequest is a list of parameters used for loot
// swagger:parameters listLootBySearch
type LootsBySearchRequest struct {
	// Name is which loot to get information about
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

// LootsBySearchResponse is a general response to a request
// swagger:response
type LootsBySearchResponse struct {
	Search *model.Loot `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Loots  model.Loots `json:"loots,omitempty"`
}

func lootRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /loot loot listLoot
		//
		// Lists loots
		//
		// This will show all available loots by default.
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
		//       200: LootsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLoot",
			"GET",
			"/loot",
			listLoot,
		},
		// swagger:route GET /loot/search loot listLootBySearch
		//
		// Search loots by name
		//
		// This will show all available loots by default.
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
		//       200: LootsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListLootBySearch",
			"GET",
			"/loot/search",
			listLootBySearch,
		},
		// swagger:route POST /loot loot createLoot
		//
		// Create an loot
		//
		// This will create an loot
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateLoot",
			"POST",
			"/loot",
			createLoot,
		},
		// swagger:route GET /loot/{ID} loot getLoot
		//
		// Get an loot
		//
		// This will get an individual loot available loots by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: LootResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetLoot",
			"GET",
			"/loot/{ID:[0-9]+}",
			getLoot,
		},
		// swagger:route PUT /loot/{ID} loot editLoot
		//
		// Edit an loot
		//
		// This will edit an loot
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: LootResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditLoot",
			"PUT",
			"/loot/{ID:[0-9]+}",
			editLoot,
		},
		// swagger:route DELETE /loot/{ID} loot deleteLoot
		//
		// Delete an loot
		//
		// This will delete an loot
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
			"DeleteLoot",
			"DELETE",
			"/loot/{ID:[0-9]+}",
			deleteLoot,
		},
	}
	return
}

func getLoot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootRequest{
		ID: getIntVar(r, "ID"),
	}

	loot := &model.Loot{
		ID: request.ID,
	}

	err = cases.GetLoot(loot, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &LootResponse{
		Loot: loot,
	}
	content = response
	return
}

func createLoot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	loot := &model.Loot{}
	err = decodeBody(r, loot)
	if err != nil {
		return
	}
	err = cases.CreateLoot(loot, user)
	if err != nil {
		return
	}
	response := &LootResponse{
		Loot: loot,
	}
	content = response
	return
}

func deleteLoot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootRequest{
		ID: getIntVar(r, "ID"),
	}

	loot := &model.Loot{
		ID: request.ID,
	}

	err = cases.DeleteLoot(loot, user)
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

func editLoot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &LootEditRequest{
		ID: getIntVar(r, "ID"),
	}

	loot := &model.Loot{
		ID: request.ID,
	}

	err = decodeBody(r, loot)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditLoot(loot, user)
	if err != nil {
		return
	}
	response := &LootResponse{
		Loot: loot,
	}
	content = response
	return
}

func listLoot(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	loots, err := cases.ListLoot(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &LootsResponse{
		Page:  page,
		Loots: loots,
	}
	content = response
	return
}

func listLootBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	loot := &model.Loot{}
	loot.Name = getQuery(r, "name")
	loots, err := cases.ListLootBySearch(page, loot, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &LootsBySearchResponse{
		Page:   page,
		Loots:  loots,
		Search: loot,
	}
	content = response
	return
}
