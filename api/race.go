package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// RaceRequest is a list of parameters used for race
// swagger:parameters deleteRace editRace getRace
type RaceRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// RaceResponse is what endpoints respond with
// swagger:response
type RaceResponse struct {
	Race *model.Race `json:"race,omitempty"`
}

// RaceCreateRequest is the body parameters for creating an race
// swagger:parameters createRace
type RaceCreateRequest struct {
	// Race details to create
	// in: body
	Race *model.Race `json:"race"`
}

// RaceEditRequest is the body parameters for creating an race
// swagger:parameters editRace
type RaceEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Race details to edit
	// in: body
	Race *model.Race `json:"race"`
}

// RacesRequest is a list of parameters used for race
// swagger:parameters listRace
type RacesRequest struct {
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

// RacesResponse is a general response to a request
// swagger:response
type RacesResponse struct {
	Page  *model.Page `json:"page,omitempty"`
	Races model.Races `json:"races,omitempty"`
}

// RacesBySearchRequest is a list of parameters used for race
// swagger:parameters listRaceBySearch
type RacesBySearchRequest struct {
	// Name is which race to get information about
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

// RacesBySearchResponse is a general response to a request
// swagger:response
type RacesBySearchResponse struct {
	Search *model.Race `json:"search,omitempty"`
	Page   *model.Page `json:"page,omitempty"`
	Races  model.Races `json:"races,omitempty"`
}

func raceRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /race race listRace
		//
		// Lists races
		//
		// This will show all available races by default.
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
		//       200: RacesResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListRace",
			"GET",
			"/race",
			listRace,
		},
		// swagger:route GET /race/search race listRaceBySearch
		//
		// Search races by name
		//
		// This will show all available races by default.
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
		//       200: RacesBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListRaceBySearch",
			"GET",
			"/race/search",
			listRaceBySearch,
		},
		// swagger:route POST /race race createRace
		//
		// Create an race
		//
		// This will create an race
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: RaceResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateRace",
			"POST",
			"/race",
			createRace,
		},
		// swagger:route GET /race/{ID} race getRace
		//
		// Get an race
		//
		// This will get an individual race available races by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: RaceResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetRace",
			"GET",
			"/race/{ID:[0-9]+}",
			getRace,
		},
		// swagger:route PUT /race/{ID} race editRace
		//
		// Edit an race
		//
		// This will edit an race
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: RaceResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditRace",
			"PUT",
			"/race/{ID:[0-9]+}",
			editRace,
		},
		// swagger:route DELETE /race/{ID} race deleteRace
		//
		// Delete an race
		//
		// This will delete an race
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
			"DeleteRace",
			"DELETE",
			"/race/{ID:[0-9]+}",
			deleteRace,
		},
	}
	return
}

func getRace(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RaceRequest{
		ID: getIntVar(r, "ID"),
	}

	race := &model.Race{
		ID: request.ID,
	}

	err = cases.GetRace(race, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &RaceResponse{
		Race: race,
	}
	content = response
	return
}

func createRace(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	race := &model.Race{}
	err = decodeBody(r, race)
	if err != nil {
		return
	}
	err = cases.CreateRace(race, user)
	if err != nil {
		return
	}
	response := &RaceResponse{
		Race: race,
	}
	content = response
	return
}

func deleteRace(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RaceRequest{
		ID: getIntVar(r, "ID"),
	}

	race := &model.Race{
		ID: request.ID,
	}

	err = cases.DeleteRace(race, user)
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

func editRace(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &RaceEditRequest{
		ID: getIntVar(r, "ID"),
	}

	race := &model.Race{
		ID: request.ID,
	}

	err = decodeBody(r, race)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditRace(race, user)
	if err != nil {
		return
	}
	response := &RaceResponse{
		Race: race,
	}
	content = response
	return
}

func listRace(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	races, err := cases.ListRace(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &RacesResponse{
		Page:  page,
		Races: races,
	}
	content = response
	return
}

func listRaceBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	race := &model.Race{
		Name: getQuery(r, "name"),
	}
	races, err := cases.ListRaceBySearch(page, race, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	log.Println(races)
	response := &RacesBySearchResponse{
		Page:   page,
		Races:  races,
		Search: race,
	}
	content = response
	return
}
