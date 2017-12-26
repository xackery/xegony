package api

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

func (a *Api) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("API_ROOT")
	if len(rootPath) == 0 {
		rootPath = "/api/"
	}

	routes := Routes{
		Route{
			"Index",
			"GET",
			"/",
			a.Index,
		},

		Route{
			"CreateCharacter",
			"POST",
			"/character",
			a.CreateCharacter,
		},

		Route{
			"DeleteCharacter",
			"DELETE",
			"/character/{characterId}",
			a.DeleteCharacter,
		},

		Route{
			"EditCharacter",
			"PUT",
			"/character/{characterId}",
			a.EditCharacter,
		},

		Route{
			"GetCharacter",
			"GET",
			"/character/{characterId}",
			a.GetCharacter,
		},

		Route{
			"ListCharacter",
			"GET",
			"/character",
			a.ListCharacter,
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
