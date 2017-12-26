package web

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/xackery/xegony/box"
)

type Site struct {
	Title       string //Title of site
	Page        string
	Section     string
	Description string //Description for oprop

}

type Web struct {
	templates map[string]*Template
}

func (a *Web) Initialize(config string) (err error) {
	a.templates = map[string]*Template{}
	return

}

func (a *Web) NewSite() (site Site) {
	site = Site{
		Title:       "lfg.link",
		Description: "Looking for group",
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
	site.Page = "index"
	site.Title = "lfg.link"

	content := Content{
		Site: site,
		Host: r.Host,
	}

	if strings.Contains(r.Host, "localhost") {
		content.Host = "localhost:8081"
	} else {
		content.Host = "api.lfg.link"
	}

	tmp := a.getTemplate("")
	if tmp == nil {
		tmp, err = a.loadTemplate(nil, "index", "index.tpl")
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
