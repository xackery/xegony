package web

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

func (a *Web) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("WEB_ROOT")

	routes := Routes{
		Route{
			"Index",
			"GET",
			"/",
			a.ListForum,
		},

		Route{
			"GetDashboard",
			"GET",
			"/dashboard",
			a.GetDashboard,
		},

		Route{
			"Login",
			"GET",
			"/login",
			a.GetLogin,
		},

		Route{
			"SearchCharacter",
			"GET",
			"/character/search",
			a.SearchCharacter,
		},

		Route{
			"SearchCharacter",
			"GET",
			"/character/search/{search}",
			a.SearchCharacter,
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
			"ListCharacter",
			"GET",
			"/character/{characterId}/inventory",
			a.ListItemByCharacter,
		},

		Route{
			"ListCharacterByRanking",
			"GET",
			"/character/ranking",
			a.ListCharacterByRanking,
		},

		Route{
			"ListCharacterByAccount",
			"GET",
			"/character/byaccount/{accountId}",
			a.ListCharacterByAccount,
		},

		Route{
			"ListCharacterByRanking",
			"GET",
			"/ranking",
			a.ListCharacterByRanking,
		},

		Route{
			"GetNpc",
			"GET",
			"/npc/{npcId}",
			a.GetNpc,
		},

		Route{
			"ListNpc",
			"GET",
			"/npc",
			a.ListNpc,
		},

		Route{
			"ListNpcByZone",
			"GET",
			"/npc/byzone",
			a.ListNpcByZone,
		},

		Route{
			"GetNpcByZone",
			"GET",
			"/npc/byzone/{zoneId}",
			a.GetNpcByZone,
		},

		Route{
			"ListNpcByFaction",
			"GET",
			"/npc/byfaction",
			a.ListNpcByFaction,
		},

		Route{
			"GetNpcByFaction",
			"GET",
			"/npc/byfaction/{factionId}",
			a.GetNpcByFaction,
		},

		Route{
			"ListTopic",
			"GET",
			"/forum/{forumId}",
			a.ListTopic,
		},

		Route{
			"GetTopic",
			"GET",
			"/topic/{topicId}/details",
			a.GetTopic,
		},

		Route{
			"ListPost",
			"GET",
			"/topic/{topicId}",
			a.ListPost,
		},

		Route{
			"GetPost",
			"GET",
			"/post/{postId}",
			a.GetPost,
		},

		Route{
			"ListForum",
			"GET",
			"/forum",
			a.ListForum,
		},

		Route{
			"GetForum",
			"GET",
			"/forum/{forumId}/details",
			a.GetForum,
		},

		Route{
			"CreateForum",
			"GET",
			"/forum/create",
			a.CreateForum,
		},

		Route{
			"GetZone",
			"GET",
			"/zone/{zoneId}",
			a.GetZone,
		},

		Route{
			"ListZone",
			"GET",
			"/zone",
			a.ListZone,
		},

		Route{
			"GetItem",
			"GET",
			"/item/{itemId}",
			a.GetItem,
		},

		Route{
			"ListItemBySlot",
			"GET",
			"/item/byslot",
			a.ListItemBySlot,
		},

		Route{
			"GetItemBySlot",
			"GET",
			"/item/byslot/{slotId}",
			a.GetItemBySlot,
		},

		Route{
			"ListItem",
			"GET",
			"/item",
			a.ListItem,
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
	router.NotFoundHandler = http.HandlerFunc(a.notFound)

	return
}
