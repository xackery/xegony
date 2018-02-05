package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellEffectFormulaRequest is a list of parameters used for spellEffectFormula
// swagger:parameters deleteSpellEffectFormula editSpellEffectFormula getSpellEffectFormula
type SpellEffectFormulaRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// SpellEffectFormulaResponse is what endpoints respond with
// swagger:response
type SpellEffectFormulaResponse struct {
	SpellEffectFormula *model.SpellEffectFormula `json:"spellEffectFormula,omitempty"`
}

// SpellEffectFormulaCreateRequest is the body parameters for creating an spellEffectFormula
// swagger:parameters createSpellEffectFormula
type SpellEffectFormulaCreateRequest struct {
	// SpellEffectFormula details to create
	// in: body
	SpellEffectFormula *model.SpellEffectFormula `json:"spellEffectFormula"`
}

// SpellEffectFormulaEditRequest is the body parameters for creating an spellEffectFormula
// swagger:parameters editSpellEffectFormula
type SpellEffectFormulaEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// SpellEffectFormula details to edit
	// in: body
	SpellEffectFormula *model.SpellEffectFormula `json:"spellEffectFormula"`
}

// SpellEffectFormulasRequest is a list of parameters used for spellEffectFormula
// swagger:parameters listSpellEffectFormula
type SpellEffectFormulasRequest struct {
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

// SpellEffectFormulasResponse is a general response to a request
// swagger:response
type SpellEffectFormulasResponse struct {
	Page                *model.Page               `json:"page,omitempty"`
	SpellEffectFormulas model.SpellEffectFormulas `json:"spellEffectFormulas,omitempty"`
}

// SpellEffectFormulasBySearchRequest is a list of parameters used for spellEffectFormula
// swagger:parameters listSpellEffectFormulaBySearch
type SpellEffectFormulasBySearchRequest struct {
	// Name is which spellEffectFormula to get information about
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

// SpellEffectFormulasBySearchResponse is a general response to a request
// swagger:response
type SpellEffectFormulasBySearchResponse struct {
	Search              *model.SpellEffectFormula `json:"search,omitempty"`
	Page                *model.Page               `json:"page,omitempty"`
	SpellEffectFormulas model.SpellEffectFormulas `json:"spellEffectFormulas,omitempty"`
}

func spellEffectFormulaRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell/effect/formula spellEffectFormula listSpellEffectFormula
		//
		// Lists spellEffectFormulas
		//
		// This will show all available spellEffectFormulas by default.
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
		//       200: SpellEffectFormulasResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellEffectFormula",
			"GET",
			"/spell/effect/formula",
			listSpellEffectFormula,
		},
		// swagger:route GET /spell/effect/formula/search spellEffectFormula listSpellEffectFormulaBySearch
		//
		// Search spellEffectFormulas by name
		//
		// This will show all available spellEffectFormulas by default.
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
		//       200: SpellEffectFormulasBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellEffectFormulaBySearch",
			"GET",
			"/spell/effect/formula/search",
			listSpellEffectFormulaBySearch,
		},
		// swagger:route POST /spell/effect/formula spellEffectFormula createSpellEffectFormula
		//
		// Create an spellEffectFormula
		//
		// This will create an spellEffectFormula
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellEffectFormulaResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpellEffectFormula",
			"POST",
			"/spell/effect/formula",
			createSpellEffectFormula,
		},
		// swagger:route GET /spell/effect/formula/{ID} spellEffectFormula getSpellEffectFormula
		//
		// Get an spellEffectFormula
		//
		// This will get an individual spellEffectFormula available spellEffectFormulas by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellEffectFormulaResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpellEffectFormula",
			"GET",
			"/spell/effect/formula/{ID:[0-9]+}",
			getSpellEffectFormula,
		},
		// swagger:route PUT /spell/effect/formula/{ID} spellEffectFormula editSpellEffectFormula
		//
		// Edit an spellEffectFormula
		//
		// This will edit an spellEffectFormula
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellEffectFormulaResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpellEffectFormula",
			"PUT",
			"/spell/effect/formula/{ID:[0-9]+}",
			editSpellEffectFormula,
		},
		// swagger:route DELETE /spell/effect/formula/{ID} spellEffectFormula deleteSpellEffectFormula
		//
		// Delete an spellEffectFormula
		//
		// This will delete an spellEffectFormula
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
			"DeleteSpellEffectFormula",
			"DELETE",
			"/spell/effect/formula/{ID:[0-9]+}",
			deleteSpellEffectFormula,
		},
	}
	return
}

func getSpellEffectFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellEffectFormulaRequest{
		ID: getIntVar(r, "ID"),
	}

	spellEffectFormula := &model.SpellEffectFormula{
		ID: request.ID,
	}

	err = cases.GetSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellEffectFormulaResponse{
		SpellEffectFormula: spellEffectFormula,
	}
	content = response
	return
}

func createSpellEffectFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spellEffectFormula := &model.SpellEffectFormula{}
	err = decodeBody(r, spellEffectFormula)
	if err != nil {
		return
	}
	err = cases.CreateSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		return
	}
	response := &SpellEffectFormulaResponse{
		SpellEffectFormula: spellEffectFormula,
	}
	content = response
	return
}

func deleteSpellEffectFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellEffectFormulaRequest{
		ID: getIntVar(r, "ID"),
	}

	spellEffectFormula := &model.SpellEffectFormula{
		ID: request.ID,
	}

	err = cases.DeleteSpellEffectFormula(spellEffectFormula, user)
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

func editSpellEffectFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellEffectFormulaEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spellEffectFormula := &model.SpellEffectFormula{
		ID: request.ID,
	}

	err = decodeBody(r, spellEffectFormula)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpellEffectFormula(spellEffectFormula, user)
	if err != nil {
		return
	}
	response := &SpellEffectFormulaResponse{
		SpellEffectFormula: spellEffectFormula,
	}
	content = response
	return
}

func listSpellEffectFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellEffectFormulas, err := cases.ListSpellEffectFormula(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellEffectFormulasResponse{
		Page:                page,
		SpellEffectFormulas: spellEffectFormulas,
	}
	content = response
	return
}

func listSpellEffectFormulaBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellEffectFormula := &model.SpellEffectFormula{
		Name: getQuery(r, "name"),
	}
	spellEffectFormulas, err := cases.ListSpellEffectFormulaBySearch(page, spellEffectFormula, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellEffectFormulasBySearchResponse{
		Page:                page,
		SpellEffectFormulas: spellEffectFormulas,
		Search:              spellEffectFormula,
	}
	content = response
	return
}
