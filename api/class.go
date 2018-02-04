package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// ClassRequest is a list of parameters used for class
// swagger:parameters deleteClass editClass getClass
type ClassRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// ClassResponse is what endpoints respond with
// swagger:response
type ClassResponse struct {
	Class *model.Class `json:"class,omitempty"`
}

// ClassCreateRequest is the body parameters for creating an class
// swagger:parameters createClass
type ClassCreateRequest struct {
	// Class details to create
	// in: body
	Class *model.Class `json:"class"`
}

// ClassEditRequest is the body parameters for creating an class
// swagger:parameters editClass
type ClassEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// Class details to edit
	// in: body
	Class *model.Class `json:"class"`
}

// ClasssRequest is a list of parameters used for class
// swagger:parameters listClass
type ClasssRequest struct {
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

// ClasssResponse is a general response to a request
// swagger:response
type ClasssResponse struct {
	Page   *model.Page  `json:"page,omitempty"`
	Classs model.Classs `json:"classs,omitempty"`
}

// ClasssBySearchRequest is a list of parameters used for class
// swagger:parameters listClassBySearch
type ClasssBySearchRequest struct {
	// Name is which class to get information about
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
	// example: name
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// ClasssBySearchResponse is a general response to a request
// swagger:response
type ClasssBySearchResponse struct {
	Search *model.Class `json:"search,omitempty"`
	Page   *model.Page  `json:"page,omitempty"`
	Classs model.Classs `json:"classs,omitempty"`
}

func classRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /class class listClass
		//
		// Lists classs
		//
		// This will show all available classs by default.
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
		//       200: ClasssResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListClass",
			"GET",
			"/class",
			listClass,
		},
		// swagger:route GET /class/search class listClassBySearch
		//
		// Search classs by name
		//
		// This will show all available classs by default.
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
		//       200: ClasssBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListClassBySearch",
			"GET",
			"/class/search",
			listClassBySearch,
		},
		// swagger:route POST /class class createClass
		//
		// Create an class
		//
		// This will create an class
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ClassResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateClass",
			"POST",
			"/class",
			createClass,
		},
		// swagger:route GET /class/{ID} class getClass
		//
		// Get an class
		//
		// This will get an individual class available classs by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: ClassResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetClass",
			"GET",
			"/class/{ID:[0-9]+}",
			getClass,
		},
		// swagger:route PUT /class/{ID} class editClass
		//
		// Edit an class
		//
		// This will edit an class
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ClassResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditClass",
			"PUT",
			"/class/{ID:[0-9]+}",
			editClass,
		},
		// swagger:route DELETE /class/{ID} class deleteClass
		//
		// Delete an class
		//
		// This will delete an class
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
			"DeleteClass",
			"DELETE",
			"/class/{ID:[0-9]+}",
			deleteClass,
		},
	}
	return
}

func getClass(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ClassRequest{
		ID: getIntVar(r, "ID"),
	}

	class := &model.Class{
		ID: request.ID,
	}

	err = cases.GetClass(class, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ClassResponse{
		Class: class,
	}
	content = response
	return
}

func createClass(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	class := &model.Class{}
	err = decodeBody(r, class)
	if err != nil {
		return
	}
	err = cases.CreateClass(class, user)
	if err != nil {
		return
	}
	response := &ClassResponse{
		Class: class,
	}
	content = response
	return
}

func deleteClass(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ClassRequest{
		ID: getIntVar(r, "ID"),
	}

	class := &model.Class{
		ID: request.ID,
	}

	err = cases.DeleteClass(class, user)
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

func editClass(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ClassEditRequest{
		ID: getIntVar(r, "ID"),
	}

	class := &model.Class{
		ID: request.ID,
	}

	err = decodeBody(r, class)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditClass(class, user)
	if err != nil {
		return
	}
	response := &ClassResponse{
		Class: class,
	}
	content = response
	return
}

func listClass(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	classs, err := cases.ListClass(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ClasssResponse{
		Page:   page,
		Classs: classs,
	}
	content = response
	return
}

func listClassBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	class := &model.Class{
		Name: getQuery(r, "name"),
	}
	classs, err := cases.ListClassBySearch(page, class, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ClasssBySearchResponse{
		Page:   page,
		Classs: classs,
		Search: class,
	}
	content = response
	return
}
