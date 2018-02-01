package api

import (
	"database/sql"
	"encoding/base64"
	"math/rand"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// UserRequest is a list of parameters used for user
// swagger:parameters deleteUser editUser getUser
type UserRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// UserResponse is what endpoints respond with
// swagger:response
type UserResponse struct {
	User *model.User `json:"user,omitempty"`
}

// UserGoogleStartRequest is a list of parameters used for initiating an oauth process
// swagger:parameters getUserGoogleStart
type UserGoogleStartRequest struct {
	// ReturnURL is the URL to return once Oauth completes
	// in: query
	// example: http://everzek.com
	ReturnURL string `json:"returnURL"`
}

// UserResponse is what endpoints respond with
// swagger:response
type UserGoogleStartResponse struct {
	RedirectURL string `json:"redirectURL"`
}

// UserCreateRequest is the body parameters for creating an user
// swagger:parameters createUser
type UserCreateRequest struct {
	// User details to create
	// in: body
	User *model.User `json:"user"`
}

// UserEditRequest is the body parameters for creating an user
// swagger:parameters editUser
type UserEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// User details to edit
	// in: body
	User *model.User `json:"user"`
}

// UsersRequest is a list of parameters used for user
// swagger:parameters listUser
type UsersRequest struct {
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

// UsersResponse is a general response to a request
// swagger:response
type UsersResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Users model.Users `json:"users,omitempty"`
}

// UsersBySearchRequest is a list of parameters used for user
// swagger:parameters listUserBySearch
type UsersBySearchRequest struct {
	// ShortName is which user to get information about
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

// UsersBySearchResponse is a general response to a request
// swagger:response
type UsersBySearchResponse struct {
	Search *model.User `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Users  model.Users `json:"users,omitempty"`
}

func userRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /user user listUser
		//
		// Lists users
		//
		// This will show all available users by default.
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
		//       200: UsersResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListUser",
			"GET",
			"/user",
			listUser,
		},
		// swagger:route GET /user/search user listUserBySearch
		//
		// Search users by name
		//
		// This will show all available users by default.
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
		//       200: UsersBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListUserBySearch",
			"GET",
			"/user/search",
			listUserBySearch,
		},
		// swagger:route POST /user user createUser
		//
		// Create an user
		//
		// This will create an user
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: UserResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateUser",
			"POST",
			"/user",
			createUser,
		},
		// swagger:route GET /user/{ID} user getUser
		//
		// Get an user
		//
		// This will get an individual user available users by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: UserResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetUser",
			"GET",
			"/user/{ID:[0-9]+}",
			getUser,
		},
		// swagger:route PUT /user/{ID} user editUser
		//
		// Edit an user
		//
		// This will edit an user
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: UserResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditUser",
			"PUT",
			"/user/{ID:[0-9]+}",
			editUser,
		},
		// swagger:route DELETE /user/{ID} user deleteUser
		//
		// Delete an user
		//
		// This will delete an user
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
			"DeleteUser",
			"DELETE",
			"/user/{ID:[0-9]+}",
			deleteUser,
		},
		// swagger:route GET /user/google/start user getUserGoogleStart
		//
		// Start a google single sign on process
		//
		// Creates a single sign on for google chain
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
		//       302: ErrRedirect
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetUserGoogleStart",
			"GET",
			"/user/google/start",
			getUserGoogleStart,
		},
	}
	return
}

func getUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserRequest{
		ID: getIntVar(r, "ID"),
	}

	focusUser := &model.User{
		ID: request.ID,
	}

	err = cases.GetUser(focusUser, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &UserResponse{
		User: focusUser,
	}
	content = response
	return
}

func createUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	focusUser := &model.User{}
	err = decodeBody(r, user)
	if err != nil {
		return
	}
	err = cases.CreateUser(focusUser, user)
	if err != nil {
		return
	}
	response := &UserResponse{
		User: focusUser,
	}
	content = response
	return
}

func deleteUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserRequest{
		ID: getIntVar(r, "ID"),
	}

	focusUser := &model.User{
		ID: request.ID,
	}

	err = cases.DeleteUser(focusUser, user)
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

func editUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserEditRequest{
		ID: getIntVar(r, "ID"),
	}

	focusUser := &model.User{
		ID: request.ID,
	}

	err = decodeBody(r, focusUser)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditUser(user, user)
	if err != nil {
		return
	}
	response := &UserResponse{
		User: focusUser,
	}
	content = response
	return
}

func listUser(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	users, err := cases.ListUser(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &UsersResponse{
		Page:  page,
		Users: users,
	}
	content = response
	return
}

func listUserBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	focusUser := &model.User{}

	user.DisplayName = getQuery(r, "displayName")
	users, err := cases.ListUserBySearch(page, focusUser, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	log.Println(users)
	response := &UsersBySearchResponse{
		Page:   page,
		Users:  users,
		Search: focusUser,
	}
	content = response
	return
}

func getUserGoogleStart(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &UserGoogleStartRequest{
		ReturnURL: getQuery(r, "returnURL"),
	}

	b := make([]byte, 16)
	rand.Read(b)

	state := base64.URLEncoding.EncodeToString(b)

	session, _ := cookieStore.Get(r, "sess")
	session.Values["state"] = state
	session.Values["returnURL"] = request.ReturnURL
	session.Save(r, w)

	redirectURL, err := cases.GetUserGoogleStart(state)
	if err != nil {
		err = errors.Wrap(err, "failed to get user google start")
		return
	}
	response := &UserGoogleStartResponse{
		RedirectURL: redirectURL,
	}
	content = response
	return
}
