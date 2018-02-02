package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// ZoneExpansionRequest is a list of parameters used for zoneExpansion
// swagger:parameters deleteZoneExpansion editZoneExpansion getZoneExpansion
type ZoneExpansionRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
}

// ZoneExpansionResponse is what endpoints respond with
// swagger:response
type ZoneExpansionResponse struct {
	ZoneExpansion *model.ZoneExpansion `json:"zoneExpansion,omitempty"`
}

// ZoneExpansionCreateRequest is the body parameters for creating an zoneExpansion
// swagger:parameters createZoneExpansion
type ZoneExpansionCreateRequest struct {
	// ZoneExpansion details to create
	// in: body
	ZoneExpansion *model.ZoneExpansion `json:"zoneExpansion"`
}

// ZoneExpansionEditRequest is the body parameters for creating an zoneExpansion
// swagger:parameters editZoneExpansion
type ZoneExpansionEditRequest struct {
	// ID to get information about
	// in: path
	// example: 1
	ID int64 `json:"ID"`
	// ZoneExpansion details to edit
	// in: body
	ZoneExpansion *model.ZoneExpansion `json:"zoneExpansion"`
}

// ZoneExpansionsRequest is a list of parameters used for zoneExpansion
// swagger:parameters listZoneExpansion
type ZoneExpansionsRequest struct {
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

// ZoneExpansionsResponse is a general response to a request
// swagger:response
type ZoneExpansionsResponse struct {
	Page           *model.Page          `json:"page,omitempty"`
	ZoneExpansions model.ZoneExpansions `json:"zoneExpansions,omitempty"`
}

// ZoneExpansionsBySearchRequest is a list of parameters used for zoneExpansion
// swagger:parameters listZoneExpansionBySearch
type ZoneExpansionsBySearchRequest struct {
	// ShortName is which zoneExpansion to get information about
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

// ZoneExpansionsBySearchResponse is a general response to a request
// swagger:response
type ZoneExpansionsBySearchResponse struct {
	Search         *model.ZoneExpansion `json:"search,omitempty"`
	Page           *model.Page          `json:"page,omitempty"`
	ZoneExpansions model.ZoneExpansions `json:"zoneExpansions,omitempty"`
}

func zoneExpansionRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /zone/expansion zone listZoneExpansion
		//
		// Lists zoneExpansions
		//
		// This will show all available zoneExpansions by default.
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
		//       200: ZoneExpansionsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListZoneExpansion",
			"GET",
			"/zone/expansion",
			listZoneExpansion,
		},
		// swagger:route GET /zone/expansion/search zone listZoneExpansionBySearch
		//
		// Search zoneExpansions by name
		//
		// This will show all available zoneExpansions by default.
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
		//       200: ZoneExpansionsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListZoneExpansionBySearch",
			"GET",
			"/zone/expansion/search",
			listZoneExpansionBySearch,
		},
		// swagger:route POST /zone/expansion zone createZoneExpansion
		//
		// Create an zoneExpansion
		//
		// This will create an zoneExpansion
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ZoneExpansionResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateZoneExpansion",
			"POST",
			"/zone/expansion",
			createZoneExpansion,
		},
		// swagger:route GET /zone/expansion/{ID} zone getZoneExpansion
		//
		// Get an zoneExpansion
		//
		// This will get an individual zoneExpansion available zoneExpansions by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: ZoneExpansionResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetZoneExpansion",
			"GET",
			"/zone/expansion/{ID:[0-9]+}",
			getZoneExpansion,
		},
		// swagger:route PUT /zone/expansion/{ID} zone editZoneExpansion
		//
		// Edit an zoneExpansion
		//
		// This will edit an zoneExpansion
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ZoneExpansionResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditZoneExpansion",
			"PUT",
			"/zone/expansion/{ID:[0-9]+}",
			editZoneExpansion,
		},
		// swagger:route DELETE /zone/expansion/{ID} zone deleteZoneExpansion
		//
		// Delete an zoneExpansion
		//
		// This will delete an zoneExpansion
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
			"DeleteZoneExpansion",
			"DELETE",
			"/zone/expansion/{ID:[0-9]+}",
			deleteZoneExpansion,
		},
	}
	return
}

func getZoneExpansion(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneExpansionRequest{
		ID: getIntVar(r, "ID"),
	}

	zoneExpansion := &model.ZoneExpansion{
		ID: request.ID,
	}

	err = cases.GetZoneExpansion(zoneExpansion, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ZoneExpansionResponse{
		ZoneExpansion: zoneExpansion,
	}
	content = response
	return
}

func createZoneExpansion(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	zoneExpansion := &model.ZoneExpansion{}
	err = decodeBody(r, zoneExpansion)
	if err != nil {
		return
	}
	err = cases.CreateZoneExpansion(zoneExpansion, user)
	if err != nil {
		return
	}
	response := &ZoneExpansionResponse{
		ZoneExpansion: zoneExpansion,
	}
	content = response
	return
}

func deleteZoneExpansion(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneExpansionRequest{
		ID: getIntVar(r, "ID"),
	}

	zoneExpansion := &model.ZoneExpansion{
		ID: request.ID,
	}

	err = cases.DeleteZoneExpansion(zoneExpansion, user)
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

func editZoneExpansion(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneExpansionEditRequest{
		ID: getIntVar(r, "ID"),
	}

	zoneExpansion := &model.ZoneExpansion{
		ID: request.ID,
	}

	err = decodeBody(r, zoneExpansion)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditZoneExpansion(zoneExpansion, user)
	if err != nil {
		return
	}
	response := &ZoneExpansionResponse{
		ZoneExpansion: zoneExpansion,
	}
	content = response
	return
}

func listZoneExpansion(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	zoneExpansions, err := cases.ListZoneExpansion(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ZoneExpansionsResponse{
		Page:           page,
		ZoneExpansions: zoneExpansions,
	}
	content = response
	return
}

func listZoneExpansionBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	zoneExpansion := &model.ZoneExpansion{}
	zoneExpansion.ShortName = getQuery(r, "shortName")
	zoneExpansions, err := cases.ListZoneExpansionBySearch(page, zoneExpansion, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	log.Println(zoneExpansions)
	response := &ZoneExpansionsBySearchResponse{
		Page:           page,
		ZoneExpansions: zoneExpansions,
		Search:         zoneExpansion,
	}
	content = response
	return
}
