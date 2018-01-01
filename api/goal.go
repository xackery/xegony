package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *API) getGoal(w http.ResponseWriter, r *http.Request) {

	listID, err := getIntVar(r, "listID")
	if err != nil {
		err = errors.Wrap(err, "listID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	entryID, err := getIntVar(r, "entryID")
	if err != nil {
		err = errors.Wrap(err, "entryID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	goal, err := a.goalRepo.Get(listID, entryID)
	if err != nil {
		if err == sql.ErrNoRows {
			writeData(w, r, "", http.StatusOK)
			return
		}
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	writeData(w, r, goal, http.StatusOK)
	return
}

func (a *API) createGoal(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	goal := &model.Goal{}
	err = decodeBody(r, goal)
	if err != nil {
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}
	err = a.goalRepo.Create(goal)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}

	writeData(w, r, goal, http.StatusCreated)
	return
}

func (a *API) deleteGoal(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	listID, err := getIntVar(r, "listID")
	if err != nil {
		err = errors.Wrap(err, "listID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	entryID, err := getIntVar(r, "entryID")
	if err != nil {
		err = errors.Wrap(err, "entryID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.goalRepo.Delete(listID, entryID)
	if err != nil {
		switch errors.Cause(err).(type) {
		case *model.ErrNoContent:
			writeData(w, r, nil, http.StatusNotModified)
			return
		default:
			err = errors.Wrap(err, "Request failed")
			writeError(w, r, err, http.StatusInternalServerError)
		}
		return
	}
	writeData(w, r, nil, http.StatusNoContent)
	return
}

func (a *API) editGoal(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	listID, err := getIntVar(r, "listID")
	if err != nil {
		err = errors.Wrap(err, "listID argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	goal := &model.Goal{}
	err = decodeBody(r, goal)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusMethodNotAllowed)
		return
	}

	err = a.goalRepo.Edit(listID, goal)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, goal, http.StatusOK)
	return
}

func (a *API) listGoal(w http.ResponseWriter, r *http.Request) {
	goals, err := a.goalRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, goals, http.StatusOK)
	return
}
