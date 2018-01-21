package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
)

const (
	//JSON is a constant string representing json
	JSON = "json"
	//XML is a constant string representing xml
	XML = "xml"
	//YAML is a constant string representing yaml
	YAML = "yaml"
)

var (
	mySigningKey = []byte("øˆ∂∆ø∆12")
)

type loginResponse struct {
	APIKey string
	User   *model.User
}

//API wraps the api server
type API struct {
	log    *log.Logger
	logErr *log.Logger

	accountRepo        *cases.AccountRepository
	activityRepo       *cases.ActivityRepository
	bazaarRepo         *cases.BazaarRepository
	characterRepo      *cases.CharacterRepository
	factionRepo        *cases.FactionRepository
	forumRepo          *cases.ForumRepository
	goalRepo           *cases.GoalRepository
	itemRepo           *cases.ItemRepository
	lootDropEntryRepo  *cases.LootDropEntryRepository
	lootDropRepo       *cases.LootDropRepository
	lootTableEntryRepo *cases.LootTableEntryRepository
	lootTableRepo      *cases.LootTableRepository
	npcLootRepo        *cases.NpcLootRepository
	npcRepo            *cases.NpcRepository
	postRepo           *cases.PostRepository
	taskRepo           *cases.TaskRepository
	topicRepo          *cases.TopicRepository
	userRepo           *cases.UserRepository
	zoneLevelRepo      *cases.ZoneLevelRepository
	zoneRepo           *cases.ZoneRepository
}

// Initialize initializes an API endpoint with the implemented storage.
// config can be empty, it will initialize based on environment variables
// or by default values.
func (a *API) Initialize(s storage.Storage, config string, w io.Writer) (err error) {
	if s == nil {
		err = fmt.Errorf("Invalid storage type passed, must be pointer reference")
		return
	}
	if w == nil {
		w = os.Stdout
	}
	a.log = log.New(w, "API: ", 0)
	a.logErr = log.New(w, "APIErr: ", 0)

	a.accountRepo = &cases.AccountRepository{}
	if err = a.accountRepo.Initialize(s); err != nil {
		return
	}
	a.activityRepo = &cases.ActivityRepository{}
	if err = a.activityRepo.Initialize(s); err != nil {
		return
	}
	a.bazaarRepo = &cases.BazaarRepository{}
	if err = a.bazaarRepo.Initialize(s); err != nil {
		return
	}
	a.characterRepo = &cases.CharacterRepository{}
	if err = a.characterRepo.Initialize(s); err != nil {
		return
	}
	a.factionRepo = &cases.FactionRepository{}
	if err = a.factionRepo.Initialize(s); err != nil {
		return
	}
	a.forumRepo = &cases.ForumRepository{}
	if err = a.forumRepo.Initialize(s); err != nil {
		return
	}
	a.goalRepo = &cases.GoalRepository{}
	if err = a.goalRepo.Initialize(s); err != nil {
		return
	}
	a.itemRepo = &cases.ItemRepository{}
	if err = a.itemRepo.Initialize(s); err != nil {
		return
	}
	a.lootDropRepo = &cases.LootDropRepository{}
	if err = a.lootDropRepo.Initialize(s); err != nil {
		return
	}
	a.lootDropEntryRepo = &cases.LootDropEntryRepository{}
	if err = a.lootDropEntryRepo.Initialize(s); err != nil {
		return
	}
	a.lootTableRepo = &cases.LootTableRepository{}
	if err = a.lootTableRepo.Initialize(s); err != nil {
		return
	}
	a.lootTableEntryRepo = &cases.LootTableEntryRepository{}
	if err = a.lootTableEntryRepo.Initialize(s); err != nil {
		return
	}
	a.npcRepo = &cases.NpcRepository{}
	if err = a.npcRepo.Initialize(s); err != nil {
		return
	}
	a.npcLootRepo = &cases.NpcLootRepository{}
	if err = a.npcLootRepo.Initialize(s); err != nil {
		return
	}
	a.postRepo = &cases.PostRepository{}
	if err = a.postRepo.Initialize(s); err != nil {
		return
	}

	a.taskRepo = &cases.TaskRepository{}
	if err = a.taskRepo.Initialize(s); err != nil {
		return
	}
	a.topicRepo = &cases.TopicRepository{}
	if err = a.topicRepo.Initialize(s); err != nil {
		return
	}
	a.userRepo = &cases.UserRepository{}
	if err = a.userRepo.Initialize(s); err != nil {
		return
	}
	a.zoneRepo = &cases.ZoneRepository{}
	if err = a.zoneRepo.Initialize(s); err != nil {
		return
	}
	a.zoneLevelRepo = &cases.ZoneLevelRepository{}
	if err = a.zoneLevelRepo.Initialize(s); err != nil {
		return
	}
	a.log.Println("Initialized")
	return
}

func (a *API) indexRoutes() (routes []*route) {
	routes = []*route{
		{
			"Index",
			"GET",
			"/",
			a.index,
		},
	}
	return
}

// Index handles the root endpoint of /api/
func (a *API) index(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	type Content struct {
		Message string `json:"message"`
	}

	content = Content{
		Message: "Please refer to documentation for more details",
	}

	return
}

func (a *API) writeData(w http.ResponseWriter, r *http.Request, content interface{}, statusCode int) {
	var err error
	if w == nil || r == nil {
		a.logErr.Println("a.writeData called with invalid writer/request")
	}
	if content == nil {
		w.WriteHeader(statusCode)
		return
	}
	vals := r.URL.Query()
	format := JSON
	if formats, ok := vals["format"]; ok {
		if len(formats) >= 1 {
			format = strings.ToLower(formats[0]) // The first `?format=val`
		}
	}

	var data []byte
	switch strings.ToLower(format) {
	case XML:
		w.Header().Set("Content-Type", "application/xml; charset=UTF-8")
		if data, err = xml.Marshal(content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	case YAML:
		w.Header().Set("Content-Type", "application/yaml; charset=UTF-8")
		if data, err = yaml.Marshal(content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default: //json
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if data, err = json.Marshal(content); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}

func (a *API) writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	type Content struct {
		Message string            `json:"message"`
		Fields  map[string]string `json:"fields,omitempty"`
	}

	//if debug mode
	/*if err, ok := err.(model.StackTracer); ok {
		for _, f := range err.StackTrace() {
			fmt.Printf("%+s:%d", f)
		}
	}*/
	//fmt.Printf("%v\n", err)

	a.logErr.Println(err.Error())

	content := Content{
		Message: fmt.Sprintf("%s", err), //errors.Cause(err).Error(),
	}

	switch tErr := errors.Cause(err).(type) {
	case *model.ErrNoContent:
		return
	case *model.ErrValidation:
		content.Fields = map[string]string{}
		content.Message = "Invalid fields"
		for f, d := range tErr.Reasons {
			if content.Message == "Invalid fields" {
				content.Message = d
			}
			content.Fields[f] = d
		}
		statusCode = http.StatusBadRequest
	case *model.ErrPermission:
		statusCode = http.StatusUnauthorized
	}

	return
}

// decodeBody is used to convert raw json body content into a specified struct
func decodeBody(r *http.Request, data interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(data)
	if err != nil {
		err = &model.ErrDecodeBody{}
		return
	}
	return
}

// getIntParam parses query parameters based on key and returns as an int64
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

// getIntVar parses a variable from the routing pattern and returns it as an int64
func getIntVar(r *http.Request, key string) (val int64, err error) {
	vars := mux.Vars(r)
	val, err = strconv.ParseInt(vars[key], 10, 64)
	if err != nil {
		err = &model.ErrInvalidArguments{}
		return
	}
	return
}

// getVar  returns with a variable inside the request based on a routing pattern assigned variable
func getVar(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
