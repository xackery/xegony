package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// ZoneRequest is a list of parameters used for zone
// swagger:parameters deleteZone editZone getZone
type ZoneRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// ZoneResponse is what endpoints respond with
// swagger:response
type ZoneResponse struct {
	Zone *model.Zone `json:"zone,omitempty"`
}

// ZoneCreateRequest is the body parameters for creating an zone
// swagger:parameters createZone
type ZoneCreateRequest struct {
	// Zone details to create
	// in: body
	Zone *model.Zone `json:"zone"`
}

// ZoneEditRequest is the body parameters for creating an zone
// swagger:parameters editZone
type ZoneEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Zone details to edit
	// in: body
	Zone *model.Zone `json:"zone"`
}

// ZonesRequest is a list of parameters used for zone
// swagger:parameters listZone
type ZonesRequest struct {
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

// ZonesResponse is a general response to a request
// swagger:response
type ZonesResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Zones model.Zones `json:"zones,omitempty"`
}

// ZonesBySearchRequest is a list of parameters used for zone
// swagger:parameters listZoneBySearch
type ZonesBySearchRequest struct {
	// ShortName is which zone to get information about
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

// ZonesBySearchResponse is a general response to a request
// swagger:response
type ZonesBySearchResponse struct {
	Search *model.Zone `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Zones  model.Zones `json:"zones,omitempty"`
}

func zoneRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /zone zone listZone
		//
		// Lists zones
		//
		// This will show all available zones by default.
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
		//       200: ZonesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListZone",
			"GET",
			"/zone",
			listZone,
		},
		// swagger:route GET /zone/search zone listZoneBySearch
		//
		// Search zones by name
		//
		// This will show all available zones by default.
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
		//       200: ZonesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListZoneBySearch",
			"GET",
			"/zone/search",
			listZoneBySearch,
		},
		// swagger:route POST /zone zone createZone
		//
		// Create an zone
		//
		// This will create an zone
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ZoneResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateZone",
			"POST",
			"/zone",
			createZone,
		},
		// swagger:route GET /zone/{ID} zone getZone
		//
		// Get an zone
		//
		// This will get an individual zone available zones by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: ZoneResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetZone",
			"GET",
			"/zone/{ID:[0-9]+}",
			getZone,
		},
		// swagger:route PUT /zone/{ID} zone editZone
		//
		// Edit an zone
		//
		// This will edit an zone
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ZoneResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditZone",
			"PUT",
			"/zone/{ID:[0-9]+}",
			editZone,
		},
		// swagger:route DELETE /zone/{ID} zone deleteZone
		//
		// Delete an zone
		//
		// This will delete an zone
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
			"DeleteZone",
			"DELETE",
			"/zone/{ID:[0-9]+}",
			deleteZone,
		},
	}
	return
}

func getZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneRequest{
		ID: getIntVar(r, "ID"),
	}

	zone := &model.Zone{
		ID: request.ID,
	}

	err = cases.GetZone(zone, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ZoneResponse{
		Zone: zone,
	}
	content = response
	return
}

func createZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	zone := &model.Zone{}
	err = decodeBody(r, zone)
	if err != nil {
		return
	}
	err = cases.CreateZone(zone, user)
	if err != nil {
		return
	}
	response := &ZoneResponse{
		Zone: zone,
	}
	content = response
	return
}

func deleteZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneRequest{
		ID: getIntVar(r, "ID"),
	}

	zone := &model.Zone{
		ID: request.ID,
	}

	err = cases.DeleteZone(zone, user)
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

func editZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ZoneEditRequest{
		ID: getIntVar(r, "ID"),
	}

	zone := &model.Zone{
		ID: request.ID,
	}

	err = decodeBody(r, zone)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditZone(zone, user)
	if err != nil {
		return
	}
	response := &ZoneResponse{
		Zone: zone,
	}
	content = response
	return
}

func listZone(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	zones, err := cases.ListZone(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ZonesResponse{
		Page:  page,
		Zones: zones,
	}
	content = response
	return
}

func listZoneBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	zone := &model.Zone{}
	zone.ShortName.String = getQuery(r, "shortName")
	zones, err := cases.ListZoneBySearch(page, zone, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	log.Println(zones)
	response := &ZonesBySearchResponse{
		Page:   page,
		Zones:  zones,
		Search: zone,
	}
	content = response
	return
}
