package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) spawnRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spawn spawn listSpawn
		//
		// Lists spawns
		//
		// This will show all available spawns by default.
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
		//       200: Spawns
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawn",
			"GET",
			"/spawn",
			a.listSpawn,
		},
		// swagger:route POST /spawn spawn createSpawn
		//
		// Create an spawn
		//
		// This will create an spawn
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
			"CreateSpawn",
			"POST",
			"/spawn",
			a.createSpawn,
		},
		// swagger:route GET /spawn/{spawnID} spawn getSpawn
		//
		// Get an spawn
		//
		// This will get an individual spawn available spawns by default.
		//
		//     Responses:
		//       default: ErrInternal
		//       200: Spawn
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpawn",
			"GET",
			"/spawn/{spawnID:[0-9]+}",
			a.getSpawn,
		},
		// swagger:route PUT /spawn/{spawnID} spawn editSpawn
		//
		// Edit an spawn
		//
		// This will edit an spawn
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
			"EditSpawn",
			"PUT",
			"/spawn/{spawnID:[0-9]+}",
			a.editSpawn,
		},
		// swagger:route DELETE /spawn/{spawnID} spawn deleteSpawn
		//
		// Delete an spawn
		//
		// This will delete an spawn
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
			"DeleteSpawn",
			"DELETE",
			"/spawn/{spawnID:[0-9]+}",
			a.deleteSpawn,
		},
	}
	return
}
func (a *API) getSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawn := &model.Spawn{
		ID: spawnID,
	}
	err = a.spawnRepo.Get(spawn, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = spawn
	return
}

func (a *API) createSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawn := &model.Spawn{}
	err = decodeBody(r, spawn)
	if err != nil {
		return
	}
	err = a.spawnRepo.Create(spawn, user)
	if err != nil {
		return
	}
	content = spawn
	return
}

func (a *API) deleteSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawn := &model.Spawn{
		ID: spawnID,
	}

	err = a.spawnRepo.Delete(spawn, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	return
}

func (a *API) editSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawn := &model.Spawn{}
	err = decodeBody(r, spawn)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	spawn.ID = spawnID

	err = a.spawnRepo.Edit(spawn, user)
	if err != nil {
		return
	}
	content = spawn
	return
}

func (a *API) listSpawn(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spawns, err := a.spawnRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = spawns
	return
}
