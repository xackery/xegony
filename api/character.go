package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// CharacterRequest is a list of parameters used for character
// swagger:parameters deleteCharacter editCharacter getCharacter
type CharacterRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
}

// CharacterResponse is what endpoints respond with
// swagger:response
type CharacterResponse struct {
	Character *model.Character `json:"character,omitempty"`
}

// CharacterInventoryResponse is what endpoints respond with
// swagger:response
type CharacterInventoryResponse struct {
	Character *model.Character `json:"character,omitempty"`
	Items     model.Items      `json:"items,omitempty"`
}

// CharacterCreateRequest is the body parameters for creating an character
// swagger:parameters createCharacter
type CharacterCreateRequest struct {
	// Character details to create
	// in: body
	Character *model.Character `json:"character"`
}

// CharacterEditRequest is the body parameters for creating an character
// swagger:parameters editCharacter
type CharacterEditRequest struct {
	// ID to get information about
	// in: path
	// example: 74887
	ID int64 `json:"ID"`
	// Character details to edit
	// in: body
	Character *model.Character `json:"character"`
}

// CharactersRequest is a list of parameters used for character
// swagger:parameters listCharacter
type CharactersRequest struct {
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

// CharactersResponse is a general response to a request
// swagger:response
type CharactersResponse struct {
	Page       *model.Page      `json:"page,omitempty"`
	Characters model.Characters `json:"characters,omitempty"`
}

// CharactersBySearchRequest is a list of parameters used for character
// swagger:parameters listCharacterBySearch
type CharactersBySearchRequest struct {
	// Name is which character to get information about
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

// CharactersBySearchResponse is a general response to a request
// swagger:response
type CharactersBySearchResponse struct {
	Search     *model.Character `json:"search,omitempty"`
	Page       *model.Page      `json:"page,omitempty"`
	Characters model.Characters `json:"characters,omitempty"`
}

func characterRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /character character listCharacter
		//
		// Lists characters
		//
		// This will show all available characters by default.
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
		//       200: CharactersResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListCharacter",
			"GET",
			"/character",
			listCharacter,
		},
		// swagger:route GET /character/search character listCharacterBySearch
		//
		// Search characters by name
		//
		// This will show all available characters by default.
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
		//       200: CharactersBySearchResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListCharacterBySearch",
			"GET",
			"/character/search",
			listCharacterBySearch,
		},
		// swagger:route POST /character character createCharacter
		//
		// Create an character
		//
		// This will create an character
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: CharacterResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateCharacter",
			"POST",
			"/character",
			createCharacter,
		},
		// swagger:route GET /character/{ID} character getCharacter
		//
		// Get an character
		//
		// This will get an individual character available characters by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: CharacterResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetCharacter",
			"GET",
			"/character/{ID:[0-9]+}",
			getCharacter,
		},
		// swagger:route GET /character/{ID}/inventory character getCharacterInventory
		//
		// Get a character's inventory
		//
		// This will get an individual character as well as inventory
		//
		//     Responses:
		//       default: ErrInternal
		//       200: CharacterInventoryResponse
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetCharacter",
			"GET",
			"/character/{ID:[0-9]+}/inventory",
			getCharacterInventory,
		},
		// swagger:route PUT /character/{ID} character editCharacter
		//
		// Edit an character
		//
		// This will edit an character
		//
		//     Security:
		//       apiKey:
		//
		//     Responses:
		//       default: ErrInternal
		//		 200: CharacterResponse
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditCharacter",
			"PUT",
			"/character/{ID:[0-9]+}",
			editCharacter,
		},
		// swagger:route DELETE /character/{ID} character deleteCharacter
		//
		// Delete an character
		//
		// This will delete an character
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
			"DeleteCharacter",
			"DELETE",
			"/character/{ID:[0-9]+}",
			deleteCharacter,
		},
	}
	return
}

func getCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &CharacterRequest{
		ID: getIntVar(r, "ID"),
	}

	character := &model.Character{
		ID: request.ID,
	}

	err = cases.GetCharacter(character, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	response := &CharacterResponse{
		Character: character,
	}
	content = response
	return
}

func getCharacterInventory(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &CharacterRequest{
		ID: getIntVar(r, "ID"),
	}

	character := &model.Character{
		ID: request.ID,
	}

	err = cases.GetCharacter(character, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}

	/*err = cases.ListItemByCharacter(character, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get inventory")
		return
	}*/
	response := &CharacterResponse{
		Character: character,
	}
	content = response
	return
}

func createCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	character := &model.Character{}
	err = decodeBody(r, character)
	if err != nil {
		return
	}
	err = cases.CreateCharacter(character, user)
	if err != nil {
		return
	}
	response := &CharacterResponse{
		Character: character,
	}
	content = response
	return
}

func deleteCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &CharacterRequest{
		ID: getIntVar(r, "ID"),
	}

	character := &model.Character{
		ID: request.ID,
	}

	err = cases.DeleteCharacter(character, user)
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

func editCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	request := &CharacterEditRequest{
		ID: getIntVar(r, "ID"),
	}

	character := &model.Character{
		ID: request.ID,
	}

	err = decodeBody(r, character)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	err = cases.EditCharacter(character, user)
	if err != nil {
		return
	}
	response := &CharacterResponse{
		Character: character,
	}
	content = response
	return
}

func listCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	characters, err := cases.ListCharacter(page, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &CharactersResponse{
		Page:       page,
		Characters: characters,
	}
	content = response
	return
}

func listCharacterBySearch(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	page := &model.Page{
		Offset:       getIntQuery(r, "offset"),
		Limit:        getIntQuery(r, "limit"),
		OrderBy:      getQuery(r, "orderBy"),
		IsDescending: getIntQuery(r, "isDescending"),
	}
	character := &model.Character{
		Name: getQuery(r, "name"),
	}
	characters, err := cases.ListCharacterBySearch(page, character, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	response := &CharactersBySearchResponse{
		Page:       page,
		Characters: characters,
		Search:     character,
	}
	content = response
	return
}
