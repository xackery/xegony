package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

// CharacterParams is a list of parameters used for character
// swagger:parameters deleteCharacter editCharacter getCharacter
type CharacterParams struct {
	//CharacterID to get information about
	// in: path
	CharacterID int64 `json:"characterID"`
	//todo: pagination
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
		//       200: Characters
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListCharacter",
			"GET",
			"/character",
			listCharacter,
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
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"CreateCharacter",
			"POST",
			"/character",
			createCharacter,
		},
		// swagger:route GET /character/{characterID} character getCharacter
		//
		// Get an character
		//
		// This will get an individual character available characters by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: Character
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetCharacter",
			"GET",
			"/character/{characterID:[0-9]+}",
			getCharacter,
		},
		// swagger:route PUT /character/{characterID} character editCharacter
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
		//		 200: ErrNoContent
		//       204: ErrNoContent
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"EditCharacter",
			"PUT",
			"/character/{characterID:[0-9]+}",
			editCharacter,
		},
		// swagger:route DELETE /character/{characterID} character deleteCharacter
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
			"/character/{characterID:[0-9]+}",
			deleteCharacter,
		},
	}
	return
}

func getCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	characterReq := &CharacterParams{}

	characterReq.CharacterID = getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "characterID argument is required")
		return
	}
	character := &model.Character{
		ID: characterReq.CharacterID,
	}
	err = cases.GetCharacter(character, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = character
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
	content = character
	return
}

func deleteCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	characterReq := &CharacterParams{}
	characterReq.CharacterID = getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "characterID argument is required")
		return
	}
	character := &model.Character{
		ID: characterReq.CharacterID,
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
	characterReq := &CharacterParams{}
	characterReq.CharacterID = getIntVar(r, "characterID")
	if err != nil {
		err = errors.Wrap(err, "characterID argument is required")
		return
	}

	character := &model.Character{
		ID: characterReq.CharacterID,
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
	content = character
	return
}

func listCharacter(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	characters, err := cases.ListCharacter(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = characters
	return
}
