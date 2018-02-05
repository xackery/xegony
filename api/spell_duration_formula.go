package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellDurationFormulaRequest is a list of parameters used for spellDurationFormula
// swagger:parameters deleteSpellDurationFormula editSpellDurationFormula getSpellDurationFormula
type SpellDurationFormulaRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// SpellDurationFormulaResponse is what endpoints respond with
// swagger:response
type SpellDurationFormulaResponse struct {
	SpellDurationFormula *model.SpellDurationFormula `json:"spellDurationFormula,omitempty"`
}

// SpellDurationFormulaCreateRequest is the body parameters for creating an spellDurationFormula
// swagger:parameters createSpellDurationFormula
type SpellDurationFormulaCreateRequest struct {
	// SpellDurationFormula details to create
	// in: body
	SpellDurationFormula *model.SpellDurationFormula `json:"spellDurationFormula"`
}

// SpellDurationFormulaEditRequest is the body parameters for creating an spellDurationFormula
// swagger:parameters editSpellDurationFormula
type SpellDurationFormulaEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// SpellDurationFormula details to edit
	// in: body
	SpellDurationFormula *model.SpellDurationFormula `json:"spellDurationFormula"`
}

// SpellDurationFormulasRequest is a list of parameters used for spellDurationFormula
// swagger:parameters listSpellDurationFormula
type SpellDurationFormulasRequest struct {
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

// SpellDurationFormulasResponse is a general response to a request
// swagger:response
type SpellDurationFormulasResponse struct {
	Page                  *model.Page                 `json:"page,omitempty"`
	SpellDurationFormulas model.SpellDurationFormulas `json:"spellDurationFormulas,omitempty"`
}

// SpellDurationFormulasBySearchRequest is a list of parameters used for spellDurationFormula
// swagger:parameters listSpellDurationFormulaBySearch
type SpellDurationFormulasBySearchRequest struct {
	// Name is which spellDurationFormula to get information about
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

// SpellDurationFormulasBySearchResponse is a general response to a request
// swagger:response
type SpellDurationFormulasBySearchResponse struct {
	Search                *model.SpellDurationFormula `json:"search,omitempty"`
	Page                  *model.Page                 `json:"page,omitempty"`
	SpellDurationFormulas model.SpellDurationFormulas `json:"spellDurationFormulas,omitempty"`
}

func spellDurationFormulaRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell/effect/formula spellDurationFormula listSpellDurationFormula
		//
		// Lists spellDurationFormulas
		//
		// This will show all available spellDurationFormulas by default.
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
		//       200: SpellDurationFormulasResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellDurationFormula",
			"GET",
			"/spell/effect/formula",
			listSpellDurationFormula,
		},
		// swagger:route GET /spell/effect/formula/search spellDurationFormula listSpellDurationFormulaBySearch
		//
		// Search spellDurationFormulas by name
		//
		// This will show all available spellDurationFormulas by default.
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
		//       200: SpellDurationFormulasBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellDurationFormulaBySearch",
			"GET",
			"/spell/effect/formula/search",
			listSpellDurationFormulaBySearch,
		},
		// swagger:route POST /spell/effect/formula spellDurationFormula createSpellDurationFormula
		//
		// Create an spellDurationFormula
		//
		// This will create an spellDurationFormula
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellDurationFormulaResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpellDurationFormula",
			"POST",
			"/spell/effect/formula",
			createSpellDurationFormula,
		},
		// swagger:route GET /spell/effect/formula/{ID} spellDurationFormula getSpellDurationFormula
		//
		// Get an spellDurationFormula
		//
		// This will get an individual spellDurationFormula available spellDurationFormulas by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellDurationFormulaResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpellDurationFormula",
			"GET",
			"/spell/effect/formula/{ID:[0-9]+}",
			getSpellDurationFormula,
		},
		// swagger:route PUT /spell/effect/formula/{ID} spellDurationFormula editSpellDurationFormula
		//
		// Edit an spellDurationFormula
		//
		// This will edit an spellDurationFormula
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellDurationFormulaResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpellDurationFormula",
			"PUT",
			"/spell/effect/formula/{ID:[0-9]+}",
			editSpellDurationFormula,
		},
		// swagger:route DELETE /spell/effect/formula/{ID} spellDurationFormula deleteSpellDurationFormula
		//
		// Delete an spellDurationFormula
		//
		// This will delete an spellDurationFormula
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
			"DeleteSpellDurationFormula",
			"DELETE",
			"/spell/effect/formula/{ID:[0-9]+}",
			deleteSpellDurationFormula,
		},
	}
	return
}

func getSpellDurationFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellDurationFormulaRequest{
		ID: getIntVar(r, "ID"),
	}

	spellDurationFormula := &model.SpellDurationFormula{
		ID: request.ID,
	}

	err = cases.GetSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellDurationFormulaResponse{
		SpellDurationFormula: spellDurationFormula,
	}
	content = response
	return
}

func createSpellDurationFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spellDurationFormula := &model.SpellDurationFormula{}
	err = decodeBody(r, spellDurationFormula)
	if err != nil {
		return
	}
	err = cases.CreateSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		return
	}
	response := &SpellDurationFormulaResponse{
		SpellDurationFormula: spellDurationFormula,
	}
	content = response
	return
}

func deleteSpellDurationFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellDurationFormulaRequest{
		ID: getIntVar(r, "ID"),
	}

	spellDurationFormula := &model.SpellDurationFormula{
		ID: request.ID,
	}

	err = cases.DeleteSpellDurationFormula(spellDurationFormula, user)
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

func editSpellDurationFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellDurationFormulaEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spellDurationFormula := &model.SpellDurationFormula{
		ID: request.ID,
	}

	err = decodeBody(r, spellDurationFormula)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpellDurationFormula(spellDurationFormula, user)
	if err != nil {
		return
	}
	response := &SpellDurationFormulaResponse{
		SpellDurationFormula: spellDurationFormula,
	}
	content = response
	return
}

func listSpellDurationFormula(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellDurationFormulas, err := cases.ListSpellDurationFormula(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellDurationFormulasResponse{
		Page: page,
		SpellDurationFormulas: spellDurationFormulas,
	}
	content = response
	return
}

func listSpellDurationFormulaBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spellDurationFormula := &model.SpellDurationFormula{
		Name: getQuery(r, "name"),
	}
	spellDurationFormulas, err := cases.ListSpellDurationFormulaBySearch(page, spellDurationFormula, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellDurationFormulasBySearchResponse{
		Page: page,
		SpellDurationFormulas: spellDurationFormulas,
		Search:                spellDurationFormula,
	}
	content = response
	return
}
