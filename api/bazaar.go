package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) bazaarRoutes() (routes []*route) {
	routes = []*route{
		{
			"CreateBazaar",
			"POST",
			"/bazaar",
			a.createBazaar,
		},
		{
			"DeleteBazaar",
			"DELETE",
			"/bazaar/{bazaarID:[0-9]+}",
			a.deleteBazaar,
		},
		{
			"EditBazaar",
			"PUT",
			"/bazaar/{bazaarID:[0-9]+}",
			a.editBazaar,
		},
		{
			"GetBazaar",
			"GET",
			"/bazaar/{bazaarID:[0-9]+}",
			a.getBazaar,
		},
		{
			"ListBazaar",
			"GET",
			"/bazaar",
			a.listBazaar,
		},
	}
	return
}

func (a *API) getBazaar(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	bazaarID, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		return
	}
	bazaar := &model.Bazaar{
		ID: bazaarID,
	}
	err = a.bazaarRepo.Get(bazaar, user)
	if err != nil {
		if err == sql.ErrNoRows {
			return
		}
		err = errors.Wrap(err, "Request error")
		return
	}
	content = bazaar
	return
}

func (a *API) createBazaar(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	bazaar := &model.Bazaar{}
	err = decodeBody(r, bazaar)
	if err != nil {
		return
	}
	err = a.bazaarRepo.Create(bazaar, user)
	if err != nil {
		return
	}
	content = bazaar
	return
}

func (a *API) deleteBazaar(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	bazaarID, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		return
	}

	bazaar := &model.Bazaar{
		ID: bazaarID,
	}

	err = a.bazaarRepo.Delete(bazaar, user)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			return
		default:
			err = errors.Wrap(err, "Request failed")
		}
		return
	}
	content = bazaar
	return
}

func (a *API) editBazaar(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {

	bazaarID, err := getIntVar(r, "bazaarID")
	if err != nil {
		err = errors.Wrap(err, "bazaarID argument is required")
		return
	}

	bazaar := &model.Bazaar{}
	err = decodeBody(r, bazaar)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	bazaar.ID = bazaarID
	err = a.bazaarRepo.Edit(bazaar, user)
	if err != nil {
		return
	}
	content = bazaar
	return
}

func (a *API) listBazaar(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	bazaars, err := a.bazaarRepo.List(user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}
	content = bazaars
	return

}
