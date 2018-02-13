package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// SizeRequest is a list of parameters used for size
// swagger:parameters deleteSize editSize getSize
type SizeRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// SizeResponse is what endpoints respond with
// swagger:response
type SizeResponse struct {
	Size *model.Size `json:"size,omitempty"`
}

// SizeCreateRequest is the body parameters for creating an size
// swagger:parameters createSize
type SizeCreateRequest struct {
	// Size details to create
	// in: body
	Size *model.Size `json:"size"`
}

// SizeEditRequest is the body parameters for creating an size
// swagger:parameters editSize
type SizeEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// Size details to edit
	// in: body
	Size *model.Size `json:"size"`
}

// SizesRequest is a list of parameters used for size
// swagger:parameters listSize
type SizesRequest struct {
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

// SizesResponse is a general response to a request
// swagger:response
type SizesResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Sizes model.Sizes `json:"sizes,omitempty"`
}

// SizesBySearchRequest is a list of parameters used for size
// swagger:parameters listSizeBySearch
type SizesBySearchRequest struct {
	// Name is which size to get information about
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

// SizesBySearchResponse is a general response to a request
// swagger:response
type SizesBySearchResponse struct {
	Search *model.Size `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Sizes  model.Sizes `json:"sizes,omitempty"`
}

func sizeRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /size size listSize
		//
		// Lists sizes
		//
		// This will show all available sizes by default.
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
		//       200: SizesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSize",
			"GET",
			"/size",
			listSize,
		},
		// swagger:route GET /size/search size listSizeBySearch
		//
		// Search sizes by name
		//
		// This will show all available sizes by default.
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
		//       200: SizesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSizeBySearch",
			"GET",
			"/size/search",
			listSizeBySearch,
		},
		// swagger:route POST /size size createSize
		//
		// Create an size
		//
		// This will create an size
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SizeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateSize",
			"POST",
			"/size",
			createSize,
		},
		// swagger:route GET /size/{ID} size getSize
		//
		// Get an size
		//
		// This will get an individual size available sizes by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SizeResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSize",
			"GET",
			"/size/{ID:[0-9]+}",
			getSize,
		},
		// swagger:route PUT /size/{ID} size editSize
		//
		// Edit an size
		//
		// This will edit an size
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: SizeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditSize",
			"PUT",
			"/size/{ID:[0-9]+}",
			editSize,
		},
		// swagger:route DELETE /size/{ID} size deleteSize
		//
		// Delete an size
		//
		// This will delete an size
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
			"DeleteSize",
			"DELETE",
			"/size/{ID:[0-9]+}",
			deleteSize,
		},
	}
	return
}

func getSize(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SizeRequest{
		ID: getIntVar(r, "ID"),
	}

	size := &model.Size{
		ID: request.ID,
	}

	err = cases.GetSize(size, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &SizeResponse{
		Size: size,
	}
	content = response
	return
}

func createSize(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	size := &model.Size{}
	err = decodeBody(r, size)
	if err != nil {
		return
	}
	err = cases.CreateSize(size, user)
	if err != nil {
		return
	}
	response := &SizeResponse{
		Size: size,
	}
	content = response
	return
}

func deleteSize(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SizeRequest{
		ID: getIntVar(r, "ID"),
	}

	size := &model.Size{
		ID: request.ID,
	}

	err = cases.DeleteSize(size, user)
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

func editSize(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &SizeEditRequest{
		ID: getIntVar(r, "ID"),
	}

	size := &model.Size{
		ID: request.ID,
	}

	err = decodeBody(r, size)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditSize(size, user)
	if err != nil {
		return
	}
	response := &SizeResponse{
		Size: size,
	}
	content = response
	return
}

func listSize(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	sizes, err := cases.ListSize(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SizesResponse{
		Page:  page,
		Sizes: sizes,
	}
	content = response
	return
}

func listSizeBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	size := &model.Size{
		Name: getQuery(r, "name"),
	}
	sizes, err := cases.ListSizeBySearch(page, size, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &SizesBySearchResponse{
		Page:   page,
		Sizes:  sizes,
		Search: size,
	}
	content = response
	return
}
