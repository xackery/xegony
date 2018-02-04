package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellAnimationTypeRequest is a list of parameters used for spellAnimationType
// swagger:parameters deleteSpellAnimationType editSpellAnimationType getSpellAnimationType
type SpellAnimationTypeRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// SpellAnimationTypeResponse is what endpoints respond with
// swagger:response
type SpellAnimationTypeResponse struct {
	SpellAnimationType *model.SpellAnimationType `json:"spellAnimationType,omitempty"`
}

// SpellAnimationTypeCreateRequest is the body parameters for creating an spellAnimationType
// swagger:parameters createSpellAnimationType
type SpellAnimationTypeCreateRequest struct {
	// SpellAnimationType details to create
	// in: body
	SpellAnimationType *model.SpellAnimationType `json:"spellAnimationType"`
}

// SpellAnimationTypeEditRequest is the body parameters for creating an spellAnimationType
// swagger:parameters editSpellAnimationType
type SpellAnimationTypeEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// SpellAnimationType details to edit
	// in: body
	SpellAnimationType *model.SpellAnimationType `json:"spellAnimationType"`
}

// SpellAnimationTypesRequest is a list of parameters used for spellAnimationType
// swagger:parameters listSpellAnimationType
type SpellAnimationTypesRequest struct {
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

// SpellAnimationTypesResponse is a general response to a request
// swagger:response
type SpellAnimationTypesResponse struct {
	Page                *model.Page               `json:"page,omitempty"`
	SpellAnimationTypes model.SpellAnimationTypes `json:"spellAnimationTypes,omitempty"`
}

// SpellAnimationTypesBySearchRequest is a list of parameters used for spellAnimationType
// swagger:parameters listSpellAnimationTypeBySearch
type SpellAnimationTypesBySearchRequest struct {
	// Name is which spellAnimationType to get information about
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

// SpellAnimationTypesBySearchResponse is a general response to a request
// swagger:response
type SpellAnimationTypesBySearchResponse struct {
	Search              *model.SpellAnimationType `json:"search,omitempty"`
	Page                *model.Page               `json:"page,omitempty"`
	SpellAnimationTypes model.SpellAnimationTypes `json:"spellAnimationTypes,omitempty"`
}

func spellAnimationTypeRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell/animation/type spellAnimationType listSpellAnimationType
		//
		// Lists spellAnimationTypes
		//
		// This will show all available spellAnimationTypes by default.
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
		//       200: SpellAnimationTypesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellAnimationType",
			"GET",
			"/spell/animation/type",
			listSpellAnimationType,
		},
		// swagger:route GET /spell/animation/type/search spellAnimationType listSpellAnimationTypeBySearch
		//
		// Search spellAnimationTypes by name
		//
		// This will show all available spellAnimationTypes by default.
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
		//       200: SpellAnimationTypesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellAnimationTypeBySearch",
			"GET",
			"/spell/animation/type/search",
			listSpellAnimationTypeBySearch,
		},
		// swagger:route POST /spell/animation/type spellAnimationType createSpellAnimationType
		//
		// Create an spellAnimationType
		//
		// This will create an spellAnimationType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellAnimationTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpellAnimationType",
			"POST",
			"/spell/animation/type",
			createSpellAnimationType,
		},
		// swagger:route GET /spell/animation/type/{ID} spellAnimationType getSpellAnimationType
		//
		// Get an spellAnimationType
		//
		// This will get an individual spellAnimationType available spellAnimationTypes by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellAnimationTypeResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpellAnimationType",
			"GET",
			"/spell/animation/type/{ID:[0-9]+}",
			getSpellAnimationType,
		},
		// swagger:route PUT /spell/animation/type/{ID} spellAnimationType editSpellAnimationType
		//
		// Edit an spellAnimationType
		//
		// This will edit an spellAnimationType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellAnimationTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpellAnimationType",
			"PUT",
			"/spell/animation/type/{ID:[0-9]+}",
			editSpellAnimationType,
		},
		// swagger:route DELETE /spell/animation/type/{ID} spellAnimationType deleteSpellAnimationType
		//
		// Delete an spellAnimationType
		//
		// This will delete an spellAnimationType
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
			"DeleteSpellAnimationType",
			"DELETE",
			"/spell/animation/type/{ID:[0-9]+}",
			deleteSpellAnimationType,
		},
	}
	return
}

func getSpellAnimationType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellAnimationTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellAnimationType := &model.SpellAnimationType{
		ID: request.ID,
	}

	err = cases.GetSpellAnimationType(spellAnimationType, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellAnimationTypeResponse{
		SpellAnimationType: spellAnimationType,
	}
	content = response
	return
}

func createSpellAnimationType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spellAnimationType := &model.SpellAnimationType{}
	err = decodeBody(r, spellAnimationType)
	if err != nil {
		return
	}
	err = cases.CreateSpellAnimationType(spellAnimationType, user)
	if err != nil {
		return
	}
	response := &SpellAnimationTypeResponse{
		SpellAnimationType: spellAnimationType,
	}
	content = response
	return
}

func deleteSpellAnimationType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellAnimationTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	spellAnimationType := &model.SpellAnimationType{
		ID: request.ID,
	}

	err = cases.DeleteSpellAnimationType(spellAnimationType, user)
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

func editSpellAnimationType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellAnimationTypeEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spellAnimationType := &model.SpellAnimationType{
		ID: request.ID,
	}

	err = decodeBody(r, spellAnimationType)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpellAnimationType(spellAnimationType, user)
	if err != nil {
		return
	}
	response := &SpellAnimationTypeResponse{
		SpellAnimationType: spellAnimationType,
	}
	content = response
	return
}

func listSpellAnimationType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellAnimationTypes, err := cases.ListSpellAnimationType(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellAnimationTypesResponse{
		Page:                page,
		SpellAnimationTypes: spellAnimationTypes,
	}
	content = response
	return
}

func listSpellAnimationTypeBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellAnimationType := &model.SpellAnimationType{
		Name: getQuery(r, "name"),
	}
	spellAnimationTypes, err := cases.ListSpellAnimationTypeBySearch(page, spellAnimationType, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellAnimationTypesBySearchResponse{
		Page:                page,
		SpellAnimationTypes: spellAnimationTypes,
		Search:              spellAnimationType,
	}
	content = response
	return
}
