package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellTravelTypeRequest is a list of parameters used for spellTravelType
// swagger:parameters deleteSpellTravelType editSpellTravelType getSpellTravelType
type SpellTravelTypeRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// SpellTravelTypeResponse is what endpoints respond with
// swagger:response
type SpellTravelTypeResponse struct {
	SpellTravelType *model.SpellTravelType `json:"spellTravelType,omitempty"`
}

// SpellTravelTypeCreateRequest is the body parameters for creating an spellTravelType
// swagger:parameters createSpellTravelType
type SpellTravelTypeCreateRequest struct {
	// SpellTravelType details to create
	// in: body
	SpellTravelType *model.SpellTravelType `json:"spellTravelType"`
}

// SpellTravelTypeEditRequest is the body parameters for creating an spellTravelType
// swagger:parameters editSpellTravelType
type SpellTravelTypeEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// SpellTravelType details to edit
	// in: body
	SpellTravelType *model.SpellTravelType `json:"spellTravelType"`
}

// SpellTravelTypesRequest is a list of parameters used for spellTravelType
// swagger:parameters listSpellTravelType
type SpellTravelTypesRequest struct {
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

// SpellTravelTypesResponse is a general response to a request
// swagger:response
type SpellTravelTypesResponse struct {
	Page             *model.Page            `json:"page,omitempty"`
	SpellTravelTypes model.SpellTravelTypes `json:"spellTravelTypes,omitempty"`
}

// SpellTravelTypesBySearchRequest is a list of parameters used for spellTravelType
// swagger:parameters listSpellTravelTypeBySearch
type SpellTravelTypesBySearchRequest struct {
	// Name is which spellTravelType to get information about
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

// SpellTravelTypesBySearchResponse is a general response to a request
// swagger:response
type SpellTravelTypesBySearchResponse struct {
	Search           *model.SpellTravelType `json:"search,omitempty"`
	Page             *model.Page            `json:"page,omitempty"`
	SpellTravelTypes model.SpellTravelTypes `json:"spellTravelTypes,omitempty"`
}

func spellTravelTypeRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell/travel/type spellTravelType listSpellTravelType
		//
		// Lists spellTravelTypes
		//
		// This will show all available spellTravelTypes by default.
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
		//       200: SpellTravelTypesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellTravelType",
			"GET",
			"/spell/travel/type",
			listSpellTravelType,
		},
		// swagger:route GET /spell/travel/type/search spellTravelType listSpellTravelTypeBySearch
		//
		// Search spellTravelTypes by name
		//
		// This will show all available spellTravelTypes by default.
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
		//       200: SpellTravelTypesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellTravelTypeBySearch",
			"GET",
			"/spell/travel/type/search",
			listSpellTravelTypeBySearch,
		},
		// swagger:route POST /spell/travel/type spellTravelType createSpellTravelType
		//
		// Create an spellTravelType
		//
		// This will create an spellTravelType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellTravelTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpellTravelType",
			"POST",
			"/spell/travel/type",
			createSpellTravelType,
		},
		// swagger:route GET /spell/travel/type/{ID} spellTravelType getSpellTravelType
		//
		// Get an spellTravelType
		//
		// This will get an individual spellTravelType available spellTravelTypes by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellTravelTypeResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpellTravelType",
			"GET",
			"/spell/travel/type/{ID:[0-9]+}",
			getSpellTravelType,
		},
		// swagger:route PUT /spell/travel/type/{ID} spellTravelType editSpellTravelType
		//
		// Edit an spellTravelType
		//
		// This will edit an spellTravelType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellTravelTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpellTravelType",
			"PUT",
			"/spell/travel/type/{ID:[0-9]+}",
			editSpellTravelType,
		},
		// swagger:route DELETE /spell/travel/type/{ID} spellTravelType deleteSpellTravelType
		//
		// Delete an spellTravelType
		//
		// This will delete an spellTravelType
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
			"DeleteSpellTravelType",
			"DELETE",
			"/spell/travel/type/{ID:[0-9]+}",
			deleteSpellTravelType,
		},
	}
	return
}

func getSpellTravelType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellTravelTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellTravelType := &model.SpellTravelType{
		ID: request.ID,
	}

	err = cases.GetSpellTravelType(spellTravelType, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellTravelTypeResponse{
		SpellTravelType: spellTravelType,
	}
	content = response
	return
}

func createSpellTravelType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spellTravelType := &model.SpellTravelType{}
	err = decodeBody(r, spellTravelType)
	if err != nil {
		return
	}
	err = cases.CreateSpellTravelType(spellTravelType, user)
	if err != nil {
		return
	}
	response := &SpellTravelTypeResponse{
		SpellTravelType: spellTravelType,
	}
	content = response
	return
}

func deleteSpellTravelType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellTravelTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellTravelType := &model.SpellTravelType{
		ID: request.ID,
	}

	err = cases.DeleteSpellTravelType(spellTravelType, user)
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

func editSpellTravelType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellTravelTypeEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spellTravelType := &model.SpellTravelType{
		ID: request.ID,
	}

	err = decodeBody(r, spellTravelType)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpellTravelType(spellTravelType, user)
	if err != nil {
		return
	}
	response := &SpellTravelTypeResponse{
		SpellTravelType: spellTravelType,
	}
	content = response
	return
}

func listSpellTravelType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellTravelTypes, err := cases.ListSpellTravelType(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellTravelTypesResponse{
		Page:             page,
		SpellTravelTypes: spellTravelTypes,
	}
	content = response
	return
}

func listSpellTravelTypeBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellTravelType := &model.SpellTravelType{
		Name: getQuery(r, "name"),
	}
	spellTravelTypes, err := cases.ListSpellTravelTypeBySearch(page, spellTravelType, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellTravelTypesBySearchResponse{
		Page:             page,
		SpellTravelTypes: spellTravelTypes,
		Search:           spellTravelType,
	}
	content = response
	return
}
