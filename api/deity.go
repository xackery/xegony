package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// DeityRequest is a list of parameters used for deity
// swagger:parameters deleteDeity editDeity getDeity
type DeityRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// DeityResponse is what endpoints respond with
// swagger:response
type DeityResponse struct {
	Deity *model.Deity `json:"deity,omitempty"`
}

// DeityCreateRequest is the body parameters for creating an deity
// swagger:parameters createDeity
type DeityCreateRequest struct {
	// Deity details to create
	// in: body
	Deity *model.Deity `json:"deity"`
}

// DeityEditRequest is the body parameters for creating an deity
// swagger:parameters editDeity
type DeityEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// Deity details to edit
	// in: body
	Deity *model.Deity `json:"deity"`
}

// DeitysRequest is a list of parameters used for deity
// swagger:parameters listDeity
type DeitysRequest struct {
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

// DeitysResponse is a general response to a request
// swagger:response
type DeitysResponse struct {
	Page   *model.Page  `json:"page,omitempty"`
	Deitys model.Deitys `json:"deitys,omitempty"`
}

// DeitysBySearchRequest is a list of parameters used for deity
// swagger:parameters listDeityBySearch
type DeitysBySearchRequest struct {
	// Name is which deity to get information about
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

// DeitysBySearchResponse is a general response to a request
// swagger:response
type DeitysBySearchResponse struct {
	Search *model.Deity `json:"search,omitempty"`
	Page   *model.Page  `json:"page,omitempty"`
	Deitys model.Deitys `json:"deitys,omitempty"`
}

func deityRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /deity deity listDeity
		//
		// Lists deitys
		//
		// This will show all available deitys by default.
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
		//       200: DeitysResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListDeity",
			"GET",
			"/deity",
			listDeity,
		},
		// swagger:route GET /deity/search deity listDeityBySearch
		//
		// Search deitys by name
		//
		// This will show all available deitys by default.
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
		//       200: DeitysBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListDeityBySearch",
			"GET",
			"/deity/search",
			listDeityBySearch,
		},
		// swagger:route POST /deity deity createDeity
		//
		// Create an deity
		//
		// This will create an deity
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: DeityResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateDeity",
			"POST",
			"/deity",
			createDeity,
		},
		// swagger:route GET /deity/{ID} deity getDeity
		//
		// Get an deity
		//
		// This will get an individual deity available deitys by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: DeityResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetDeity",
			"GET",
			"/deity/{ID:[0-9]+}",
			getDeity,
		},
		// swagger:route PUT /deity/{ID} deity editDeity
		//
		// Edit an deity
		//
		// This will edit an deity
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: DeityResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditDeity",
			"PUT",
			"/deity/{ID:[0-9]+}",
			editDeity,
		},
		// swagger:route DELETE /deity/{ID} deity deleteDeity
		//
		// Delete an deity
		//
		// This will delete an deity
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
			"DeleteDeity",
			"DELETE",
			"/deity/{ID:[0-9]+}",
			deleteDeity,
		},
	}
	return
}

func getDeity(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &DeityRequest{
		ID: getIntVar(r, "ID"),
	}

	deity := &model.Deity{
		ID: request.ID,
	}

	err = cases.GetDeity(deity, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &DeityResponse{
		Deity: deity,
	}
	content = response
	return
}

func createDeity(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	deity := &model.Deity{}
	err = decodeBody(r, deity)
	if err != nil {
		return
	}
	err = cases.CreateDeity(deity, user)
	if err != nil {
		return
	}
	response := &DeityResponse{
		Deity: deity,
	}
	content = response
	return
}

func deleteDeity(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &DeityRequest{
		ID: getIntVar(r, "ID"),
	}

	deity := &model.Deity{
		ID: request.ID,
	}

	err = cases.DeleteDeity(deity, user)
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

func editDeity(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &DeityEditRequest{
		ID: getIntVar(r, "ID"),
	}

	deity := &model.Deity{
		ID: request.ID,
	}

	err = decodeBody(r, deity)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditDeity(deity, user)
	if err != nil {
		return
	}
	response := &DeityResponse{
		Deity: deity,
	}
	content = response
	return
}

func listDeity(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	deitys, err := cases.ListDeity(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &DeitysResponse{
		Page:   page,
		Deitys: deitys,
	}
	content = response
	return
}

func listDeityBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	deity := &model.Deity{
		Name: getQuery(r, "name"),
	}
	deitys, err := cases.ListDeityBySearch(page, deity, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &DeitysBySearchResponse{
		Page:   page,
		Deitys: deitys,
		Search: deity,
	}
	content = response
	return
}
