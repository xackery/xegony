package api

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/model"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc func(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error)
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

		router.
			Methods(route.Method).
			Name(route.Name).
			Path(rootPath + route.Pattern).
			Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				start := time.Now()

				auth, err := GetAuthClaim(r)
				user := &model.User{}
				if err == nil {
					user = auth.User
				}

				statusCode := http.StatusOK
				content, err := route.HandlerFunc(w, r, auth, user, statusCode)
				if err != nil {
					a.writeError(w, r, err, statusCode)
				} else {
					a.writeData(w, r, content, statusCode)
				}
				a.log.Printf(
					"%s %s: %s in %s",
					r.Method,
					r.RequestURI,
					route.Name,
					time.Since(start),
				)
			}))
	}

	return
}
