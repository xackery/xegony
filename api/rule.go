package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// RuleRequest is a list of parameters used for rule
// swagger:parameters deleteRule editRule getRule
type RuleRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// RuleResponse is what endpoints respond with
// swagger:response
type RuleResponse struct {
	Rule *model.Rule `json:"rule,omitempty"`
}

// RuleCreateRequest is the body parameters for creating an rule
// swagger:parameters createRule
type RuleCreateRequest struct {
	// Rule details to create
	// in: body
	Rule *model.Rule `json:"rule"`
}

// RuleEditRequest is the body parameters for creating an rule
// swagger:parameters editRule
type RuleEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Rule details to edit
	// in: body
	Rule *model.Rule `json:"rule"`
}

// RulesRequest is a list of parameters used for rule
// swagger:parameters listRule
type RulesRequest struct {
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

// RulesResponse is a general response to a request
// swagger:response
type RulesResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Rules model.Rules `json:"rules,omitempty"`
}

// RulesBySearchRequest is a list of parameters used for rule
// swagger:parameters listRuleBySearch
type RulesBySearchRequest struct {
	// ShortName is which rule to get information about
	// example: xackery
	// in: query
	ShortName string `json:"shortName"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: short_name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// RulesBySearchResponse is a general response to a request
// swagger:response
type RulesBySearchResponse struct {
	Search *model.Rule `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Rules  model.Rules `json:"rules,omitempty"`
}

func ruleRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /rule rule listRule
		//
		// Lists rules
		//
		// This will show all available rules by default.
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
		//       200: RulesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListRule",
			"GET",
			"/rule",
			listRule,
		},
		// swagger:route GET /rule/search rule listRuleBySearch
		//
		// Search rules by name
		//
		// This will show all available rules by default.
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
		//       200: RulesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListRuleBySearch",
			"GET",
			"/rule/search",
			listRuleBySearch,
		},
		// swagger:route POST /rule rule createRule
		//
		// Create an rule
		//
		// This will create an rule
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: RuleResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateRule",
			"POST",
			"/rule",
			createRule,
		},
		// swagger:route GET /rule/{ID} rule getRule
		//
		// Get an rule
		//
		// This will get an individual rule available rules by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: RuleResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetRule",
			"GET",
			"/rule/{ID:[0-9]+}",
			getRule,
		},
		// swagger:route PUT /rule/{ID} rule editRule
		//
		// Edit an rule
		//
		// This will edit an rule
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: RuleResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditRule",
			"PUT",
			"/rule/{ID:[0-9]+}",
			editRule,
		},
		// swagger:route DELETE /rule/{ID} rule deleteRule
		//
		// Delete an rule
		//
		// This will delete an rule
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
			"DeleteRule",
			"DELETE",
			"/rule/{ID:[0-9]+}",
			deleteRule,
		},
	}
	return
}

func getRule(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleRequest{
		ID: getIntVar(r, "ID"),
	}

	rule := &model.Rule{
		ID: request.ID,
	}

	err = cases.GetRule(rule, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &RuleResponse{
		Rule: rule,
	}
	content = response
	return
}

func createRule(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	rule := &model.Rule{}
	err = decodeBody(r, rule)
	if err != nil {
		return
	}
	err = cases.CreateRule(rule, user)
	if err != nil {
		return
	}
	response := &RuleResponse{
		Rule: rule,
	}
	content = response
	return
}

func deleteRule(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleRequest{
		ID: getIntVar(r, "ID"),
	}

	rule := &model.Rule{
		ID: request.ID,
	}

	err = cases.DeleteRule(rule, user)
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

func editRule(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RuleEditRequest{
		ID: getIntVar(r, "ID"),
	}

	rule := &model.Rule{
		ID: request.ID,
	}

	err = decodeBody(r, rule)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditRule(rule, user)
	if err != nil {
		return
	}
	response := &RuleResponse{
		Rule: rule,
	}
	content = response
	return
}

func listRule(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	rules, err := cases.ListRule(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &RulesResponse{
		Page:  page,
		Rules: rules,
	}
	content = response
	return
}

func listRuleBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	rule := &model.Rule{}
	rule.Name = getQuery(r, "name")
	rules, err := cases.ListRuleBySearch(page, rule, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &RulesBySearchResponse{
		Page:   page,
		Rules:  rules,
		Search: rule,
	}
	content = response
	return
}
