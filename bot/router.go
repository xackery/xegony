package bot

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/model"
)

//ApplyRoutes applies routes to given mux router
func (a *Bot) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("BOT_ROOT")
	if len(rootPath) == 0 {
		rootPath = "/bot"
	}

	type Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc func(w http.ResponseWriter, r *http.Request, auth *model.AuthClaim, user *model.User, statusCode int) (content interface{}, err error)
	}

	routes := []Route{
		{
			"Index",
			"GET",
			"/",
			a.index,
		},
		/*
			//NpcLoot
			{
				"NpcLootStatus",
				"GET",
				"/npcloot",
				a.npcLootStatus,
			},
			{
				"NpcLootCreate",
				"POST",
				"/npcloot",
				a.npcLootCreate,
			},

			//ZoneLevel
			{
				"ZoneLevelsStatus",
				"GET",
				"/zonelevels",
				a.zoneLevelsStatus,
			},
			{
				"ZoneLevelsCreate",
				"POST",
				"/zonelevels",
				a.zoneLevelsCreate,
			},

			//ZoneMap
			{
				"ZoneMapStatus",
				"GET",
				"/zonemap",
				a.zoneMapStatus,
			},
			{
				"ZoneMapCreate",
				"POST",
				"/zonemap",
				a.zoneMapCreate,
			},

			//CharacterGraph
			{
				"CharacterGraphStatus",
				"GET",
				"/charactergraph",
				a.characterGraphStatus,
			},
			{
				"CharacterGraphCreate",
				"POST",
				"/charactergraph",
				a.characterGraphCreate,
			},
		*/
	}

	for _, route := range routes {

		router.HandleFunc(rootPath+route.Pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			auth, err := api.GetAuthClaim(r)
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
	return
}
