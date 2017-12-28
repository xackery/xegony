package api

import (
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Api) GetGoal(w http.ResponseWriter, r *http.Request) {

	listId, err := getIntVar(r, "listId")
	if err != nil {
		err = errors.Wrap(err, "listId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	entryId, err := getIntVar(r, "entryId")
	if err != nil {
		err = errors.Wrap(err, "entryId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}
	goal, err := a.goalRepo.Get(listId, entryId)
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

func (a *Api) CreateGoal(w http.ResponseWriter, r *http.Request) {
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

func (a *Api) DeleteGoal(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsAdmin(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	listId, err := getIntVar(r, "listId")
	if err != nil {
		err = errors.Wrap(err, "listId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	entryId, err := getIntVar(r, "entryId")
	if err != nil {
		err = errors.Wrap(err, "entryId argument is required")
		writeError(w, r, err, http.StatusBadRequest)
		return
	}

	err = a.goalRepo.Delete(listId, entryId)
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

func (a *Api) EditGoal(w http.ResponseWriter, r *http.Request) {
	var err error

	if err = IsModerator(r); err != nil {
		writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	listId, err := getIntVar(r, "listId")
	if err != nil {
		err = errors.Wrap(err, "listId argument is required")
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

	err = a.goalRepo.Edit(listId, goal)
	if err != nil {
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, goal, http.StatusOK)
	return
}

func (a *Api) ListGoal(w http.ResponseWriter, r *http.Request) {
	goals, err := a.goalRepo.List()
	if err != nil {
		err = errors.Wrap(err, "Request error")
		writeError(w, r, err, http.StatusInternalServerError)
		return
	}
	writeData(w, r, goals, http.StatusOK)
	return
}
