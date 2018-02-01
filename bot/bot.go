package bot

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	alog "log"
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

var (
	log    *alog.Logger
	logErr *alog.Logger

	statusTracker sync.Map
)

//Status groups together bot status related entries
type Status struct {
	Name      string
	State     string
	StartTime time.Time
	RunTime   time.Duration
}

// Initialize initializes an API endpoint with the implemented storage.
// config can be empty, it will initialize based on environment variables
// or by default values.
func Initialize(sr storage.Reader, sw storage.Writer, si storage.Initializer, config string, w io.Writer, wErr io.Writer) (err error) {
	if sr == nil {
		err = fmt.Errorf("Invalid reader type passed, must be pointer reference")
		return
	}
	if sw == nil {
		err = fmt.Errorf("Invalid writer type passed, must be pointer reference")
		return
	}
	if si == nil {
		err = fmt.Errorf("Invalid initializer type passed, must be pointer reference")
		return
	}
	if w == nil {
		w = os.Stdout
	}
	log = alog.New(w, "BOT: ", 0)
	logErr = alog.New(w, "BOTErr: ", 0)

	err = cases.InitializeAll(sr, sw, si)
	if err != nil {
		err = errors.Wrap(err, "failed to initialize all")
		return
	}
	log.Println("Initialized")
	return
}

func (b *Status) getRuntime() time.Duration {
	if b.State == StateRunning {
		return time.Since(b.StartTime)
	}
	return b.RunTime
}

func startBot(key string) (err error) {
	bot := &Status{
		Name:      key,
		State:     StateIdle,
		StartTime: time.Now(),
	}

	rawBot, loaded := statusTracker.LoadOrStore(key, bot)
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
	statusTracker.Store(key, bot)
	return
}

func endBot(key string) (err error) {

	rawBot, loaded := statusTracker.Load(key)
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
	statusTracker.Store(key, bot)
	return
}

func getStatus(key string) (bot *Status, err error) {
	bot = &Status{
		Name:      key,
		State:     StateIdle,
		StartTime: time.Now(),
	}

	rawBot, loaded := statusTracker.LoadOrStore(key, bot)
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

// Index handles the root endpoint of /api/
func index(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error) {
	type Content struct {
		Message string `json:"message"`
	}
	content = Content{
		Message: "Please refer to documentation for more details",
	}
	return

}

func writeData(w http.ResponseWriter, r *http.Request, content interface{}, statusCode int) {

	var err error
	if w == nil || r == nil {
		logErr.Println("writeData called with invalid writer/request")
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

func writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {
	type Content struct {
		Message string            `json:"message"`
		Fields  map[string]string `json:"fields,omitempty"`
	}

	content := Content{
		Message: fmt.Sprintf("%s", err), //errors.Cause(err).Error(),
	}

	logErr.Println(err.Error())

	switch errors.Cause(err).(type) {
	case *model.ErrNoContent:
		return
	case *model.ErrPermission:
		statusCode = http.StatusUnauthorized
	}

	writeData(w, r, content, statusCode)
	return
}
