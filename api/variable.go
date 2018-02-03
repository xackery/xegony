package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// VariableRequest is a list of parameters used for variable
// swagger:parameters deleteVariable editVariable getVariable
type VariableRequest struct {
	// Name to get information about
	// in: path
	// example: 74887
	Name string `json:"name"`
}

// VariableResponse is what endpoints respond with
// swagger:response
type VariableResponse struct {
	Variable *model.Variable `json:"variable,omitempty"`
}

// VariableCreateRequest is the body parameters for creating an variable
// swagger:parameters createVariable
type VariableCreateRequest struct {
	// Variable details to create
	// in: body
	Variable *model.Variable `json:"variable"`
}

// VariableEditRequest is the body parameters for creating an variable
// swagger:parameters editVariable
type VariableEditRequest struct {
	// Name to get information about
	// in: path
	// example: 74887
	Name string `json:"name"`
	// Variable details to edit
	// in: body
	Variable *model.Variable `json:"variable"`
}

// VariablesRequest is a list of parameters used for variable
// swagger:parameters listVariable
type VariablesRequest struct {
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

// VariablesResponse is a general response to a request
// swagger:response
type VariablesResponse struct {
	Page      *model.Page     `json:"page,omitempty"`
	Variables model.Variables `json:"variables,omitempty"`
}

// VariablesBySearchRequest is a list of parameters used for variable
// swagger:parameters listVariableBySearch
type VariablesBySearchRequest struct {
	// ShortName is which variable to get information about
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

// VariablesBySearchResponse is a general response to a request
// swagger:response
type VariablesBySearchResponse struct {
	Search    *model.Variable `json:"search,omitempty"`
	Page      *model.Page     `json:"page,omitempty"`
	Variables model.Variables `json:"variables,omitempty"`
}

func variableRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /variable variable listVariable
		//
		// Lists variables
		//
		// This will show all available variables by default.
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
		//       200: VariablesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListVariable",
			"GET",
			"/variable",
			listVariable,
		},
		// swagger:route GET /variable/search variable listVariableBySearch
		//
		// Search variables by name
		//
		// This will show all available variables by default.
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
		//       200: VariablesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListVariableBySearch",
			"GET",
			"/variable/search",
			listVariableBySearch,
		},
		// swagger:route POST /variable variable createVariable
		//
		// Create an variable
		//
		// This will create an variable
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: VariableResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateVariable",
			"POST",
			"/variable",
			createVariable,
		},
		// swagger:route GET /variable/{name} variable getVariable
		//
		// Get an variable
		//
		// This will get an individual variable available variables by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: VariableResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetVariable",
			"GET",
			"/variable/{name:[a-zA-Z_]+}",
			getVariable,
		},
		// swagger:route PUT /variable/{name} variable editVariable
		//
		// Edit an variable
		//
		// This will edit an variable
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: VariableResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditVariable",
			"PUT",
			"/variable/{name:[a-zA-Z_]+}",
			editVariable,
		},
		// swagger:route DELETE /variable/{name} variable deleteVariable
		//
		// Delete an variable
		//
		// This will delete an variable
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
			"DeleteVariable",
			"DELETE",
			"/variable/{name:[0-9]+}",
			deleteVariable,
		},
	}
	return
}

func getVariable(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &VariableRequest{
		Name: getVar(r, "name"),
	}

	variable := &model.Variable{
		Name: request.Name,
	}

	err = cases.GetVariable(variable, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &VariableResponse{
		Variable: variable,
	}
	content = response
	return
}

func createVariable(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	variable := &model.Variable{}
	err = decodeBody(r, variable)
	if err != nil {
		return
	}
	err = cases.CreateVariable(variable, user)
	if err != nil {
		return
	}
	response := &VariableResponse{
		Variable: variable,
	}
	content = response
	return
}

func deleteVariable(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &VariableRequest{
		Name: getVar(r, "name"),
	}

	variable := &model.Variable{
		Name: request.Name,
	}

	err = cases.DeleteVariable(variable, user)
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

func editVariable(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &VariableEditRequest{
		Name: getVar(r, "name"),
	}

	variable := &model.Variable{
		Name: request.Name,
	}

	err = decodeBody(r, variable)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditVariable(variable, user)
	if err != nil {
		return
	}
	response := &VariableResponse{
		Variable: variable,
	}
	content = response
	return
}

func listVariable(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	variables, err := cases.ListVariable(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &VariablesResponse{
		Page:      page,
		Variables: variables,
	}
	content = response
	return
}

func listVariableBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	variable := &model.Variable{}
	variable.Name = getQuery(r, "name")
	variables, err := cases.ListVariableBySearch(page, variable, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	log.Println(variables)
	response := &VariablesBySearchResponse{
		Page:      page,
		Variables: variables,
		Search:    variable,
	}
	content = response
	return
}
