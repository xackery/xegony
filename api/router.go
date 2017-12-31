package api

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//ApplyRoutes applies routes to given mux router
func (a *Api) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("API_ROOT")
	if len(rootPath) == 0 {
		rootPath = "/api"
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

		{
			"CreateCharacter",
			"POST",
			"/character",
			a.createCharacter,
		},

		{
			"DeleteCharacter",
			"DELETE",
			"/character/{characterID}",
			a.deleteCharacter,
		},

		{
			"EditCharacter",
			"PUT",
			"/character/{characterID}",
			a.editCharacter,
		},

		{
			"GetCharacter",
			"GET",
			"/character/{characterID}",
			a.getCharacter,
		},

		{
			"ListCharacter",
			"GET",
			"/character",
			a.listCharacter,
		},

		{
			"CreateAccount",
			"POST",
			"/account",
			a.createAccount,
		},

		{
			"DeleteAccount",
			"DELETE",
			"/account/{accountID}",
			a.deleteAccount,
		},

		{
			"EditAccount",
			"PUT",
			"/account/{accountID}",
			a.editAccount,
		},

		{
			"GetAccount",
			"GET",
			"/account/{accountID}",
			a.getAccount,
		},

		{
			"ListAccount",
			"GET",
			"/account",
			a.listAccount,
		},

		{
			"CreateForum",
			"POST",
			"/forum",
			a.createForum,
		},

		{
			"DeleteForum",
			"DELETE",
			"/forum/{forumID}",
			a.deleteForum,
		},

		{
			"EditForum",
			"PUT",
			"/forum/{forumID}",
			a.editForum,
		},

		{
			"GetForum",
			"GET",
			"/forum/{forumID}",
			a.getForum,
		},

		{
			"ListForum",
			"GET",
			"/forum",
			a.listForum,
		},

		{
			"ListFaction",
			"GET",
			"/faction",
			a.listFaction,
		},

		{
			"CreateFaction",
			"POST",
			"/faction",
			a.createFaction,
		},

		{
			"DeleteFaction",
			"DELETE",
			"/faction/{factionID}",
			a.deleteFaction,
		},

		{
			"EditFaction",
			"PUT",
			"/faction/{factionID}",
			a.editFaction,
		},

		{
			"GetFaction",
			"GET",
			"/faction/{factionID}",
			a.getFaction,
		},

		{
			"CreateTopic",
			"POST",
			"/topic",
			a.createTopic,
		},

		{
			"DeleteTopic",
			"DELETE",
			"/topic/{topicID}",
			a.deleteTopic,
		},

		{
			"EditTopic",
			"PUT",
			"/topic/{topicID}",
			a.editTopic,
		},

		{
			"GetTopic",
			"GET",
			"/topic/{topicID}",
			a.getTopic,
		},

		{
			"ListItem",
			"GET",
			"/item",
			a.listItem,
		},

		{
			"CreateItem",
			"POST",
			"/item",
			a.createItem,
		},

		{
			"DeleteItem",
			"DELETE",
			"/item/{itemID}",
			a.deleteItem,
		},

		{
			"EditItem",
			"PUT",
			"/item/{itemID}",
			a.editItem,
		},

		{
			"GetItem",
			"GET",
			"/item/{itemID}",
			a.getItem,
		},

		{
			"GetItemTooltip",
			"GET",
			"/item/{itemID}/tooltip",
			a.getItemTooltip,
		},

		{
			"ListItem",
			"GET",
			"/item",
			a.listItem,
		},

		{
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
