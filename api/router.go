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
		rootPath = "/api"
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

		Route{
			"CreateAccount",
			"POST",
			"/account",
			a.CreateAccount,
		},

		Route{
			"DeleteAccount",
			"DELETE",
			"/account/{accountId}",
			a.DeleteAccount,
		},

		Route{
			"EditAccount",
			"PUT",
			"/account/{accountId}",
			a.EditAccount,
		},

		Route{
			"GetAccount",
			"GET",
			"/account/{accountId}",
			a.GetAccount,
		},

		Route{
			"ListAccount",
			"GET",
			"/account",
			a.ListAccount,
		},

		Route{
			"CreateForum",
			"POST",
			"/forum",
			a.CreateForum,
		},

		Route{
			"DeleteForum",
			"DELETE",
			"/forum/{forumId}",
			a.DeleteForum,
		},

		Route{
			"EditForum",
			"PUT",
			"/forum/{forumId}",
			a.EditForum,
		},

		Route{
			"GetForum",
			"GET",
			"/forum/{forumId}",
			a.GetForum,
		},

		Route{
			"ListForum",
			"GET",
			"/forum",
			a.ListForum,
		},

		Route{
			"ListFaction",
			"GET",
			"/faction",
			a.ListFaction,
		},

		Route{
			"CreateFaction",
			"POST",
			"/faction",
			a.CreateFaction,
		},

		Route{
			"DeleteFaction",
			"DELETE",
			"/faction/{factionId}",
			a.DeleteFaction,
		},

		Route{
			"EditFaction",
			"PUT",
			"/faction/{factionId}",
			a.EditFaction,
		},

		Route{
			"GetFaction",
			"GET",
			"/faction/{factionId}",
			a.GetFaction,
		},

		Route{
			"CreateTopic",
			"POST",
			"/topic",
			a.CreateTopic,
		},

		Route{
			"DeleteTopic",
			"DELETE",
			"/topic/{topicId}",
			a.DeleteTopic,
		},

		Route{
			"EditTopic",
			"PUT",
			"/topic/{topicId}",
			a.EditTopic,
		},

		Route{
			"GetTopic",
			"GET",
			"/topic/{topicId}",
			a.GetTopic,
		},

		Route{
			"ListItem",
			"GET",
			"/item",
			a.ListItem,
		},

		Route{
			"CreateItem",
			"POST",
			"/item",
			a.CreateItem,
		},

		Route{
			"DeleteItem",
			"DELETE",
			"/item/{itemId}",
			a.DeleteItem,
		},

		Route{
			"EditItem",
			"PUT",
			"/item/{itemId}",
			a.EditItem,
		},

		Route{
			"GetItem",
			"GET",
			"/item/{itemId}",
			a.GetItem,
		},

		Route{
			"GetItemTooltip",
			"GET",
			"/item/{itemId}/tooltip",
			a.GetItemTooltip,
		},

		Route{
			"ListItem",
			"GET",
			"/item",
			a.ListItem,
		},

		Route{
			"PostLogin",
			"POST",
			"/login",
			a.PostLogin,
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
