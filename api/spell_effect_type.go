package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellEffectTypeRequest is a list of parameters used for spellEffectType
// swagger:parameters deleteSpellEffectType editSpellEffectType getSpellEffectType
type SpellEffectTypeRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// SpellEffectTypeResponse is what endpoints respond with
// swagger:response
type SpellEffectTypeResponse struct {
	SpellEffectType *model.SpellEffectType `json:"spellEffectType,omitempty"`
}

// SpellEffectTypeCreateRequest is the body parameters for creating an spellEffectType
// swagger:parameters createSpellEffectType
type SpellEffectTypeCreateRequest struct {
	// SpellEffectType details to create
	// in: body
	SpellEffectType *model.SpellEffectType `json:"spellEffectType"`
}

// SpellEffectTypeEditRequest is the body parameters for creating an spellEffectType
// swagger:parameters editSpellEffectType
type SpellEffectTypeEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// SpellEffectType details to edit
	// in: body
	SpellEffectType *model.SpellEffectType `json:"spellEffectType"`
}

// SpellEffectTypesRequest is a list of parameters used for spellEffectType
// swagger:parameters listSpellEffectType
type SpellEffectTypesRequest struct {
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

// SpellEffectTypesResponse is a general response to a request
// swagger:response
type SpellEffectTypesResponse struct {
	Page             *model.Page            `json:"page,omitempty"`
	SpellEffectTypes model.SpellEffectTypes `json:"spellEffectTypes,omitempty"`
}

// SpellEffectTypesBySearchRequest is a list of parameters used for spellEffectType
// swagger:parameters listSpellEffectTypeBySearch
type SpellEffectTypesBySearchRequest struct {
	// Name is which spellEffectType to get information about
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

// SpellEffectTypesBySearchResponse is a general response to a request
// swagger:response
type SpellEffectTypesBySearchResponse struct {
	Search           *model.SpellEffectType `json:"search,omitempty"`
	Page             *model.Page            `json:"page,omitempty"`
	SpellEffectTypes model.SpellEffectTypes `json:"spellEffectTypes,omitempty"`
}

func spellEffectTypeRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell/effect/type spellEffectType listSpellEffectType
		//
		// Lists spellEffectTypes
		//
		// This will show all available spellEffectTypes by default.
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
		//       200: SpellEffectTypesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellEffectType",
			"GET",
			"/spell/effect/type",
			listSpellEffectType,
		},
		// swagger:route GET /spell/effect/type/search spellEffectType listSpellEffectTypeBySearch
		//
		// Search spellEffectTypes by name
		//
		// This will show all available spellEffectTypes by default.
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
		//       200: SpellEffectTypesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellEffectTypeBySearch",
			"GET",
			"/spell/effect/type/search",
			listSpellEffectTypeBySearch,
		},
		// swagger:route POST /spell/effect/type spellEffectType createSpellEffectType
		//
		// Create an spellEffectType
		//
		// This will create an spellEffectType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellEffectTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpellEffectType",
			"POST",
			"/spell/effect/type",
			createSpellEffectType,
		},
		// swagger:route GET /spell/effect/type/{ID} spellEffectType getSpellEffectType
		//
		// Get an spellEffectType
		//
		// This will get an individual spellEffectType available spellEffectTypes by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellEffectTypeResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpellEffectType",
			"GET",
			"/spell/effect/type/{ID:[0-9]+}",
			getSpellEffectType,
		},
		// swagger:route PUT /spell/effect/type/{ID} spellEffectType editSpellEffectType
		//
		// Edit an spellEffectType
		//
		// This will edit an spellEffectType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellEffectTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpellEffectType",
			"PUT",
			"/spell/effect/type/{ID:[0-9]+}",
			editSpellEffectType,
		},
		// swagger:route DELETE /spell/effect/type/{ID} spellEffectType deleteSpellEffectType
		//
		// Delete an spellEffectType
		//
		// This will delete an spellEffectType
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
			"DeleteSpellEffectType",
			"DELETE",
			"/spell/effect/type/{ID:[0-9]+}",
			deleteSpellEffectType,
		},
	}
	return
}

func getSpellEffectType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellEffectTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellEffectType := &model.SpellEffectType{
		ID: request.ID,
	}

	err = cases.GetSpellEffectType(spellEffectType, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellEffectTypeResponse{
		SpellEffectType: spellEffectType,
	}
	content = response
	return
}

func createSpellEffectType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spellEffectType := &model.SpellEffectType{}
	err = decodeBody(r, spellEffectType)
	if err != nil {
		return
	}
	err = cases.CreateSpellEffectType(spellEffectType, user)
	if err != nil {
		return
	}
	response := &SpellEffectTypeResponse{
		SpellEffectType: spellEffectType,
	}
	content = response
	return
}

func deleteSpellEffectType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellEffectTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellEffectType := &model.SpellEffectType{
		ID: request.ID,
	}

	err = cases.DeleteSpellEffectType(spellEffectType, user)
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

func editSpellEffectType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellEffectTypeEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spellEffectType := &model.SpellEffectType{
		ID: request.ID,
	}

	err = decodeBody(r, spellEffectType)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpellEffectType(spellEffectType, user)
	if err != nil {
		return
	}
	response := &SpellEffectTypeResponse{
		SpellEffectType: spellEffectType,
	}
	content = response
	return
}

func listSpellEffectType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellEffectTypes, err := cases.ListSpellEffectType(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellEffectTypesResponse{
		Page:             page,
		SpellEffectTypes: spellEffectTypes,
	}
	content = response
	return
}

func listSpellEffectTypeBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellEffectType := &model.SpellEffectType{
		Name: getQuery(r, "name"),
	}
	spellEffectTypes, err := cases.ListSpellEffectTypeBySearch(page, spellEffectType, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellEffectTypesBySearchResponse{
		Page:             page,
		SpellEffectTypes: spellEffectTypes,
		Search:           spellEffectType,
	}
	content = response
	return
}
