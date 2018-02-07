package web

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	alog "log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/xackery/xegony/box"
	"github.com/xackery/xegony/model"
	"github.com/xackery/xegony/storage"
)

var (
	log    *alog.Logger
	logErr *alog.Logger
)

func indexRoutes() (routes []*route) {
	routes = []*route{
		{
			"Index",
			"GET",
			"/",
			getIndex,
		},
	}
	return
}

// Initialize initializes Web endpoint with the implemented storage.
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
	log = alog.New(w, "WEB: ", 0)
	logErr = alog.New(w, "WEBError: ", 0)

	log.Println("Initialized")
	return
}

func getIndex(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error) {
	return listForum(w, r, user, statusCode)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL, "404 Not Found")
	var err error
	path := "www/" + r.URL.Path[1:]

	if _, err = os.Stat(path); err == nil {
		http.ServeFile(w, r, path)
		return
	}

	var bData []byte
	if bData, err = box.ReadFile(path); err == nil {
		reader := bytes.NewReader(bData)
		http.ServeContent(w, r, path, time.Now(), reader)
		return
	}

	//All failed, this is a true 404
	err = fmt.Errorf("404 - Not Found: %s", r.URL)
	writeError(w, r, err, http.StatusNotFound)
	return
}

func writeError(w http.ResponseWriter, r *http.Request, err error, statusCode int) {

	type Content struct {
		Site    site
		Message string
		URL     string
	}
	site := newSite(r, nil)
	site.Page = fmt.Sprintf("%d", statusCode)
	site.Title = "Error"

	tmp, tErr := loadTemplate(nil, "404", "404.tpl")
	if tErr != nil {
		logErr.Println(tErr.Error())
	}

	switch statusCode {
	case http.StatusUnauthorized:
		if tmp == nil {
			tmp, tErr = loadTemplate(nil, "401", "401.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "401 - Unauthorized"
	case http.StatusBadRequest:
		if tmp == nil {
			tmp, tErr = loadTemplate(nil, "400", "400.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "400 - Bad Request"
	case http.StatusNotFound: //404
		if tmp == nil {
			tmp, tErr = loadTemplate(nil, "404", "404.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "404 - Not Found"
	default: //500
		statusCode = http.StatusInternalServerError
		if tmp == nil {
			tmp, tErr = loadTemplate(nil, "500", "500.tpl")
			if tErr != nil {
				err = errors.Wrap(err, tErr.Error())
			}
		}
		site.Title = "500 - Internal Server Error"
	}

	content := Content{
		Site:    site,
		Message: fmt.Sprintf("%s", errors.Cause(err).Error()),
	}
	logErr.Println(err.Error())
	if r != nil {
		content.URL = r.URL.String()
	}

	writeData(w, r, tmp, content, statusCode)
}

func writeData(w http.ResponseWriter, r *http.Request, tmp *template.Template, content interface{}, statusCode int) {
	var err error
	w.WriteHeader(statusCode)

	if tmp == nil {
		logErr.Println("Failed to load template", content)
		return
	}

	if err = tmp.Execute(w, content); err != nil {
		logErr.Println("Failed to execute template:", err.Error())
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
