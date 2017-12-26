package web

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/box"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
)

type Site struct {
	Title       string //Title of site
	Page        string
	Section     string
	Description string //Description for oprop

}

type Web struct {
	templates     map[string]*Template
	characterRepo *cases.CharacterRepository
	userRepo      *cases.UserRepository
	accountRepo   *cases.AccountRepository
	forumRepo     *cases.ForumRepository
	topicRepo     *cases.TopicRepository
	npcRepo       *cases.NpcRepository
	zoneRepo      *cases.ZoneRepository
	factionRepo   *cases.FactionRepository
}

func (a *Web) Initialize(s storage.Storage, config string) (err error) {
	a.templates = map[string]*Template{}

	if s == nil {
		err = fmt.Errorf("Invalid storage type passed, must be pointer reference")
		return
	}

	a.characterRepo = &cases.CharacterRepository{}
	a.characterRepo.Initialize(s)

	a.userRepo = &cases.UserRepository{}
	a.userRepo.Initialize(s)

	a.accountRepo = &cases.AccountRepository{}
	a.accountRepo.Initialize(s)

	a.forumRepo = &cases.ForumRepository{}
	a.forumRepo.Initialize(s)

	a.topicRepo = &cases.TopicRepository{}
	a.topicRepo.Initialize(s)

	a.npcRepo = &cases.NpcRepository{}
	a.npcRepo.Initialize(s)

	a.zoneRepo = &cases.ZoneRepository{}
	a.zoneRepo.Initialize(s)

	a.factionRepo = &cases.FactionRepository{}
	a.factionRepo.Initialize(s)
	return
}

func (a *Web) NewSite() (site Site) {
	site = Site{
		Title:       "Xegony",
		Description: "Xegony",
	}
	return
}

func (a *Web) Index(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Host string
	}

	site := a.NewSite()
	site.Page = "forum"
	site.Title = "Xegony"

	content := Content{
		Site: site,
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "body", "index.tpl")
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}
		tmp, err = a.loadStandardTemplate(tmp)
		if err != nil {
			a.writeError(w, r, err, http.StatusInternalServerError)
			return
		}

		a.setTemplate("index", tmp)
	}

	a.writeData(w, r, tmp, content, http.StatusOK)
	return
}

func (a *Web) notFound(w http.ResponseWriter, r *http.Request) {
	var err error

	path := "www/" + r.URL.Path[1:]

	if _, err = os.Stat(path); err == nil {
		http.ServeFile(w, r, path)
		return
	}
	var bData []byte
	if bData, err = box.ReadFile(path); err != nil {
		//box open
		a.writeError(w, r, err, http.StatusNotFound)
		return
	}
	reader := bytes.NewReader(bData)
	http.ServeContent(w, r, path, time.Now(), reader)
	return
}

func (a *Web) writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {

	site := a.NewSite()
	site.Page = fmt.Sprintf("%d", statusCode)
	site.Title = "Error"

	tmp := a.getTemplate("")

	var tErr error
	switch statusCode {
	case http.StatusNotFound: //404
		if tmp == nil {
			tmp, tErr = a.loadTemplate(nil, "404", "404.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "404 - Not Found"
	case http.StatusInternalServerError: //500
		if tmp == nil {
			tmp, tErr = a.loadTemplate(nil, "500", "500.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "500 - Internal Server Error"
	}

	type Content struct {
		Site    Site
		Message string
		Url     string
	}

	content := Content{
		Site:    site,
		Message: err.Error(),
	}
	if r != nil {
		content.Url = r.URL.String()
	}
	a.writeData(w, r, tmp, content, statusCode)
}

func (a *Web) writeData(w http.ResponseWriter, r *http.Request, tmp *template.Template, content interface{}, statusCode int) {
	var err error
	w.WriteHeader(statusCode)
	if tmp == nil {
		log.Println("Failed to load template", content)
		return
	}
	if err = tmp.Execute(w, content); err != nil {
		log.Println("Failed to execute template:", err.Error())
		return
	}
}

func getIntVar(r *http.Request, key string) (val int64, err error) {
	vars := mux.Vars(r)
	val, err = strconv.ParseInt(vars[key], 10, 64)
	if err != nil {
		err = &model.ErrInvalidArguments{}
		return
	}
	return
}

func getVar(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
