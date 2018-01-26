package api

/*
import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) spawnEntryRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateSpawnEntry",
			"POST",
			"/spawn/{spawnID:[0-9]+}/entry",
			a.createSpawnEntry,
		},
		{
			"DeleteSpawnEntry",
			"DELETE",
			"/spawn/{spawnID:[0-9]+}/entry/{spawnEntryID:[0-9]+}",
			a.deleteSpawnEntry,
		},
		{
			"EditSpawnEntry",
			"PUT",
			"/spawn/{spawnID:[0-9]+}/entry/{spawnEntryID::[0-9]+}",
			a.editSpawnEntry,
		},
		{
			"GetSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry/{spawnEntryID:[0-9]+}",
			a.getSpawnEntry,
		},
		{
			"ListSpawnEntry",
			"GET",
			"/spawn/{spawnID:[0-9]+}/entry",
			a.listSpawnEntry,
		},
	}
	return
}

func (a *API) getSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	spawnGroupID, err := getIntVar(r, "spawnGroupID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
				return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
				return
	}
	spawnEntry := &model.SpawnEntry{
		NpcID:        npcID,
		SpawngroupID: spawnGroupID,
	}
	err = a.spawnEntryRepo.Get(spawnEntry, user)
	if err != nil {
		if err == sql.ErrNoRows {
						return
		}
		err = errors.Wrap(err, "Request error")
				return
	}
		return
}

func (a *API) createSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	if err = IsAdmin(r); err != nil {
				return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
				return
	}
	err = a.spawnEntryRepo.Create(spawnEntry, user)
	if err != nil {
				return
	}

		return
}

func (a *API) deleteSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {


	if err = IsAdmin(r); err != nil {
				return
	}

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
				return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
				return
	}

	spawn := &model.Spawn{
		ID:    spawnID,
		NpcID: npcID,
	}
	_, err = a.spawnEntryRepo.Delete(spawn, user)
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


	if err = IsModerator(r); err != nil {
				return
	}
	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
				return
	}

	npcID, err := getIntVar(r, "npcID")
	if err != nil {
		err = errors.Wrap(err, "npcID argument is required")
				return
	}

	spawnEntry := &model.SpawnEntry{}
	err = decodeBody(r, spawnEntry)
	if err != nil {
		err = errors.Wrap(err, "Request error")
				return
	}

	_, err = a.spawnEntryRepo.Edit(spawnID, npcID, spawnEntry)
	if err != nil {
				return
	}
		return
}

func (a *API) listSpawnEntry(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {

	spawnID, err := getIntVar(r, "spawnID")
	if err != nil {
		err = errors.Wrap(err, "spawnID argument is required")
				return
	}

	spawnEntrys, _, err := a.spawnEntryRepo.List(spawnID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
				return
	}
		return
}
*/
