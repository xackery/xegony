package web

import (
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site      Site
		Activitys []*model.Activity
		Task      *model.Task
	}

	site := a.NewSite(r)
	site.Page = "activity"
	site.Title = "Activity"

	if strings.ToLower(getVar(r, "taskId")) == "create" {
		a.CreateTask(w, r)
		return
	}

	taskId, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	activitys, err := a.activityRepo.List(taskId)
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	task, err := a.taskRepo.Get(taskId)
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

func (a *Web) GetActivity(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site     Site
		Activity *model.Activity
		Task     *model.Task
	}

	activityId, err := getIntVar(r, "activityId")
	if err != nil {
		err = errors.Wrap(err, "activityId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	taskId, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	activity, err := a.activityRepo.Get(taskId, activityId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	task, err := a.taskRepo.Get(taskId)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
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
