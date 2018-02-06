package api

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/cases"
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
	rootPath := cases.GetConfigValue("apiSuffix")

	var routes []*route
	var newRoutes []*route

	newRoutes = indexRoutes()
	for _, r := range newRoutes {
		routes = append(routes, r)
	}

	routes = append(routes, accountRoutes()...)
	routes = append(routes, characterRoutes()...)
	routes = append(routes, classRoutes()...)
	routes = append(routes, configRoutes()...)
	routes = append(routes, deityRoutes()...)
	routes = append(routes, itemRoutes()...)
	routes = append(routes, npcRoutes()...)
	routes = append(routes, oauthTypeRoutes()...)
	routes = append(routes, raceRoutes()...)
	routes = append(routes, ruleRoutes()...)
	routes = append(routes, ruleEntryRoutes()...)
	routes = append(routes, spawnRoutes()...)
	routes = append(routes, spawnEntryRoutes()...)
	routes = append(routes, spawnNpcRoutes()...)
	routes = append(routes, spellRoutes()...)
	routes = append(routes, spellAnimationRoutes()...)
	routes = append(routes, spellAnimationTypeRoutes()...)
	routes = append(routes, spellEffectFormulaRoutes()...)
	routes = append(routes, spellEffectTypeRoutes()...)
	routes = append(routes, spellTargetTypeRoutes()...)
	routes = append(routes, spellTravelTypeRoutes()...)
	routes = append(routes, userRoutes()...)
	routes = append(routes, userGoogleRoutes()...)
	routes = append(routes, variableRoutes()...)
	routes = append(routes, zoneRoutes()...)
	routes = append(routes, zoneExpansionRoutes()...)

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
