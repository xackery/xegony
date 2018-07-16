package web

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error)
}

//ApplyRoutes applies routes to given mux router
func ApplyRoutes(router *mux.Router) {
	rootPath := cases.GetConfigValue("webSuffix")

	if rootPath[len(rootPath)-1:] == "/" {
		rootPath = rootPath[0 : len(rootPath)-1]
	}
	var routes []*route

	routes = append(routes, characterRoutes()...)
	routes = append(routes, forumRoutes()...)
	routes = append(routes, indexRoutes()...)
	routes = append(routes, itemRoutes()...)
	routes = append(routes, npcRoutes()...)
	routes = append(routes, spellRoutes()...)
	routes = append(routes, spawnRoutes()...)
	routes = append(routes, zoneRoutes()...)

	type Content struct {
		User    *model.User `json:"user"`
		content interface{}
	}
	for i := range routes {
		route := routes[i]

		router.HandleFunc(rootPath+route.Pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			auth, err := api.GetAuthClaim(r)
			user := &model.User{}
			if err == nil {
				user = auth.User
			}

			tmp := &template.Template{}
			statusCode := http.StatusOK
			content, tmp, err := route.HandlerFunc(w, r, user, statusCode)
			if err != nil {
				writeError(w, r, err, statusCode)
			} else {
				writeData(w, r, tmp, content, statusCode)
			}
			log.Printf(
				"%s %s -> %s %s",
				r.Method,
				r.RequestURI,
				route.Name,
				time.Since(start),
			)
		})).
			Methods(route.Method).
			Name(route.Name)

	}

	router.NotFoundHandler = http.HandlerFunc(notFound)
	return
}
