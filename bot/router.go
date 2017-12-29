package bot

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func (a *Bot) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("BOT_ROOT")
	if len(rootPath) == 0 {
		rootPath = "/bot"
	}

	routes := Routes{
		Route{
			"Index",
			"GET",
			"/",
			a.Index,
		},
		Route{
			"Index",
			"GET",
			"/npcloot",
			a.NpcLootStatus,
		},
	}

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(rootPath + route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return
}
