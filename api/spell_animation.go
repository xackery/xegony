package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellAnimationRequest is a list of parameters used for spellAnimation
// swagger:parameters deleteSpellAnimation editSpellAnimation getSpellAnimation
type SpellAnimationRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// SpellAnimationResponse is what endpoints respond with
// swagger:response
type SpellAnimationResponse struct {
	SpellAnimation *model.SpellAnimation `json:"spellAnimation,omitempty"`
}

// SpellAnimationCreateRequest is the body parameters for creating an spellAnimation
// swagger:parameters createSpellAnimation
type SpellAnimationCreateRequest struct {
	// SpellAnimation details to create
	// in: body
	SpellAnimation *model.SpellAnimation `json:"spellAnimation"`
}

// SpellAnimationEditRequest is the body parameters for creating an spellAnimation
// swagger:parameters editSpellAnimation
type SpellAnimationEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// SpellAnimation details to edit
	// in: body
	SpellAnimation *model.SpellAnimation `json:"spellAnimation"`
}

// SpellAnimationsRequest is a list of parameters used for spellAnimation
// swagger:parameters listSpellAnimation
type SpellAnimationsRequest struct {
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

// SpellAnimationsResponse is a general response to a request
// swagger:response
type SpellAnimationsResponse struct {
	Page            *model.Page           `json:"page,omitempty"`
	SpellAnimations model.SpellAnimations `json:"spellAnimations,omitempty"`
}

// SpellAnimationsBySearchRequest is a list of parameters used for spellAnimation
// swagger:parameters listSpellAnimationBySearch
type SpellAnimationsBySearchRequest struct {
	// Name is which spellAnimation to get information about
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

// SpellAnimationsBySearchResponse is a general response to a request
// swagger:response
type SpellAnimationsBySearchResponse struct {
	Search          *model.SpellAnimation `json:"search,omitempty"`
	Page            *model.Page           `json:"page,omitempty"`
	SpellAnimations model.SpellAnimations `json:"spellAnimations,omitempty"`
}

func spellAnimationRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell/animation spellAnimation listSpellAnimation
		//
		// Lists spellAnimations
		//
		// This will show all available spellAnimations by default.
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
		//       200: SpellAnimationsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellAnimation",
			"GET",
			"/spell/animation",
			listSpellAnimation,
		},
		// swagger:route GET /spell/animation/search spellAnimation listSpellAnimationBySearch
		//
		// Search spellAnimations by name
		//
		// This will show all available spellAnimations by default.
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
		//       200: SpellAnimationsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellAnimationBySearch",
			"GET",
			"/spell/animation/search",
			listSpellAnimationBySearch,
		},
		// swagger:route POST /spell/animation spellAnimation createSpellAnimation
		//
		// Create an spellAnimation
		//
		// This will create an spellAnimation
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellAnimationResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpellAnimation",
			"POST",
			"/spell/animation",
			createSpellAnimation,
		},
		// swagger:route GET /spell/animation/{ID} spellAnimation getSpellAnimation
		//
		// Get an spellAnimation
		//
		// This will get an individual spellAnimation available spellAnimations by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellAnimationResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpellAnimation",
			"GET",
			"/spell/animation/{ID:[0-9]+}",
			getSpellAnimation,
		},
		// swagger:route PUT /spell/animation/{ID} spellAnimation editSpellAnimation
		//
		// Edit an spellAnimation
		//
		// This will edit an spellAnimation
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellAnimationResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpellAnimation",
			"PUT",
			"/spell/animation/{ID:[0-9]+}",
			editSpellAnimation,
		},
		// swagger:route DELETE /spell/animation/{ID} spellAnimation deleteSpellAnimation
		//
		// Delete an spellAnimation
		//
		// This will delete an spellAnimation
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
			"DeleteSpellAnimation",
			"DELETE",
			"/spell/animation/{ID:[0-9]+}",
			deleteSpellAnimation,
		},
	}
	return
}

func getSpellAnimation(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellAnimationRequest{
		ID: getIntVar(r, "ID"),
	}

	spellAnimation := &model.SpellAnimation{
		ID: request.ID,
	}

	err = cases.GetSpellAnimation(spellAnimation, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellAnimationResponse{
		SpellAnimation: spellAnimation,
	}
	content = response
	return
}

func createSpellAnimation(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spellAnimation := &model.SpellAnimation{}
	err = decodeBody(r, spellAnimation)
	if err != nil {
		return
	}
	err = cases.CreateSpellAnimation(spellAnimation, user)
	if err != nil {
		return
	}
	response := &SpellAnimationResponse{
		SpellAnimation: spellAnimation,
	}
	content = response
	return
}

func deleteSpellAnimation(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellAnimationRequest{
		ID: getIntVar(r, "ID"),
	}

	spellAnimation := &model.SpellAnimation{
		ID: request.ID,
	}

	err = cases.DeleteSpellAnimation(spellAnimation, user)
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

func editSpellAnimation(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellAnimationEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spellAnimation := &model.SpellAnimation{
		ID: request.ID,
	}

	err = decodeBody(r, spellAnimation)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpellAnimation(spellAnimation, user)
	if err != nil {
		return
	}
	response := &SpellAnimationResponse{
		SpellAnimation: spellAnimation,
	}
	content = response
	return
}

func listSpellAnimation(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellAnimations, err := cases.ListSpellAnimation(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellAnimationsResponse{
		Page:            page,
		SpellAnimations: spellAnimations,
	}
	content = response
	return
}

func listSpellAnimationBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellAnimation := &model.SpellAnimation{
		Name: getQuery(r, "name"),
	}
	spellAnimations, err := cases.ListSpellAnimationBySearch(page, spellAnimation, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	log.Println(spellAnimations)
	response := &SpellAnimationsBySearchResponse{
		Page:            page,
		SpellAnimations: spellAnimations,
		Search:          spellAnimation,
	}
	content = response
	return
}
