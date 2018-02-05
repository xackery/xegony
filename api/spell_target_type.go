package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellTargetTypeRequest is a list of parameters used for spellTargetType
// swagger:parameters deleteSpellTargetType editSpellTargetType getSpellTargetType
type SpellTargetTypeRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// SpellTargetTypeResponse is what endpoints respond with
// swagger:response
type SpellTargetTypeResponse struct {
	SpellTargetType *model.SpellTargetType `json:"spellTargetType,omitempty"`
}

// SpellTargetTypeCreateRequest is the body parameters for creating an spellTargetType
// swagger:parameters createSpellTargetType
type SpellTargetTypeCreateRequest struct {
	// SpellTargetType details to create
	// in: body
	SpellTargetType *model.SpellTargetType `json:"spellTargetType"`
}

// SpellTargetTypeEditRequest is the body parameters for creating an spellTargetType
// swagger:parameters editSpellTargetType
type SpellTargetTypeEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// SpellTargetType details to edit
	// in: body
	SpellTargetType *model.SpellTargetType `json:"spellTargetType"`
}

// SpellTargetTypesRequest is a list of parameters used for spellTargetType
// swagger:parameters listSpellTargetType
type SpellTargetTypesRequest struct {
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

// SpellTargetTypesResponse is a general response to a request
// swagger:response
type SpellTargetTypesResponse struct {
	Page             *model.Page            `json:"page,omitempty"`
	SpellTargetTypes model.SpellTargetTypes `json:"spellTargetTypes,omitempty"`
}

// SpellTargetTypesBySearchRequest is a list of parameters used for spellTargetType
// swagger:parameters listSpellTargetTypeBySearch
type SpellTargetTypesBySearchRequest struct {
	// Name is which spellTargetType to get information about
	// example: xackery
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
	// example: id
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// SpellTargetTypesBySearchResponse is a general response to a request
// swagger:response
type SpellTargetTypesBySearchResponse struct {
	Search           *model.SpellTargetType `json:"search,omitempty"`
	Page             *model.Page            `json:"page,omitempty"`
	SpellTargetTypes model.SpellTargetTypes `json:"spellTargetTypes,omitempty"`
}

func spellTargetTypeRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell/target/type spellTargetType listSpellTargetType
		//
		// Lists spellTargetTypes
		//
		// This will show all available spellTargetTypes by default.
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
		//       200: SpellTargetTypesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellTargetType",
			"GET",
			"/spell/target/type",
			listSpellTargetType,
		},
		// swagger:route GET /spell/target/type/search spellTargetType listSpellTargetTypeBySearch
		//
		// Search spellTargetTypes by name
		//
		// This will show all available spellTargetTypes by default.
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
		//       200: SpellTargetTypesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellTargetTypeBySearch",
			"GET",
			"/spell/target/type/search",
			listSpellTargetTypeBySearch,
		},
		// swagger:route POST /spell/target/type spellTargetType createSpellTargetType
		//
		// Create an spellTargetType
		//
		// This will create an spellTargetType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellTargetTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpellTargetType",
			"POST",
			"/spell/target/type",
			createSpellTargetType,
		},
		// swagger:route GET /spell/target/type/{ID} spellTargetType getSpellTargetType
		//
		// Get an spellTargetType
		//
		// This will get an individual spellTargetType available spellTargetTypes by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellTargetTypeResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpellTargetType",
			"GET",
			"/spell/target/type/{ID:[0-9]+}",
			getSpellTargetType,
		},
		// swagger:route PUT /spell/target/type/{ID} spellTargetType editSpellTargetType
		//
		// Edit an spellTargetType
		//
		// This will edit an spellTargetType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellTargetTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpellTargetType",
			"PUT",
			"/spell/target/type/{ID:[0-9]+}",
			editSpellTargetType,
		},
		// swagger:route DELETE /spell/target/type/{ID} spellTargetType deleteSpellTargetType
		//
		// Delete an spellTargetType
		//
		// This will delete an spellTargetType
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
			"DeleteSpellTargetType",
			"DELETE",
			"/spell/target/type/{ID:[0-9]+}",
			deleteSpellTargetType,
		},
	}
	return
}

func getSpellTargetType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellTargetTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellTargetType := &model.SpellTargetType{
		ID: request.ID,
	}

	err = cases.GetSpellTargetType(spellTargetType, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellTargetTypeResponse{
		SpellTargetType: spellTargetType,
	}
	content = response
	return
}

func createSpellTargetType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spellTargetType := &model.SpellTargetType{}
	err = decodeBody(r, spellTargetType)
	if err != nil {
		return
	}
	err = cases.CreateSpellTargetType(spellTargetType, user)
	if err != nil {
		return
	}
	response := &SpellTargetTypeResponse{
		SpellTargetType: spellTargetType,
	}
	content = response
	return
}

func deleteSpellTargetType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellTargetTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellTargetType := &model.SpellTargetType{
		ID: request.ID,
	}

	err = cases.DeleteSpellTargetType(spellTargetType, user)
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

func editSpellTargetType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellTargetTypeEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spellTargetType := &model.SpellTargetType{
		ID: request.ID,
	}

	err = decodeBody(r, spellTargetType)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpellTargetType(spellTargetType, user)
	if err != nil {
		return
	}
	response := &SpellTargetTypeResponse{
		SpellTargetType: spellTargetType,
	}
	content = response
	return
}

func listSpellTargetType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellTargetTypes, err := cases.ListSpellTargetType(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellTargetTypesResponse{
		Page:             page,
		SpellTargetTypes: spellTargetTypes,
	}
	content = response
	return
}

func listSpellTargetTypeBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellTargetType := &model.SpellTargetType{
		Name: getQuery(r, "name"),
	}
	spellTargetTypes, err := cases.ListSpellTargetTypeBySearch(page, spellTargetType, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellTargetTypesBySearchResponse{
		Page:             page,
		SpellTargetTypes: spellTargetTypes,
		Search:           spellTargetType,
	}
	content = response
	return
}
