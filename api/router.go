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
	HandlerFunc func(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error)
}

//ApplyRoutes applies routes to given mux router
func ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("API_ROOT")
	if len(rootPath) == 0 {
		rootPath = "/api"
	}

	var routes []*route
	var newRoutes []*route

	newRoutes = indexRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}

	routes = append(routes, accountRoutes()...)
	routes = append(routes, characterRoutes()...)
	routes = append(routes, zoneRoutes()...)

	for i := range routes {
		route := routes[i]
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
				content, err := route.HandlerFunc(w, r, user, statusCode)
				if err != nil {
					writeError(w, r, err, statusCode)
				} else {
					writeData(w, r, content, statusCode)
				}
				log.Printf(
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
