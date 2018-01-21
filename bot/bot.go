package bot

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

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

	//StateIdle is an idle bot, ready to run
	StateIdle = "Idle"
	//StateRunning means it is currently doing a job
	StateRunning = "Running"
	//StateError means it errored and is waiting to be addressed
	StateError = "Error"
)

//Bot wraps all routing endpoints
type Bot struct {
	log    *log.Logger
	logErr *log.Logger

	statusTracker      sync.Map
	accountRepo        *cases.AccountRepository
	activityRepo       *cases.ActivityRepository
	bazaarRepo         *cases.BazaarRepository
	characterRepo      *cases.CharacterRepository
	factionRepo        *cases.FactionRepository
	forumRepo          *cases.ForumRepository
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

//Status groups together bot status related entries
type Status struct {
	Name      string
	State     string
	StartTime time.Time
	RunTime   time.Duration
}

func (b *Status) getRuntime() time.Duration {
	if b.State == StateRunning {
		return time.Since(b.StartTime)
	}
	return b.RunTime
}

func (a *Bot) startBot(key string) (err error) {
	bot := &Status{
		Name:      key,
		State:     StateIdle,
		StartTime: time.Now(),
	}

	rawBot, loaded := a.statusTracker.LoadOrStore(key, bot)
	if !loaded { //did not exist, started
		return
	}
	bot, ok := rawBot.(*Status)
	if !ok {
		err = fmt.Errorf("Invalid Status %s", key)
		return
	}
	if bot.State != StateIdle {
		err = fmt.Errorf("Bot is already running")
		return
	}
	bot.State = StateRunning
	a.statusTracker.Store(key, bot)
	return
}

func (a *Bot) endBot(key string) (err error) {

	rawBot, loaded := a.statusTracker.Load(key)
	if !loaded { //did not exist, started
		return
	}
	bot, ok := rawBot.(*Status)
	if !ok {
		err = fmt.Errorf("Invalid Status %s", key)
		return
	}
	if bot.State == StateIdle {
		err = fmt.Errorf("Bot is already ended")
		return
	}
	bot.State = StateIdle
	bot.RunTime = time.Since(bot.StartTime)
	a.statusTracker.Store(key, bot)
	return
}

func (a *Bot) getStatus(key string) (bot *Status, err error) {
	bot = &Status{
		Name:      key,
		State:     StateIdle,
		StartTime: time.Now(),
	}

	rawBot, loaded := a.statusTracker.LoadOrStore(key, bot)
	if !loaded { //did not exist, started
		return
	}
	bot, ok := rawBot.(*Status)
	if !ok {
		err = fmt.Errorf("Invalid Status %s", key)
		return
	}
	return
}

// Initialize initializes an API endpoint with the implemented storage.
// config can be empty, it will initialize based on environment variables
// or by default values.
func (a *Bot) Initialize(s storage.Storage, config string, w io.Writer) (err error) {
	if s == nil {
		err = fmt.Errorf("Invalid storage type passed, must be pointer reference")
		return
	}

	if w == nil {
		w = os.Stdout
	}
	a.log = log.New(w, "BOT: ", 0)
	a.logErr = log.New(w, "BOTErr: ", 0)

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

// Index handles the root endpoint of /api/
func (a *Bot) index(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error) {
	type Content struct {
		Message string `json:"message"`
	}
	content = Content{
		Message: "Please refer to documentation for more details",
	}
	return

}

func (a *Bot) writeData(w http.ResponseWriter, r *http.Request, content interface{}, statusCode int) {

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
	return
}

func (a *Bot) writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	type Content struct {
		Message string            `json:"message"`
		Fields  map[string]string `json:"fields,omitempty"`
	}

	content := Content{
		Message: fmt.Sprintf("%s", err), //errors.Cause(err).Error(),
	}

	a.logErr.Println(err.Error())

	switch errors.Cause(err).(type) {
	case *model.ErrNoContent:
		return
	case *model.ErrPermission:
		statusCode = http.StatusUnauthorized
	}

	a.writeData(w, r, content, statusCode)
	return
}
