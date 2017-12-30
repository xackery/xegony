//Package api provides /api/ endpoints
package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
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
	JSON = "json"
	XML  = "xml"
	YAML = "yaml"
)

var (
	mySigningKey = []byte("øˆ∂∆ø∆12")
)

type LoginResponse struct {
	ApiKey string
	User   *model.User
}

type Api struct {
	accountRepo        *cases.AccountRepository
	activityRepo       *cases.ActivityRepository
	bazaarRepo         *cases.BazaarRepository
	characterRepo      *cases.CharacterRepository
	factionRepo        *cases.FactionRepository
	forumRepo          *cases.ForumRepository
	goalRepo           *cases.GoalRepository
	itemRepo           *cases.ItemRepository
	lootDropRepo       *cases.LootDropRepository
	lootDropEntryRepo  *cases.LootDropEntryRepository
	lootTableRepo      *cases.LootTableRepository
	lootTableEntryRepo *cases.LootTableEntryRepository
	npcRepo            *cases.NpcRepository
	postRepo           *cases.PostRepository
	taskRepo           *cases.TaskRepository
	topicRepo          *cases.TopicRepository
	userRepo           *cases.UserRepository
	zoneRepo           *cases.ZoneRepository
}

// Initialize initializes an API endpoint with the implemented storage.
// config can be empty, it will initialize based on environment variables
// or by default values.
func (a *Api) Initialize(s storage.Storage, config string) (err error) {
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
	a.goalRepo = &cases.GoalRepository{}
	a.goalRepo.Initialize(s)
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

// Index handles the root endpoint of /api/
func (a *Api) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("index")
	type Content struct {
		Message string `json:"message"`
	}

	content := Content{
		Message: "Please refer to documentation for more details",
	}

	writeData(w, r, content, http.StatusOK)
}

// writeData is the final step of all http responses. All routes should end here.
func writeData(w http.ResponseWriter, r *http.Request, content interface{}, statusCode int) {
	var err error
	if w == nil || r == nil {
		log.Println("writeData called with invalid writer/request")
		return
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

// writeError gracefully handles errors occurred during the routing.
// Calling this will call writeData, so you can safely return once it is called.
func writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
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

	content := Content{
		Message: fmt.Sprintf("%s", err), //errors.Cause(err).Error(),
	}

	switch tErr := errors.Cause(err).(type) {
	case *model.ErrNoContent:
		writeData(w, r, nil, http.StatusNotModified)
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

	writeData(w, r, content, statusCode)
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
