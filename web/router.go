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
			a.listForum,
		},

		Route{
			"GetDashboard",
			"GET",
			"/dashboard",
			a.getDashboard,
		},

		Route{
			"Login",
			"GET",
			"/login",
			a.getLogin,
		},

		Route{
			"Logout",
			"GET",
			"/logout",
			a.getLogout,
		},

		Route{
			"SearchCharacter",
			"GET",
			"/character/search",
			a.searchCharacter,
		},

		Route{
			"SearchCharacter",
			"GET",
			"/character/search/{search}",
			a.searchCharacter,
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
			"ListCharacter",
			"GET",
			"/character/{characterId}/inventory",
			a.listItemByCharacter,
		},

		Route{
			"ListCharacterByRanking",
			"GET",
			"/character/ranking",
			a.listCharacterByRanking,
		},

		Route{
			"ListCharacterByOnline",
			"GET",
			"/character/byonline",
			a.listCharacterByOnline,
		},

		Route{
			"ListCharacterByAccount",
			"GET",
			"/character/byaccount/{accountId}",
			a.listCharacterByAccount,
		},

		Route{
			"ListCharacterByRanking",
			"GET",
			"/ranking",
			a.listCharacterByRanking,
		},

		Route{
			"GetNpc",
			"GET",
			"/npc/{npcId}",
			a.getNpc,
		},

		Route{
			"ListNpc",
			"GET",
			"/npc",
			a.listNpc,
		},

		Route{
			"ListNpcByZone",
			"GET",
			"/npc/byzone",
			a.listNpcByZone,
		},

		Route{
			"GetNpcByZone",
			"GET",
			"/npc/byzone/{zoneId}",
			a.getNpcByZone,
		},

		Route{
			"ListNpcByFaction",
			"GET",
			"/npc/byfaction",
			a.listNpcByFaction,
		},

		Route{
			"GetNpcByFaction",
			"GET",
			"/npc/byfaction/{factionId}",
			a.getNpcByFaction,
		},

		Route{
			"ListTopic",
			"GET",
			"/forum/{forumID}",
			a.listTopic,
		},

		Route{
			"GetTopic",
			"GET",
			"/topic/{topicId}/details",
			a.getTopic,
		},

		Route{
			"ListActivity",
			"GET",
			"/task/{taskId}",
			a.listActivity,
		},

		Route{
			"GetActivity",
			"GET",
			"/task/{taskId}/activity/{activityId}",
			a.getActivity,
		},

		Route{
			"GetLootTable",
			"GET",
			"/loottable/{lootTableId}",
			a.getLootTable,
		},

		Route{
			"ListLootTable",
			"GET",
			"/loottable",
			a.listLootTable,
		},

		Route{
			"GetLootDropEntry",
			"GET",
			"/lootdrop/{lootDropId}/{itemId}",
			a.getLootDropEntry,
		},

		Route{
			"ListLootDropEntry",
			"GET",
			"/lootdrop/{lootDropId}",
			a.listLootDropEntry,
		},

		Route{
			"ListPost",
			"GET",
			"/topic/{topicId}",
			a.listPost,
		},

		Route{
			"GetPost",
			"GET",
			"/post/{postId}",
			a.getPost,
		},

		Route{
			"ListForum",
			"GET",
			"/forum",
			a.listForum,
		},

		Route{
			"GetForum",
			"GET",
			"/forum/{forumID}/details",
			a.getForum,
		},

		Route{
			"CreateForum",
			"GET",
			"/forum/create",
			a.createForum,
		},

		Route{
			"GetTask",
			"GET",
			"/task/{taskId}/details",
			a.getTask,
		},

		Route{
			"ListTask",
			"GET",
			"/task",
			a.listTask,
		},

		Route{
			"GetZone",
			"GET",
			"/zone/{zoneId}",
			a.getZone,
		},

		Route{
			"ListZone",
			"GET",
			"/zone",
			a.listZone,
		},

		Route{
			"ListZoneByHotzone",
			"GET",
			"/zone/byhotzone",
			a.listZoneByHotzone,
		},

		Route{
			"GetItem",
			"GET",
			"/item/{itemId}",
			a.getItem,
		},

		Route{
			"SearchItem",
			"GET",
			"/item/search",
			a.searchItem,
		},

		Route{
			"SearchItemByAccount",
			"GET",
			"/item/search/byaccount",
			a.searchItemByAccount,
		},

		Route{
			"ListItemBySlot",
			"GET",
			"/item/byslot",
			a.listItemBySlot,
		},

		Route{
			"ListItemByZone",
			"GET",
			"/item/byzone",
			a.listItemByZone,
		},

		Route{
			"GetItemByZone",
			"GET",
			"/item/byzone/{zoneId}",
			a.getItemByZone,
		},

		Route{
			"GetItemBySlot",
			"GET",
			"/item/byslot/{slotId}",
			a.getItemBySlot,
		},

		Route{
			"ListItem",
			"GET",
			"/item",
			a.listItem,
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
	router.NotFoundHandler = http.HandlerFunc(a.notFound)

	return
}
