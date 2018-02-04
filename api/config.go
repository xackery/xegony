package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// ConfigRequest is a list of parameters used for config
// swagger:parameters deleteConfig editConfig getConfig
type ConfigRequest struct {
	// Key to get information about
	// in: path
	// example: 74887
	Key string `json:"key"`
}

// ConfigResponse is what endpoints respond with
// swagger:response
type ConfigResponse struct {
	Config *model.Config `json:"config,omitempty"`
}

// ConfigCreateRequest is the body parameters for creating an config
// swagger:parameters createConfig
type ConfigCreateRequest struct {
	// Config details to create
	// in: body
	Config *model.Config `json:"config"`
}

// ConfigEditRequest is the body parameters for creating an config
// swagger:parameters editConfig
type ConfigEditRequest struct {
	// Key to get information about
	// in: path
	// example: googleToken
	Key string `json:"key"`
	// Config details to edit
	// in: body
	Config *model.Config `json:"config"`
}

// ConfigsRequest is a list of parameters used for config
// swagger:parameters listConfig
type ConfigsRequest struct {
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

// ConfigsResponse is a general response to a request
// swagger:response
type ConfigsResponse struct {
	Page    *model.Page   `json:"page,omitempty"`
	Configs model.Configs `json:"configs,omitempty"`
}

// ConfigsBySearchRequest is a list of parameters used for config
// swagger:parameters listConfigBySearch
type ConfigsBySearchRequest struct {
	// Key is which config to get information about
	// example: googleToken
	// in: query
	Key string `json:"key"`
	// Offset is pagination, offset*limit
	// example: 0
	// in: query
	Offset int64 `json:"offset"`
	// Limit to how many items per page
	// example: 10
	// in: query
	Limit int64 `json:"limit"`
	// OrderBy is which field to order a page by
	// example: key
	// in: query
	OrderBy string `json:"orderBy"`
	// IsDescending will change sort order when true
	// example: 0
	// in: query
	IsDescending int64 `json:"isDescending"`
}

// ConfigsBySearchResponse is a general response to a request
// swagger:response
type ConfigsBySearchResponse struct {
	Search  *model.Config `json:"search,omitempty"`
	Page    *model.Page   `json:"page,omitempty"`
	Configs model.Configs `json:"configs,omitempty"`
}

func configRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /config config listConfig
		//
		// Lists configs
		//
		// This will show all available configs by default.
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
		//       200: ConfigsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListConfig",
			"GET",
			"/config",
			listConfig,
		},
		// swagger:route GET /config/search config listConfigBySearch
		//
		// Search configs by name
		//
		// This will show all available configs by default.
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
		//       200: ConfigsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListConfigBySearch",
			"GET",
			"/config/search",
			listConfigBySearch,
		},
		// swagger:route POST /config config createConfig
		//
		// Create an config
		//
		// This will create an config
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ConfigResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateConfig",
			"POST",
			"/config",
			createConfig,
		},
		// swagger:route GET /config/{key} config getConfig
		//
		// Get an config
		//
		// This will get an individual config available configs by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: ConfigResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetConfig",
			"GET",
			"/config/{key:[a-zA-Z]+}",
			getConfig,
		},
		// swagger:route PUT /config/{key} config editConfig
		//
		// Edit an config
		//
		// This will edit an config
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ConfigResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditConfig",
			"PUT",
			"/config/{key:[a-zA-Z]+}",
			editConfig,
		},
		// swagger:route DELETE /config/{key} config deleteConfig
		//
		// Delete an config
		//
		// This will delete an config
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
			"DeleteConfig",
			"DELETE",
			"/config/{key:[a-zA-Z]+}",
			deleteConfig,
		},
	}
	return
}

func getConfig(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ConfigRequest{
		Key: getVar(r, "key"),
	}

	config := &model.Config{
		Key: request.Key,
	}

	err = cases.GetConfig(config, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ConfigResponse{
		Config: config,
	}
	content = response
	return
}

func createConfig(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	config := &model.Config{}
	err = decodeBody(r, config)
	if err != nil {
		return
	}
	err = cases.CreateConfig(config, user)
	if err != nil {
		return
	}
	response := &ConfigResponse{
		Config: config,
	}
	content = response
	return
}

func deleteConfig(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ConfigRequest{
		Key: getVar(r, "Key"),
	}

	config := &model.Config{
		Key: request.Key,
	}

	err = cases.DeleteConfig(config, user)
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

func editConfig(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ConfigEditRequest{
		Key: getVar(r, "Key"),
	}

	config := &model.Config{
		Key: request.Key,
	}

	err = decodeBody(r, config)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditConfig(config, user)
	if err != nil {
		return
	}
	response := &ConfigResponse{
		Config: config,
	}
	content = response
	return
}

func listConfig(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	configs, err := cases.ListConfig(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ConfigsResponse{
		Page:    page,
		Configs: configs,
	}
	content = response
	return
}

func listConfigBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	config := &model.Config{
		Key: getQuery(r, "key"),
	}
	configs, err := cases.ListConfigBySearch(page, config, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ConfigsBySearchResponse{
		Page:    page,
		Configs: configs,
		Search:  config,
	}
	content = response
	return
}
