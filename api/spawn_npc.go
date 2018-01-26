package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) spawnNpcRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spawn/{spawnID}/npc spawn listSpawnNpc
		//
		// Lists spawn Npcs
		//
		// This will show all available npcs for a spawns by default.
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
		//       200: SpawnNpcs
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"ListSpawnNpc",
			"GET",
			"/spawn/{spawnID:[0-9]+}/npc",
			a.listSpawnNpc,
		},
		// swagger:route POST /spawn/{spawnID}/npc spawn createSpawnNpc
		//
		// Create an npc spawn
		//
		// This will create an npc spawn
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
			"/spawn/{spawnID:[0-9]+}/npc",
			a.createSpawnNpc,
		},
		// swagger:route GET /spawn/{spawnID}/npc/{spawnNpcID} spawn getSpawnNpc
		//
		// Get a npc spawn
		//
		// This will get an individual npc in a spawn
		//
		//     Responses:
		//       default: ErrInternal
		//       200: SpawnNpc
		//       400: ErrValidation
		//		 401: ErrPermission
		{
			"GetSpawnNpc",
			"GET",
			"/spawn/{spawnID:[0-9]+}/npc/{spawnNpcID:[0-9]+}",
			a.getSpawnNpc,
		},
		// swagger:route PUT /spawn/{spawnID}/npc/{spawnNpcID} spawn editSpawnNpc
		//
		// Edit an npc spawn
		//
		// This will edit an npc spawn
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
			"EditSpawnNpc",
			"PUT",
			"/spawn/{spawnID:[0-9]+}/npc/{spawnNpcID:[0-9]+}",
			a.editSpawnNpc,
		},
		// swagger:route DELETE /spawn/{spawnID}/npc/{spawnNpcID} spawn deleteSpawn
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
			"DeleteSpawnNpc",
			"DELETE",
			"/spawn/{spawnID:[0-9]+}/npc/{spawnNpcID:[0-9]+}",
			a.deleteSpawnNpc,
		},
	}
	return
}

func (a *API) getSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawnNpcID, err := getIntVar(r, "spawnNpcID")
	if err != nil {
		err = errors.Wrap(err, "spawnNpcID argument is required")
		return
	}

	spawnNpc := &model.SpawnNpc{
		NpcID:   spawnNpcID,
		SpawnID: spawnID,
	}
	err = a.spawnNpcRepo.Get(spawnNpc, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = spawnNpc
	return
}

func (a *API) createSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawnNpc := &model.SpawnNpc{
		SpawnID: spawnID,
	}
	err = decodeBody(r, spawnNpc)
	if err != nil {
		return
	}
	err = a.spawnNpcRepo.Create(spawnNpc, user)
	if err != nil {
		return
	}
	content = spawnNpc
	return
}

func (a *API) deleteSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}
	spawnNpcID, err := getIntVar(r, "spawnNpcID")
	if err != nil {
		err = errors.Wrap(err, "spawnNpcID argument is required")
		return
	}

	spawnNpc := &model.SpawnNpc{
		NpcID:   spawnNpcID,
		SpawnID: spawnID,
	}

	err = a.spawnNpcRepo.Delete(spawnNpc, user)
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

func (a *API) editSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawnNpcID, err := getIntVar(r, "spawnNpcID")
	if err != nil {
		err = errors.Wrap(err, "spawnNpcID argument is required")
		return
	}

	spawnNpc := &model.SpawnNpc{}
	err = decodeBody(r, spawnNpc)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	spawnNpc.NpcID = spawnNpcID
	spawnNpc.SpawnID = spawnID
	err = a.spawnNpcRepo.Edit(spawnNpc, user)
	if err != nil {
		return
	}
	content = spawnNpc
	return
}

func (a *API) listSpawnNpc(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spawnNpcs, err := a.spawnNpcRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = spawnNpcs
	return
}
