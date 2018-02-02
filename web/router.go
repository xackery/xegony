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

	var routes []*route
	var newRoutes []*route

	newRoutes = accountRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}
	/*
		newRoutes = activityRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = bazaarRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = characterRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = dashboardRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = fishingRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = forageRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = forumRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = hackerRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = indexRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = itemRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = loginRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = lootDropRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = lootDropEntryRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = lootTableRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = lootTableEntryRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = merchantRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = merchantEntryRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = npcRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = postRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = recipeRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = recipeEntryRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = ruleRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = spawnRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = spawnEntryRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = spawnNpcRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = spellRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = taskRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = topicRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = variableRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
		newRoutes = zoneRoutes()
		for _, r := range newRoutes {
			routes = append(routes, r)
		}
	*/
	for i := range routes {
		route := routes[i]
		//		log.Println("path:", route.Pattern)

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
