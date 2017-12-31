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
			a.index,
		},

		Route{
			"CreateCharacter",
			"POST",
			"/character",
			a.createCharacter,
		},

		Route{
			"DeleteCharacter",
			"DELETE",
			"/character/{characterId}",
			a.deleteCharacter,
		},

		Route{
			"EditCharacter",
			"PUT",
			"/character/{characterId}",
			a.editCharacter,
		},

		Route{
			"GetCharacter",
			"GET",
			"/character/{characterId}",
			a.getCharacter,
		},

		Route{
			"ListCharacter",
			"GET",
			"/character",
			a.listCharacter,
		},

		Route{
			"CreateAccount",
			"POST",
			"/account",
			a.createAccount,
		},

		Route{
			"DeleteAccount",
			"DELETE",
			"/account/{accountId}",
			a.deleteAccount,
		},

		Route{
			"EditAccount",
			"PUT",
			"/account/{accountId}",
			a.editAccount,
		},

		Route{
			"GetAccount",
			"GET",
			"/account/{accountId}",
			a.getAccount,
		},

		Route{
			"ListAccount",
			"GET",
			"/account",
			a.listAccount,
		},

		Route{
			"CreateForum",
			"POST",
			"/forum",
			a.createForum,
		},

		Route{
			"DeleteForum",
			"DELETE",
			"/forum/{forumID}",
			a.deleteForum,
		},

		Route{
			"EditForum",
			"PUT",
			"/forum/{forumID}",
			a.editForum,
		},

		Route{
			"GetForum",
			"GET",
			"/forum/{forumID}",
			a.getForum,
		},

		Route{
			"ListForum",
			"GET",
			"/forum",
			a.listForum,
		},

		Route{
			"ListFaction",
			"GET",
			"/faction",
			a.listFaction,
		},

		Route{
			"CreateFaction",
			"POST",
			"/faction",
			a.createFaction,
		},

		Route{
			"DeleteFaction",
			"DELETE",
			"/faction/{factionId}",
			a.deleteFaction,
		},

		Route{
			"EditFaction",
			"PUT",
			"/faction/{factionId}",
			a.editFaction,
		},

		Route{
			"GetFaction",
			"GET",
			"/faction/{factionId}",
			a.getFaction,
		},

		Route{
			"CreateTopic",
			"POST",
			"/topic",
			a.createTopic,
		},

		Route{
			"DeleteTopic",
			"DELETE",
			"/topic/{topicId}",
			a.deleteTopic,
		},

		Route{
			"EditTopic",
			"PUT",
			"/topic/{topicId}",
			a.editTopic,
		},

		Route{
			"GetTopic",
			"GET",
			"/topic/{topicId}",
			a.getTopic,
		},

		Route{
			"ListItem",
			"GET",
			"/item",
			a.listItem,
		},

		Route{
			"CreateItem",
			"POST",
			"/item",
			a.createItem,
		},

		Route{
			"DeleteItem",
			"DELETE",
			"/item/{itemId}",
			a.deleteItem,
		},

		Route{
			"EditItem",
			"PUT",
			"/item/{itemId}",
			a.editItem,
		},

		Route{
			"GetItem",
			"GET",
			"/item/{itemId}",
			a.getItem,
		},

		Route{
			"GetItemTooltip",
			"GET",
			"/item/{itemId}/tooltip",
			a.getItemTooltip,
		},

		Route{
			"ListItem",
			"GET",
			"/item",
			a.listItem,
		},

		Route{
			"PostLogin",
			"POST",
			"/login",
			a.postLogin,
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
