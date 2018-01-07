package bot

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
		HandlerFunc http.HandlerFunc
	}

	routes := []Route{
		{
			"Index",
			"GET",
			"/",
			a.index,
		},
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
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(rootPath + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return
}
