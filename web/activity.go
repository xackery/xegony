package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) listActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      site
		Activitys []*model.Activity
		Task      *model.Task
	}

	site := a.newSite(r)
	site.Page = "activity"
	site.Title = "Activity"

	if strings.ToLower(getVar(r, "taskID")) == "create" {
		a.createTask(w, r)
		return
	}

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	activitys, err := a.activityRepo.List(taskID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	task, err := a.taskRepo.Get(taskID)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:      site,
		Activitys: activitys,
		Task:      task,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "activity/list.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("activity", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) getActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     site
		Activity *model.Activity
		Task     *model.Task
	}

	activityID, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activity, err := a.activityRepo.Get(taskID, activityID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	task, err := a.taskRepo.Get(taskID)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.newSite(r)
	site.Page = "activity"
	site.Title = "Activity"

	content := Content{
		Site:     site,
		Activity: activity,
		Task:     task,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "activity/get.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("activity", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
