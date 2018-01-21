package web

import (
	"html/template"
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/model"
)

func (a *Web) taskRoutes() (routes []*route) {
	routes = []*route{
		//Task
		{
			"GetTask",
			"GET",
			"/task/{taskID:[0-9]+}/details",
			a.getTask,
		},
		{
			"ListTask",
			"GET",
			"/task",
			a.listTask,
		},
		{
			"CreateTask",
			"GET",
			"/task/{taskID:[0-9]+}/create",
			a.listActivity,
		},
	}
	return
}

func (a *Web) listTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site  site
		Tasks []*model.Task
	}

	site := a.newSite(r)
	site.Page = "task"
	site.Title = "Task"

	tasks, err := a.taskRepo.List(user)
	if err != nil {
		return
	}
	content = Content{
		Site:  site,
		Tasks: tasks,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "task/list.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("task", tmp)
	}

	return
}

func (a *Web) getTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
		Task *model.Task
	}

	taskID, err := getIntVar(r, "taskID")
	if err != nil {
		err = errors.Wrap(err, "taskID argument is required")
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
	site.Page = "task"
	site.Title = "Task"

	content = Content{
		Site: site,
		Task: task,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "task/get.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("task", tmp)
	}

	return
}

func (a *Web) createTask(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {

	type Content struct {
		Site site
	}

	site := a.newSite(r)
	site.Page = "task"
	site.Title = "Task"

	content = Content{
		Site: site,
	}

	tmp = a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "task/create.tpl")
		if err != nil {
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			return
		}

		a.setTemplate("task", tmp)
	}

	return
}
