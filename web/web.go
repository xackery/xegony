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
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/box"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
)

type Site struct {
	Title       string //Title of site
	Name        string
	Page        string
	Section     string
	Description string //Description for oprop
	User        *model.User
	PageNumber  int64
	PageSize    int64
	ResultCount int64
}

type Web struct {
	accountRepo        *cases.AccountRepository
	activityRepo       *cases.ActivityRepository
	bazaarRepo         *cases.BazaarRepository
	characterRepo      *cases.CharacterRepository
	factionRepo        *cases.FactionRepository
	forumRepo          *cases.ForumRepository
	itemRepo           *cases.ItemRepository
	lootDropRepo       *cases.LootDropRepository
	lootDropEntryRepo  *cases.LootDropEntryRepository
	lootTableRepo      *cases.LootTableRepository
	lootTableEntryRepo *cases.LootTableEntryRepository
	npcRepo            *cases.NpcRepository
	npcLootRepo        *cases.NpcLootRepository
	postRepo           *cases.PostRepository
	taskRepo           *cases.TaskRepository
	templates          map[string]*Template
	topicRepo          *cases.TopicRepository
	userRepo           *cases.UserRepository
	zoneRepo           *cases.ZoneRepository
}

func (s Site) PageList() template.HTML {
	page := `<div class="btn-group pull-right">`
	curPage := s.PageNumber
	var curElement int64
	if s.PageNumber > 0 {
		page += fmt.Sprintf("\n"+`<button type="button" class="btn btn-default"><a href="/item?pageNumber=%d"><i class="fa fa-chevron-left"></i></a></button>`, s.PageNumber-1)
	}

	curElement = (s.PageNumber - 6) * s.PageSize
	curPage -= 6
	numCount := 0

	for curElement <= s.ResultCount {
		if curPage < 0 {
			curPage++
			curElement += s.PageSize
			continue
		}
		curPage++
		if curPage == s.PageNumber {
			page += fmt.Sprintf("\n"+` <button class="btn btn-default active"><a href="/item/?pageNumber=%d">%d</a></button>`, curPage, curPage)
		} else {
			page += fmt.Sprintf("\n"+` <button class="btn btn-default"><a href="/item/?pageNumber=%d">%d</a></button>`, curPage, curPage)
		}
		curElement += s.PageSize
		numCount++
		if numCount >= 10 {
			break
		}
	}
	if s.PageNumber*s.PageSize < s.ResultCount {
		page += fmt.Sprintf("\n"+`<button type="button" class="btn btn-default"><a href="/item?pageNumber=%d"><i class="fa fa-chevron-right"></a></i></button>`, s.PageNumber+1)
	}
	page += "\n</div>"
	return template.HTML(page)
}

func (a *Web) NewSite(r *http.Request) (site Site) {
	site = Site{
		Name:        "Xegony",
		Title:       "Xegony",
		Description: "Xegony",
		PageSize:    getIntParam(r, "pageSize"),
		PageNumber:  getIntParam(r, "pageNumber"),
	}

	claims, err := api.GetAuthClaims(r)
	if err != nil && err.Error() != "No token provided" {
		//flush cookie
		log.Println("Bad auth", err.Error())
	}

	if claims != nil {
		site.User = claims.User
	}
	return
}

func (a *Web) Initialize(s storage.Storage, config string) (err error) {
	a.templates = map[string]*Template{}

	if s == nil {
		err = fmt.Errorf("Invalid storage type passed, must be pointer reference")
		return
	}

	a.accountRepo = &cases.AccountRepository{}
	a.accountRepo.Initialize(s)
	a.activityRepo = &cases.ActivityRepository{}
	a.activityRepo.Initialize(s)
	a.bazaarRepo = &cases.BazaarRepository{}
	a.bazaarRepo.Initialize(s)
	a.characterRepo = &cases.CharacterRepository{}
	a.characterRepo.Initialize(s)
	a.factionRepo = &cases.FactionRepository{}
	a.factionRepo.Initialize(s)
	a.forumRepo = &cases.ForumRepository{}
	a.forumRepo.Initialize(s)
	a.itemRepo = &cases.ItemRepository{}
	a.itemRepo.Initialize(s)
	a.lootDropRepo = &cases.LootDropRepository{}
	a.lootDropRepo.Initialize(s)
	a.lootDropEntryRepo = &cases.LootDropEntryRepository{}
	a.lootDropEntryRepo.Initialize(s)
	a.lootTableRepo = &cases.LootTableRepository{}
	a.lootTableRepo.Initialize(s)
	a.lootTableEntryRepo = &cases.LootTableEntryRepository{}
	a.lootTableEntryRepo.Initialize(s)
	a.npcRepo = &cases.NpcRepository{}
	a.npcRepo.Initialize(s)
	a.npcLootRepo = &cases.NpcLootRepository{}
	a.npcLootRepo.Initialize(s)
	a.postRepo = &cases.PostRepository{}
	a.postRepo.Initialize(s)
	a.taskRepo = &cases.TaskRepository{}
	a.taskRepo.Initialize(s)
	a.topicRepo = &cases.TopicRepository{}
	a.topicRepo.Initialize(s)
	a.userRepo = &cases.UserRepository{}
	a.userRepo.Initialize(s)
	a.zoneRepo = &cases.ZoneRepository{}
	a.zoneRepo.Initialize(s)
	return
}

func (a *Web) Index(w http.ResponseWriter, r *http.Request) {
	var err error

	type Content struct {
		Site Site
		Host string
	}

	site := a.NewSite(r)
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

	site := a.NewSite(r)
	site.Page = fmt.Sprintf("%d", statusCode)
	site.Title = "Error"

	tmp := a.getTemplate("")

	var tErr error
	switch statusCode {
	case http.StatusUnauthorized:
		if tmp == nil {
			tmp, tErr = a.loadTemplate(nil, "401", "401.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "401 - Unauthorized"
	case http.StatusBadRequest:
		if tmp == nil {
			tmp, tErr = a.loadTemplate(nil, "400", "400.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "400 - Bad Request"
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

func getIntParam(r *http.Request, key string) int64 {
	var val int64
	vals := r.URL.Query()
	keyTypes, ok := vals[key]
	if ok {
		if len(keyTypes) > 0 {
			val, _ = strconv.ParseInt(keyTypes[0], 10, 64)
			return val
		}
	}
	return 0
}

func getVar(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}

func getParam(r *http.Request, key string) string {
	val := getVar(r, key)
	if val != "" {
		return val
	}
	vals := r.URL.Query()
	keyTypes, ok := vals[key]
	if ok {
		if len(keyTypes) > 0 {
			return keyTypes[0]
		}
	}
	return ""
}
