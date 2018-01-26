package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) spawnEntryRoutes() (routes []*route) {

	routes = []*route{
		// swagger:route GET /spawn/{spawnID}/entry spawn listSpawn
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
			"ListSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry",
			a.listSpawnEntry,
		},
		// swagger:route POST /spawn/{spawnID}/entry spawn createSpawn
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
			"CreateSpawnEntry",
			"POST",
			"/spawn/{spawnID:[0-9]+}/entry",
			a.createSpawnEntry,
		},
		// swagger:route GET /spawn/{spawnID}/entry/{spawnEntryID} spawn getSpawn
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
			"GetSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry/{spawnEntryID:[0-9]+}",
			a.getSpawnEntry,
		},
		// swagger:route PUT /spawn/{spawnID}/entry/{spawnEntryID} spawn editSpawn
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
			"EditSpawnEntry",
			"PUT",
			"/spawn/{spawnID:[0-9]+}/entry/{spawnEntryID:[0-9]+}",
			a.editSpawnEntry,
		},
		// swagger:route DELETE /spawn/{spawnID}/entry/{spawnEntryID} spawn deleteSpawn
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
			"DeleteSpawnEntry",
			"DELETE",
			"/spawn/{spawnID:[0-9]+}/entry/{spawnEntryID:[0-9]+}",
			a.deleteSpawnEntry,
		},
	}
	return
}

func (a *API) getSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawnEntryID, err := getIntVar(r, "spawnEntryID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
		return
	}

	spawnEntry := &model.SpawnEntry{
		ID:      spawnEntryID,
		SpawnID: spawnID,
	}
	err = a.spawnEntryRepo.Get(spawnEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = spawnEntry
	return
}

func (a *API) createSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawnEntry := &model.SpawnEntry{
		SpawnID: spawnID,
	}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		return
	}
	err = a.spawnEntryRepo.Create(spawnEntry, user)
	if err != nil {
		return
	}
	content = spawnEntry
	return
}

func (a *API) deleteSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}
	spawnEntryID, err := getIntVar(r, "spawnEntryID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
		return
	}

	spawnEntry := &model.SpawnEntry{
		ID:      spawnEntryID,
		SpawnID: spawnID,
	}

	err = a.spawnEntryRepo.Delete(spawnEntry, user)
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

func (a *API) editSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	err = user.IsAdmin()
	if err != nil {
		return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
		return
	}

	spawnEntryID, err := getIntVar(r, "spawnEntryID")
	if err != nil {
		err = errors.Wrap(err, "spawnEntryID argument is required")
		return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	spawnEntry.ID = spawnEntryID
	spawnEntry.SpawnID = spawnID
	err = a.spawnEntryRepo.Edit(spawnEntry, user)
	if err != nil {
		return
	}
	content = spawnEntry
	return
}

func (a *API) listSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	spawnEntrys, err := a.spawnEntryRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = spawnEntrys
	return
}
