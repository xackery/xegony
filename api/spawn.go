package api

/*
import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) spawnRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateSpawn",
			"POST",
			"/spawn",
			a.createSpawn,
		},
		{
			"DeleteSpawn",
			"DELETE",
			"/spawn/{spawnID:[0-9]+}",
			a.deleteSpawn,
		},
		{
			"EditSpawn",
			"PUT",
			"/spawn/{spawnID:[0-9]+}",
			a.editSpawn,
		},
		{
			"GetSpawn",
			"GET",
			"/spawn/{spawnID:[0-9]+}",
			a.getSpawn,
		},
		{
			"ListSpawn",
			"GET",
			"/spawn",
			a.listSpawn,
		},
	}
	return
}
func (a *API) getSpawn(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

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
		return
}

func (a *API) createSpawn(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	if err = IsAdmin(r); err != nil {
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

		return
}

func (a *API) deleteSpawn(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {


	if err = IsAdmin(r); err != nil {
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

func (a *API) editSpawn(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {


	if err = IsModerator(r); err != nil {
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
		return
}

func (a *API) listSpawn(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	spawns, err := a.spawnRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
				return
	}
		return
}
*/
