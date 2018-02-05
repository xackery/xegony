// Package api Xegony API.
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /api
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Xackery<xackery@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//     - application/xml
//     - application/yaml
//
//     Security:
//     - apiKey:
//
//     SecurityDefinitions:
//     apiKey:
//          type: apiKey
//          name: Authorization
//          in: header
//
//     Extensions:
//     x-meta-value: value
//     x-meta-array:
//       - value1
//       - value2
//     x-meta-array-obj:
//       - name: obj
//         value: field
//
// swagger:meta
package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	alog "log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-yaml/yaml"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
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
	log          *alog.Logger
	logErr       *alog.Logger
	cookieStore  *sessions.CookieStore
)

type loginResponse struct {
	APIKey string
	User   *model.User
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
	log = alog.New(w, "API: ", 0)
	logErr = alog.New(wErr, "APIErr: ", 0)

	cookieStore = sessions.NewCookieStore([]byte("™£ˆø®™£ˆ∆®lewifjwofij"))

	log.Println("Initialized")
	return
}

func indexRoutes() (routes []*route) {
	routes = []*route{
		{
			"Index",
			"GET",
			"/",
			index,
		},
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

	accept := strings.ToLower(r.Header.Get("accept"))
	if len(accept) == 0 {
		accept = format
	} else {
		if strings.Index(accept, "application/") >= 0 {
			accept = accept[strings.Index(accept, "application/")+12:]
		}
	}
	var data []byte

	switch accept {
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
		Message: fmt.Sprintf("%s", errors.Cause(err).Error()),
	}

	switch tErr := errors.Cause(err).(type) {
	case *model.ErrNoContent:
		statusCode = http.StatusNoContent
	case *model.ErrDecodeBody:
		statusCode = http.StatusBadRequest
		content.Message = "failed to decode body: " + content.Message
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
		err = errors.Wrap(err, content.Message)
	case *model.ErrPermission:
		statusCode = http.StatusUnauthorized
	default:
		statusCode = http.StatusInternalServerError
	}

	logErr.Println(r.URL, statusCode, err.Error())

	writeData(w, r, content, statusCode)
	return
}

// decodeBody is used to convert raw json body content into a specified struct
func decodeBody(r *http.Request, data interface{}) (err error) {
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(data)
	if err != nil {
		err = &model.ErrDecodeBody{
			Message: err.Error(),
		}
		return
	}
	return
}

// getIntQuery parses query parameters based on key and returns as an int64
func getIntQuery(r *http.Request, key string) int64 {
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

// getQuery parses query parameters based on key and returns as a string
func getQuery(r *http.Request, key string) string {
	vals := r.URL.Query()
	keyTypes, ok := vals[key]
	if ok {
		if len(keyTypes) > 0 {
			return keyTypes[0]
		}
	}
	return ""
}

// getIntVar parses a variable from the routing pattern and returns it as an int64
func getIntVar(r *http.Request, key string) int64 {
	vars := mux.Vars(r)
	val, err := strconv.ParseInt(vars[key], 10, 64)
	if err != nil {
		return 0
	}
	return val
}

// getVar  returns with a variable inside the request based on a routing pattern assigned variable
func getVar(r *http.Request, key string) string {
	vars := mux.Vars(r)
	return vars[key]
}
