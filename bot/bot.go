//Package bot provides /bot/ endpoints, managing backround processes like cache
package bot

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
)

const (
	//JSON Defines an encoding type
	JSON = "json"
	//XML Defines an encoding type
	XML = "xml"
	//YAML Defines an encoding type
	YAML = "yaml"
)

//Bot wraps all routing endpoints
type Bot struct {
	accountRepo        *cases.AccountRepository
	activityRepo       *cases.ActivityRepository
	bazaarRepo         *cases.BazaarRepository
	characterRepo      *cases.CharacterRepository
	factionRepo        *cases.FactionRepository
	forumRepo          *cases.ForumRepository
	itemRepo           *cases.ItemRepository
	goalRepo           *cases.GoalRepository
	lootDropRepo       *cases.LootDropRepository
	lootDropEntryRepo  *cases.LootDropEntryRepository
	lootTableRepo      *cases.LootTableRepository
	lootTableEntryRepo *cases.LootTableEntryRepository
	npcRepo            *cases.NpcRepository
	npcLootRepo        *cases.NpcLootRepository
	postRepo           *cases.PostRepository
	taskRepo           *cases.TaskRepository
	topicRepo          *cases.TopicRepository
	userRepo           *cases.UserRepository
	zoneRepo           *cases.ZoneRepository
}

// Initialize initializes an API endpoint with the implemented storage.
// config can be empty, it will initialize based on environment variables
// or by default values.
func (a *Bot) Initialize(s storage.Storage, config string) (err error) {
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

// Index handles the root endpoint of /api/
func (a *Bot) index(w http.ResponseWriter, r *http.Request) {
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

func writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	type Content struct {
		Message string            `json:"message"`
		Fields  map[string]string `json:"fields,omitempty"`
	}

	content := Content{
		Message: fmt.Sprintf("%s", err), //errors.Cause(err).Error(),
	}

	switch errors.Cause(err).(type) {
	case *model.ErrNoContent:
		writeData(w, r, nil, http.StatusNotModified)
		return
	case *model.ErrPermission:
		statusCode = http.StatusUnauthorized
	}

	writeData(w, r, content, statusCode)
	return
}
