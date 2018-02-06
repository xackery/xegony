package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// OauthTypeRequest is a list of parameters used for oauthType
// swagger:parameters deleteOauthType editOauthType getOauthType
type OauthTypeRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// OauthTypeResponse is what endpoints respond with
// swagger:response
type OauthTypeResponse struct {
	OauthType *model.OauthType `json:"oauthType,omitempty"`
}

// OauthTypeCreateRequest is the body parameters for creating an oauthType
// swagger:parameters createOauthType
type OauthTypeCreateRequest struct {
	// OauthType details to create
	// in: body
	OauthType *model.OauthType `json:"oauthType"`
}

// OauthTypeEditRequest is the body parameters for creating an oauthType
// swagger:parameters editOauthType
type OauthTypeEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// OauthType details to edit
	// in: body
	OauthType *model.OauthType `json:"oauthType"`
}

// OauthTypesRequest is a list of parameters used for oauthType
// swagger:parameters listOauthType
type OauthTypesRequest struct {
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

// OauthTypesResponse is a general response to a request
// swagger:response
type OauthTypesResponse struct {
	Page       *model.Page      `json:"page,omitempty"`
	OauthTypes model.OauthTypes `json:"oauthTypes,omitempty"`
}

// OauthTypesBySearchRequest is a list of parameters used for oauthType
// swagger:parameters listOauthTypeBySearch
type OauthTypesBySearchRequest struct {
	// Name is which oauthType to get information about
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

// OauthTypesBySearchResponse is a general response to a request
// swagger:response
type OauthTypesBySearchResponse struct {
	Search     *model.OauthType `json:"search,omitempty"`
	Page       *model.Page      `json:"page,omitempty"`
	OauthTypes model.OauthTypes `json:"oauthTypes,omitempty"`
}

func oauthTypeRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /oauth/type oauthType listOauthType
		//
		// Lists oauthTypes
		//
		// This will show all available oauthTypes by default.
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
		//       200: OauthTypesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListOauthType",
			"GET",
			"/oauth/type",
			listOauthType,
		},
		// swagger:route GET /oauth/type/search oauthType listOauthTypeBySearch
		//
		// Search oauthTypes by name
		//
		// This will show all available oauthTypes by default.
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
		//       200: OauthTypesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListOauthTypeBySearch",
			"GET",
			"/oauth/type/search",
			listOauthTypeBySearch,
		},
		// swagger:route POST /oauth/type oauthType createOauthType
		//
		// Create an oauthType
		//
		// This will create an oauthType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: OauthTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateOauthType",
			"POST",
			"/oauth/type",
			createOauthType,
		},
		// swagger:route GET /oauth/type/{ID} oauthType getOauthType
		//
		// Get an oauthType
		//
		// This will get an individual oauthType available oauthTypes by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: OauthTypeResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetOauthType",
			"GET",
			"/oauth/type/{ID:[0-9]+}",
			getOauthType,
		},
		// swagger:route PUT /oauth/type/{ID} oauthType editOauthType
		//
		// Edit an oauthType
		//
		// This will edit an oauthType
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: OauthTypeResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditOauthType",
			"PUT",
			"/oauth/type/{ID:[0-9]+}",
			editOauthType,
		},
		// swagger:route DELETE /oauth/type/{ID} oauthType deleteOauthType
		//
		// Delete an oauthType
		//
		// This will delete an oauthType
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
			"DeleteOauthType",
			"DELETE",
			"/oauth/type/{ID:[0-9]+}",
			deleteOauthType,
		},
	}
	return
}

func getOauthType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &OauthTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	oauthType := &model.OauthType{
		ID: request.ID,
	}

	err = cases.GetOauthType(oauthType, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &OauthTypeResponse{
		OauthType: oauthType,
	}
	content = response
	return
}

func createOauthType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	oauthType := &model.OauthType{}
	err = decodeBody(r, oauthType)
	if err != nil {
		return
	}
	err = cases.CreateOauthType(oauthType, user)
	if err != nil {
		return
	}
	response := &OauthTypeResponse{
		OauthType: oauthType,
	}
	content = response
	return
}

func deleteOauthType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &OauthTypeRequest{
		ID: getIntVar(r, "ID"),
	}

	oauthType := &model.OauthType{
		ID: request.ID,
	}

	err = cases.DeleteOauthType(oauthType, user)
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

func editOauthType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &OauthTypeEditRequest{
		ID: getIntVar(r, "ID"),
	}

	oauthType := &model.OauthType{
		ID: request.ID,
	}

	err = decodeBody(r, oauthType)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditOauthType(oauthType, user)
	if err != nil {
		return
	}
	response := &OauthTypeResponse{
		OauthType: oauthType,
	}
	content = response
	return
}

func listOauthType(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	oauthTypes, err := cases.ListOauthType(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &OauthTypesResponse{
		Page:       page,
		OauthTypes: oauthTypes,
	}
	content = response
	return
}

func listOauthTypeBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	oauthType := &model.OauthType{
		Name: getQuery(r, "name"),
	}
	oauthTypes, err := cases.ListOauthTypeBySearch(page, oauthType, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &OauthTypesBySearchResponse{
		Page:       page,
		OauthTypes: oauthTypes,
		Search:     oauthType,
	}
	content = response
	return
}
