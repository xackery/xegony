package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//ApplyRoutes applies routes to given mux router
func (a *API) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("API_ROOT")
	if len(rootPath) == 0 {
		rootPath = "/api"
	}

	var routes []*route
	var newRoutes []*route

	newRoutes = a.indexRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.accountRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
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
	newRoutes = a.factionRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.forumRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.goalRoutes()
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
	newRoutes = a.npcRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.postRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.spawnRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	newRoutes = a.spawnEntryRoutes()
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
	newRoutes = a.zoneRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = a.logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(rootPath + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return
}
