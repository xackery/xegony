package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) activityRoutes() (routes []*route) {
	routes = []*route{
		{
			"ListActivity",
			"GET",
			"/task/{taskID:[0-9]+}",
			a.listActivity,
		},
		{
			"GetActivity",
			"GET",
			"/task/{taskID}:[0-9]+/activity/{activityID:[0-9]+}",
			a.getActivity,
		},
	}
	return
}

func (a *Web) listActivity(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site      site
		Activitys []*model.Activity
		Task      *model.Task
	}

	site := a.newSite(r)
	site.Page = "activity"
	site.Title = "Activity"

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}

	task := &model.Task{
		ID: taskID,
	}
	activitys, err := a.activityRepo.ListByTask(task, user)
	if err != nil {
		return
	}
	err = a.taskRepo.Get(task, user)
	if err != nil {
		return
	}
	content = Content{
		Site:      site,
		Activitys: activitys,
		Task:      task,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "activity/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("activity", tmp)
	}

	return
}

func (a *Web) getActivity(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site     site
		Activity *model.Activity
		Task     *model.Task
	}

	activityID, err := getIntVar(r, "activityID")
	if err != nil {
		err = errors.Wrap(err, "activityID argument is required")
		return
	}

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
		return
	}
	activity := &model.Activity{
		ActivityID: activityID,
		TaskID:     taskID,
	}
	err = a.activityRepo.Get(activity, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	task := &model.Task{
		ID: taskID,
	}
	err = a.taskRepo.Get(task, user)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		return
	}

	site := a.newSite(r)
	site.Page = "activity"
	site.Title = "Activity"

	content = Content{
		Site:     site,
		Activity: activity,
		Task:     task,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "activity/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("activity", tmp)
	}

	return
}
