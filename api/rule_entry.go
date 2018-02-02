package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// RuleEntryRequest is a list of parameters used for ruleEntry
// swagger:parameters deleteRuleEntry editRuleEntry getRuleEntry
type RuleEntryRequest struct {
	// RuleID to get information about
	// in: path
	// example: 1
	RuleID int64 `json:"ruleID"`
	// Name to get information about
	// in: path
	// example: AA:ExpPerPoint
	Name string `json:"name"`
}

// RuleEntryResponse is what endpoints respond with
// swagger:response
type RuleEntryResponse struct {
	Rule      *model.Rule      `json:"rule,omitempty"`
	RuleEntry *model.RuleEntry `json:"ruleEntry,omitempty"`
}

// RuleEntryCreateRequest is the body parameters for creating an ruleEntry
// swagger:parameters createRuleEntry
type RuleEntryCreateRequest struct {
	// RuleID to get information about
	// in: path
	// example: 1
	RuleID int64 `json:"ruleID"`
	// RuleEntry details to create
	// in: body
	RuleEntry *model.RuleEntry `json:"ruleEntry"`
}

// RuleEntryEditRequest is the body parameters for creating an ruleEntry
// swagger:parameters editRuleEntry
type RuleEntryEditRequest struct {
	// RuleID to get information about
	// in: path
	// example: 1
	RuleID int64 `json:"RuleID"`
	// name to get information about
	// in: path
	// example: AA:ExpPerPoint
	Name string `json:"name"`
	// RuleEntry details to edit
	// in: body
	RuleEntry *model.RuleEntry `json:"ruleEntry"`
}

// RuleEntrysRequest is a list of parameters used for ruleEntry
// swagger:parameters listRuleEntry
type RuleEntrysRequest struct {
	// RuleID to get information about
	// in: path
	// example: 1
	RuleID int64 `json:"ruleID"`
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

// RuleEntrysResponse is a general response to a request
// swagger:response
type RuleEntrysResponse struct {
	Page       *model.Page      `json:"page,omitempty"`
	Rule       *model.Rule      `json:"rule"`
	RuleEntrys model.RuleEntrys `json:"ruleEntrys,omitempty"`
}

// RuleEntrysBySearchRequest is a list of parameters used for ruleEntry
// swagger:parameters listRuleEntryBySearch
type RuleEntrysBySearchRequest struct {
	// RuleID to get information about
	// in: path
	// example: 1
	RuleID int64 `json:"ruleID"`
	// Name is which ruleEntry to get information about
	// example: AA:ExpPerPoint
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

// RuleEntrysBySearchResponse is a general response to a request
// swagger:response
type RuleEntrysBySearchResponse struct {
	Search     *model.RuleEntry `json:"search,omitempty"`
	Page       *model.Page      `json:"page,omitempty"`
	Rule       *model.Rule      `json:"rule,omitempty"`
	RuleEntrys model.RuleEntrys `json:"ruleEntrys,omitempty"`
}

func ruleEntryRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /rule/{ruleID}/entry rule listRuleEntry
		//
		// Lists ruleEntrys
		//
		// This will show all available ruleEntrys by default.
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
		//       200: RuleEntrysResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListRuleEntry",
			"GET",
			"/rule/{ruleID:[0-9]+}/entry",
			listRuleEntry,
		},
		// swagger:route GET /rule/{ruleID}/entry/search rule listRuleEntryBySearch
		//
		// Search ruleEntrys by name
		//
		// This will show all available ruleEntrys by default.
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
		//       200: RuleEntrysBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListRuleEntryBySearch",
			"GET",
			"/rule/{ruleID:[0-9]+}/entry/search",
			listRuleEntryBySearch,
		},
		// swagger:route POST /rule/{ruleID}/entry/{Name} rule createRuleEntry
		//
		// Create an ruleEntry
		//
		// This will create an ruleEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: RuleEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateRuleEntry",
			"POST",
			"/rule/{ruleID:[0-9]+}/entry/{name:[a-zA-Z]+:[a-zA-Z]+}",
			createRuleEntry,
		},
		// swagger:route GET /rule/{ruleID}/entry/{Name} rule getRuleEntry
		//
		// Get an ruleEntry
		//
		// This will get an individual ruleEntry available ruleEntrys by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: RuleEntryResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetRuleEntry",
			"GET",
			"/rule/{ruleID:[0-9]+}/entry/{name:[a-zA-Z]+:[a-zA-Z]+}",
			getRuleEntry,
		},
		// swagger:route PUT /rule/{ruleID}/entry/{Name} rule editRuleEntry
		//
		// Edit an ruleEntry
		//
		// This will edit an ruleEntry
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: RuleEntryResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditRuleEntry",
			"PUT",
			"/rule/{ruleID:[0-9]+}/entry/{name:[a-zA-Z]+:[a-zA-Z]+}",
			editRuleEntry,
		},
		// swagger:route DELETE /rule/{ruleID}/entry/{Name} rule deleteRuleEntry
		//
		// Delete an ruleEntry
		//
		// This will delete an ruleEntry
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
			"DeleteRuleEntry",
			"DELETE",
			"/rule/{ruleID:[0-9]+}/entry/{name:[a-zA-Z]+:[a-zA-Z]+}",
			deleteRuleEntry,
		},
	}
	return
}

func getRuleEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleEntryRequest{
		RuleID: getIntVar(r, "ruleID"),
		Name:   getVar(r, "name"),
	}

	rule := &model.Rule{
		ID: request.RuleID,
	}

	err = cases.GetRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}

	ruleEntry := &model.RuleEntry{
		RuleID: request.RuleID,
		Name:   request.Name,
	}

	err = cases.GetRuleEntry(rule, ruleEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &RuleEntryResponse{
		Rule:      rule,
		RuleEntry: ruleEntry,
	}
	content = response
	return
}

func createRuleEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleEntryCreateRequest{
		RuleID: getIntVar(r, "ruleID"),
	}

	rule := &model.Rule{
		ID: request.RuleID,
	}

	err = cases.GetRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}

	ruleEntry := &model.RuleEntry{}
	err = decodeBody(r, ruleEntry)
	if err != nil {
		return
	}

	err = cases.CreateRuleEntry(rule, ruleEntry, user)
	if err != nil {
		return
	}
	response := &RuleEntryResponse{
		Rule:      rule,
		RuleEntry: ruleEntry,
	}
	content = response
	return
}

func deleteRuleEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleEntryRequest{
		RuleID: getIntVar(r, "ruleID"),
		Name:   getVar(r, "name"),
	}

	rule := &model.Rule{
		ID: request.RuleID,
	}

	err = cases.GetRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}

	ruleEntry := &model.RuleEntry{
		RuleID: request.RuleID,
		Name:   request.Name,
	}

	err = cases.DeleteRuleEntry(ruleEntry, rule, user)
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

func editRuleEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleEntryEditRequest{
		RuleID: getIntVar(r, "ruleID"),
		Name:   getVar(r, "name"),
	}

	rule := &model.Rule{
		ID:   request.RuleID,
		Name: request.Name,
	}

	err = cases.GetRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}

	ruleEntry := &model.RuleEntry{}

	err = decodeBody(r, ruleEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	ruleEntry.Name = request.Name
	ruleEntry.RuleID = request.RuleID

	err = cases.EditRuleEntry(rule, ruleEntry, user)
	if err != nil {
		return
	}
	response := &RuleEntryResponse{
		Rule:      rule,
		RuleEntry: ruleEntry,
	}
	content = response
	return
}

func listRuleEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleEntrysRequest{
		RuleID: getIntVar(r, "ruleID"),
	}

	rule := &model.Rule{
		ID: request.RuleID,
	}

	err = cases.GetRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	ruleEntrys, err := cases.ListRuleEntry(page, rule, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &RuleEntrysResponse{
		Page:       page,
		Rule:       rule,
		RuleEntrys: ruleEntrys,
	}
	content = response
	return
}

func listRuleEntryBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleEntrysBySearchRequest{
		RuleID: getIntVar(r, "ruleID"),
	}

	rule := &model.Rule{
		ID: request.RuleID,
	}

	err = cases.GetRule(rule, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get rule")
		return
	}

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}

	ruleEntry := &model.RuleEntry{
		RuleID: request.RuleID,
	}
	ruleEntry.Name = getQuery(r, "name")

	ruleEntrys, err := cases.ListRuleEntryBySearch(page, rule, ruleEntry, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &RuleEntrysBySearchResponse{
		Page:       page,
		Rule:       rule,
		RuleEntrys: ruleEntrys,
		Search:     ruleEntry,
	}
	content = response
	return
}
