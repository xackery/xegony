package web

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/model"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, tmp *template.Template, err error)
}

//ApplyRoutes applies routes to given mux router
func (a *Web) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("WEB_ROOT")
	if len(rootPath) == 0 {
		rootPath = ""
	}

	var routes []*route
	var newRoutes []*route

	newRoutes = a.activityRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.bazaarRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.characterRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.dashboardRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.fishingRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.forageRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.forumRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.hackerRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.indexRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.itemRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.loginRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.lootDropRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.lootDropEntryRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.lootTableRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.lootTableEntryRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.merchantRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.merchantEntryRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.npcRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.postRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.recipeRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.recipeEntryRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.ruleRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.spellRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.taskRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.topicRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.variableRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.zoneRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}

	for i, _ := range routes {
		route := routes[i]
		a.log.Println("path:", route.Pattern)

		router.HandleFunc(rootPath+route.Pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			start := time.Now()

			auth, err := api.GetAuthClaim(r)
			user := &model.User{}
			if err == nil {
				user = auth.User
			}

			tmp := &template.Template{}
			statusCode := http.StatusOK
			content, tmp, err := route.HandlerFunc(w, r, auth, user, statusCode)
			if err != nil {
				a.writeError(w, r, err, statusCode)
			} else {
				a.writeData(w, r, tmp, content, statusCode)
			}
			a.log.Printf(
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

	router.NotFoundHandler = http.HandlerFunc(a.notFound)
	return
}
