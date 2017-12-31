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
			"/character/{characterID}",
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
			"/character/{characterID}/inventory",
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
			"/character/byaccount/{accountID}",
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
			"/npc/{npcID}",
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
			"/npc/byzone/{zoneID}",
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
			"/npc/byfaction/{factionID}",
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
			"/topic/{topicID}/details",
			a.getTopic,
		},

		Route{
			"ListActivity",
			"GET",
			"/task/{taskID}",
			a.listActivity,
		},

		Route{
			"GetActivity",
			"GET",
			"/task/{taskID}/activity/{activityID}",
			a.getActivity,
		},

		Route{
			"GetLootTable",
			"GET",
			"/loottable/{lootTableID}",
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
			"/lootdrop/{lootDropID}/{itemID}",
			a.getLootDropEntry,
		},

		Route{
			"ListLootDropEntry",
			"GET",
			"/lootdrop/{lootDropID}",
			a.listLootDropEntry,
		},

		Route{
			"ListPost",
			"GET",
			"/topic/{topicID}",
			a.listPost,
		},

		Route{
			"GetPost",
			"GET",
			"/post/{postID}",
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
			"/task/{taskID}/details",
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
			"/zone/{zoneID}",
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
			"/item/{itemID}",
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
			"/item/byzone/{zoneID}",
			a.getItemByZone,
		},

		Route{
			"GetItemBySlot",
			"GET",
			"/item/byslot/{slotID}",
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
