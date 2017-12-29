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
	JSON = "json"
	XML  = "xml"
	YAML = "yaml"
)

type Bot struct {
	npcRepo *cases.NpcRepository
}

func (a *Bot) Initialize(s storage.Storage, config string) (err error) {
	if s == nil {
		err = fmt.Errorf("Invalid storage type passed, must be pointer reference")
		return
	}
	a.npcRepo = &cases.NpcRepository{}
	a.npcRepo.Initialize(s)
	return
}

func (a *Bot) Index(w http.ResponseWriter, r *http.Request) {
	log.Println("index")
	type Content struct {
		Message string `json:"message"`
	}
	content := Content{
		Message: "Please refer to documentation for more details",
	}
	writeData(w, r, content, http.StatusOK)
}

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
