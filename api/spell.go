package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SpellRequest is a list of parameters used for spell
// swagger:parameters deleteSpell editSpell getSpell
type SpellRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
}

// SpellResponse is what endpoints respond with
// swagger:response
type SpellResponse struct {
	Spell *model.Spell `json:"spell,omitempty"`
}

// SpellCreateRequest is the body parameters for creating an spell
// swagger:parameters createSpell
type SpellCreateRequest struct {
	// Spell details to create
	// in: body
	Spell *model.Spell `json:"spell"`
}

// SpellEditRequest is the body parameters for creating an spell
// swagger:parameters editSpell
type SpellEditRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
	// Spell details to edit
	// in: body
	Spell *model.Spell `json:"spell"`
}

// SpellsRequest is a list of parameters used for spell
// swagger:parameters listSpell
type SpellsRequest struct {
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

// SpellsResponse is a general response to a request
// swagger:response
type SpellsResponse struct {
	Page   *model.Page  `json:"page,omitempty"`
	Spells model.Spells `json:"spells,omitempty"`
}

// SpellsBySearchRequest is a list of parameters used for spell
// swagger:parameters listSpellBySearch
type SpellsBySearchRequest struct {
	// Name is which spell to get information about
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

// SpellsBySearchResponse is a general response to a request
// swagger:response
type SpellsBySearchResponse struct {
	Search *model.Spell `json:"search,omitempty"`
	Page   *model.Page  `json:"page,omitempty"`
	Spells model.Spells `json:"spells,omitempty"`
}

func spellRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spell spell listSpell
		//
		// Lists spells
		//
		// This will show all available spells by default.
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
		//       200: SpellsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpell",
			"GET",
			"/spell",
			listSpell,
		},
		// swagger:route GET /spell/search spell listSpellBySearch
		//
		// Search spells by name
		//
		// This will show all available spells by default.
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
		//       200: SpellsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpellBySearch",
			"GET",
			"/spell/search",
			listSpellBySearch,
		},
		// swagger:route POST /spell spell createSpell
		//
		// Create an spell
		//
		// This will create an spell
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSpell",
			"POST",
			"/spell",
			createSpell,
		},
		// swagger:route GET /spell/{ID} spell getSpell
		//
		// Get an spell
		//
		// This will get an individual spell available spells by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpellResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpell",
			"GET",
			"/spell/{ID:[0-9]+}",
			getSpell,
		},
		// swagger:route PUT /spell/{ID} spell editSpell
		//
		// Edit an spell
		//
		// This will edit an spell
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SpellResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSpell",
			"PUT",
			"/spell/{ID:[0-9]+}",
			editSpell,
		},
		// swagger:route DELETE /spell/{ID} spell deleteSpell
		//
		// Delete an spell
		//
		// This will delete an spell
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
			"DeleteSpell",
			"DELETE",
			"/spell/{ID:[0-9]+}",
			deleteSpell,
		},
	}
	return
}

func getSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellRequest{
		ID: getIntVar(r, "ID"),
	}

	spell := &model.Spell{
		ID: request.ID,
	}

	err = cases.GetSpell(spell, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SpellResponse{
		Spell: spell,
	}
	content = response
	return
}

func createSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spell := &model.Spell{}
	err = decodeBody(r, spell)
	if err != nil {
		return
	}
	err = cases.CreateSpell(spell, user)
	if err != nil {
		return
	}
	response := &SpellResponse{
		Spell: spell,
	}
	content = response
	return
}

func deleteSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellRequest{
		ID: getIntVar(r, "ID"),
	}

	spell := &model.Spell{
		ID: request.ID,
	}

	err = cases.DeleteSpell(spell, user)
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

func editSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SpellEditRequest{
		ID: getIntVar(r, "ID"),
	}

	spell := &model.Spell{
		ID: request.ID,
	}

	err = decodeBody(r, spell)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSpell(spell, user)
	if err != nil {
		return
	}
	response := &SpellResponse{
		Spell: spell,
	}
	content = response
	return
}

func listSpell(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spells, err := cases.ListSpell(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellsResponse{
		Page:   page,
		Spells: spells,
	}
	content = response
	return
}

func listSpellBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	spell := &model.Spell{}
	spell.Name.String = getQuery(r, "name")
	spells, err := cases.ListSpellBySearch(page, spell, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SpellsBySearchResponse{
		Page:   page,
		Spells: spells,
		Search: spell,
	}
	content = response
	return
}
