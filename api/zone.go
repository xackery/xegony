package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) zoneRoutes() (routes []*route) {
	routes = []*route{}
	return
}

func (a *API) getZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		return
	}
	zone := &model.Zone{
		ZoneIDNumber: zoneID,
	}
	err = a.zoneRepo.Get(zone, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = zone
	return
}

func (a *API) createZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	zone := &model.Zone{}
	err = decodeBody(r, zone)
	if err != nil {
		return
	}

	err = a.zoneRepo.Create(zone, user)
	if err != nil {
		return
	}
	content = zone
	return
}

func (a *API) deleteZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		return
	}

	zone := &model.Zone{
		ZoneIDNumber: zoneID,
	}
	err = a.zoneRepo.Delete(zone, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = zone
	return
}

func (a *API) editZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	zoneID, err := getIntVar(r, "zoneID")
	if err != nil {
		err = errors.Wrap(err, "zoneID argument is required")
		return
	}

	zone := &model.Zone{}
	err = decodeBody(r, zone)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	zone.ZoneIDNumber = zoneID
	err = a.zoneRepo.Edit(zone, user)
	if err != nil {
		return
	}
	content = zone
	return
}

func (a *API) listZone(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	zones, err := a.zoneRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = zones
	return
}
