package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// ForumRequest is a list of parameters used for forum
// swagger:parameters deleteForum editForum getForum
type ForumRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
}

// ForumResponse is what endpoints respond with
// swagger:response
type ForumResponse struct {
	Forum *model.Forum `json:"forum,omitempty"`
}

// ForumCreateRequest is the body parameters for creating an forum
// swagger:parameters createForum
type ForumCreateRequest struct {
	// Forum details to create
	// in: body
	Forum *model.Forum `json:"forum"`
}

// ForumEditRequest is the body parameters for creating an forum
// swagger:parameters editForum
type ForumEditRequest struct {
	// ID to get information about
	// in: path
	// example: 12
	ID int64 `json:"ID"`
	// Forum details to edit
	// in: body
	Forum *model.Forum `json:"forum"`
}

// ForumsRequest is a list of parameters used for forum
// swagger:parameters listForum
type ForumsRequest struct {
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

// ForumsResponse is a general response to a request
// swagger:response
type ForumsResponse struct {
	Page   *model.Page  `json:"page,omitempty"`
	Forums model.Forums `json:"forums,omitempty"`
}

// ForumsBySearchRequest is a list of parameters used for forum
// swagger:parameters listForumBySearch
type ForumsBySearchRequest struct {
	// Name is which forum to get information about
	// example: heal
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

// ForumsBySearchResponse is a general response to a request
// swagger:response
type ForumsBySearchResponse struct {
	Search *model.Forum `json:"search,omitempty"`
	Page   *model.Page  `json:"page,omitempty"`
	Forums model.Forums `json:"forums,omitempty"`
}

func forumRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /forum forum listForum
		//
		// Lists forums
		//
		// This will show all available forums by default.
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
		//       200: ForumsResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListForum",
			"GET",
			"/forum",
			listForum,
		},
		// swagger:route GET /forum/search forum listForumBySearch
		//
		// Search forums by name
		//
		// This will show all available forums by default.
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
		//       200: ForumsBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListForumBySearch",
			"GET",
			"/forum/search",
			listForumBySearch,
		},
		// swagger:route POST /forum forum createForum
		//
		// Create an forum
		//
		// This will create an forum
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ForumResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateForum",
			"POST",
			"/forum",
			createForum,
		},
		// swagger:route GET /forum/{ID} forum getForum
		//
		// Get an forum
		//
		// This will get an individual forum available forums by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: ForumResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetForum",
			"GET",
			"/forum/{ID:[0-9]+}",
			getForum,
		},
		// swagger:route PUT /forum/{ID} forum editForum
		//
		// Edit an forum
		//
		// This will edit an forum
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: ForumResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditForum",
			"PUT",
			"/forum/{ID:[0-9]+}",
			editForum,
		},
		// swagger:route DELETE /forum/{ID} forum deleteForum
		//
		// Delete an forum
		//
		// This will delete an forum
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
			"DeleteForum",
			"DELETE",
			"/forum/{ID:[0-9]+}",
			deleteForum,
		},
	}
	return
}

func getForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ForumRequest{
		ID: getIntVar(r, "ID"),
	}

	forum := &model.Forum{
		ID: request.ID,
	}

	err = cases.GetForum(forum, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &ForumResponse{
		Forum: forum,
	}
	content = response
	return
}

func createForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	forum := &model.Forum{}
	err = decodeBody(r, forum)
	if err != nil {
		return
	}
	err = cases.CreateForum(forum, user)
	if err != nil {
		return
	}
	response := &ForumResponse{
		Forum: forum,
	}
	content = response
	return
}

func deleteForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ForumRequest{
		ID: getIntVar(r, "ID"),
	}

	forum := &model.Forum{
		ID: request.ID,
	}

	err = cases.DeleteForum(forum, user)
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

func editForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &ForumEditRequest{
		ID: getIntVar(r, "ID"),
	}

	forum := &model.Forum{
		ID: request.ID,
	}

	err = decodeBody(r, forum)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditForum(forum, user)
	if err != nil {
		return
	}
	response := &ForumResponse{
		Forum: forum,
	}
	content = response
	return
}

func listForum(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	forums, err := cases.ListForum(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ForumsResponse{
		Page:   page,
		Forums: forums,
	}
	content = response
	return
}

func listForumBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	forum := &model.Forum{}
	forum.Name = getQuery(r, "name")
	forums, err := cases.ListForumBySearch(page, forum, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &ForumsBySearchResponse{
		Page:   page,
		Forums: forums,
		Search: forum,
	}
	content = response
	return
}
