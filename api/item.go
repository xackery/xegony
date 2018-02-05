package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// ItemRequest is a list of parameters used for item
// swagger:parameters deleteItem editItem getItem
type ItemRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// ItemResponse is what endpoints respond with
// swagger:response
type ItemResponse struct {
	Item *model.Item `json:"item,omitempty"`
}

// ItemCreateRequest is the body parameters for creating an item
// swagger:parameters createItem
type ItemCreateRequest struct {
	// Item details to create
	// in: body
	Item *model.Item `json:"item"`
}

// ItemEditRequest is the body parameters for creating an item
// swagger:parameters editItem
type ItemEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Item details to edit
	// in: body
	Item *model.Item `json:"item"`
}

// ItemsRequest is a list of parameters used for item
// swagger:parameters listItem
type ItemsRequest struct {
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

// ItemsResponse is a general response to a request
// swagger:response
type ItemsResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Items model.Items `json:"items,omitempty"`
}

// ItemsBySearchRequest is a list of parameters used for item
// swagger:parameters listItemBySearch
type ItemsBySearchRequest struct {
	// Name is which item to get information about
	// example: singing short sword
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

// ItemsBySearchResponse is a general response to a request
// swagger:response
type ItemsBySearchResponse struct {
	Search *model.Item `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Items  model.Items `json:"items,omitempty"`
}

func itemRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /item item listItem
		//
		// Lists items
		//
		// This will show all available items by default.
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
		//       200: ItemsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListItem",
			"GET",
			"/item",
			listItem,
		},
		// swagger:route GET /item/search item listItemBySearch
		//
		// Search items by name
		//
		// This will show all available items by default.
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
		//       200: ItemsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListItemBySearch",
			"GET",
			"/item/search",
			listItemBySearch,
		},
		// swagger:route POST /item item createItem
		//
		// Create an item
		//
		// This will create an item
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ItemResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateItem",
			"POST",
			"/item",
			createItem,
		},
		// swagger:route GET /item/{ID} item getItem
		//
		// Get an item
		//
		// This will get an individual item available items by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: ItemResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetItem",
			"GET",
			"/item/{ID:[0-9]+}",
			getItem,
		},
		// swagger:route PUT /item/{ID} item editItem
		//
		// Edit an item
		//
		// This will edit an item
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ItemResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditItem",
			"PUT",
			"/item/{ID:[0-9]+}",
			editItem,
		},
		// swagger:route DELETE /item/{ID} item deleteItem
		//
		// Delete an item
		//
		// This will delete an item
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
			"DeleteItem",
			"DELETE",
			"/item/{ID:[0-9]+}",
			deleteItem,
		},
	}
	return
}

func getItem(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ItemRequest{
		ID: getIntVar(r, "ID"),
	}

	item := &model.Item{
		ID: request.ID,
	}

	err = cases.GetItem(item, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ItemResponse{
		Item: item,
	}
	content = response
	return
}

func createItem(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	item := &model.Item{}
	err = decodeBody(r, item)
	if err != nil {
		return
	}
	err = cases.CreateItem(item, user)
	if err != nil {
		return
	}
	response := &ItemResponse{
		Item: item,
	}
	content = response
	return
}

func deleteItem(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ItemRequest{
		ID: getIntVar(r, "ID"),
	}

	item := &model.Item{
		ID: request.ID,
	}

	err = cases.DeleteItem(item, user)
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

func editItem(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ItemEditRequest{
		ID: getIntVar(r, "ID"),
	}

	item := &model.Item{
		ID: request.ID,
	}

	err = decodeBody(r, item)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditItem(item, user)
	if err != nil {
		return
	}
	response := &ItemResponse{
		Item: item,
	}
	content = response
	return
}

func listItem(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	items, err := cases.ListItem(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ItemsResponse{
		Page:  page,
		Items: items,
	}
	content = response
	return
}

func listItemBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	item := &model.Item{
		Name: getQuery(r, "name"),
	}
	items, err := cases.ListItemBySearch(page, item, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ItemsBySearchResponse{
		Page:   page,
		Items:  items,
		Search: item,
	}
	content = response
	return
}
