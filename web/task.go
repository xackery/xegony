package web

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/model"
)

func (a *Web) ListTask(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site  Site
		Tasks []*model.Task
	}

	site := a.NewSite(r)
	site.Page = "task"
	site.Title = "Task"

	tasks, err := a.taskRepo.List()
	if err != nil {
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	content := Content{
		Site:  site,
		Tasks: tasks,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "listtask.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("task", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) GetTask(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Task *model.Task
	}

	id, err := getIntVar(r, "taskId")
	if err != nil {
		err = errors.Wrap(err, "taskId argument is required")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}
	task, err := a.taskRepo.Get(id)
	if err != nil {
		err = errors.Wrap(err, "Request error")
		a.writeError(w, r, err, http.StatusBadRequest)
		return
	}

	site := a.NewSite(r)
	site.Page = "task"
	site.Title = "Task"

	content := Content{
		Site: site,
		Task: task,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "gettask.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("task", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) CreateTask(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
	}

	site := a.NewSite(r)
	site.Page = "task"
	site.Title = "Task"

	if err = api.IsAdmin(r); err != nil {
		a.writeError(w, r, err, http.StatusUnauthorized)
		return
	}

	content := Content{
		Site: site,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "createtask.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("task", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}
