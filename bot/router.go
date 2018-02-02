package bot

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/cases"
	"github.com/xackery/xegony/model"
)

//ApplyRoutes applies routes to given mux router
func ApplyRoutes(router *mux.Router) {
	rootPath := cases.GetConfigValue("botSuffix")

	type Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc func(w http.ResponseWriter, r *http.Request, user *model.User, statusCode int) (content interface{}, err error)
	}

	routes := []Route{
		{
			"Index",
			"GET",
			"/",
			index,
		},
		/*
			//NpcLoot
			{
				"NpcLootStatus",
				"GET",
				"/npcloot",
				npcLootStatus,
			},
			{
				"NpcLootCreate",
				"POST",
				"/npcloot",
				npcLootCreate,
			},

			//ZoneLevel
			{
				"ZoneLevelsStatus",
				"GET",
				"/zonelevels",
				zoneLevelsStatus,
			},
			{
				"ZoneLevelsCreate",
				"POST",
				"/zonelevels",
				zoneLevelsCreate,
			},

			//ZoneMap
			{
				"ZoneMapStatus",
				"GET",
				"/zonemap",
				zoneMapStatus,
			},
			{
				"ZoneMapCreate",
				"POST",
				"/zonemap",
				zoneMapCreate,
			},
		*/
		/*//CharacterGraph
		{
			"GetCharacterGraph",
			"GET",
			"/charactergraph",
			characterGraphStatus,
		},
		{
			"CreateCharacterGraph",
			"POST",
			"/charactergraph",
			characterGraphCreate,
		},*/
	}

	for i := range routes {
		route := routes[i]
		router.
			Methods(route.Method).
			Path(rootPath + route.Pattern).
			Name(route.Name).
			Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				start := time.Now()

				auth, err := api.GetAuthClaim(r)
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
					"%s %s -> %s %s",
					r.Method,
					r.RequestURI,
					route.Name,
					time.Since(start),
				)
			}))

	}
	return
}
