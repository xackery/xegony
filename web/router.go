package web

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

//ApplyRoutes applies routes to given mux router
func (a *Web) ApplyRoutes(router *mux.Router) {
	rootPath := os.Getenv("WEB_ROOT")
	if len(rootPath) == 0 {
		rootPath = ""
	}

	type Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}

	routes := []Route{
		//Index
		{
			"Index",
			"GET",
			"/",
			a.listForum,
		},
		//Dashboard
		{
			"GetDashboard",
			"GET",
			"/dashboard",
			a.getDashboard,
		},
		//Login
		{
			"Login",
			"GET",
			"/login",
			a.getLogin,
		},
		{
			"Logout",
			"GET",
			"/logout",
			a.getLogout,
		},
		//Character
		{
			"SearchCharacter",
			"GET",
			"/character/search",
			a.searchCharacter,
		},
		{
			"SearchCharacter",
			"GET",
			"/character/search/{search}",
			a.searchCharacter,
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
			"ListCharacter",
			"GET",
			"/character/{characterID}/inventory",
			a.listItemByCharacter,
		},
		{
			"ListCharacterByRanking",
			"GET",
			"/character/ranking",
			a.listCharacterByRanking,
		},
		{
			"ListCharacterByOnline",
			"GET",
			"/character/byonline",
			a.listCharacterByOnline,
		},
		{
			"ListCharacterByAccount",
			"GET",
			"/character/byaccount/{accountID}",
			a.listCharacterByAccount,
		},
		{
			"ListCharacterByRanking",
			"GET",
			"/ranking",
			a.listCharacterByRanking,
		},
		//Error
		{
			"SearchError",
			"GET",
			"/error/search/{search}",
			a.searchError,
		},
		{
			"GetError",
			"GET",
			"/error/{errorID}",
			a.getError,
		},
		{
			"ListError",
			"GET",
			"/error",
			a.listError,
		},
		//Fishing
		{
			"GetFishing",
			"GET",
			"/fishing/{fishingID}",
			a.getFishing,
		},
		{
			"ListFishing",
			"GET",
			"/fishing",
			a.listFishing,
		},
		//Forage
		{
			"GetForage",
			"GET",
			"/forage/{forageID}",
			a.getForage,
		},
		{
			"ListForage",
			"GET",
			"/forage",
			a.listForage,
		},
		//Spell
		{
			"SearchSpell",
			"GET",
			"/spell/search/{search}",
			a.searchSpell,
		},
		{
			"GetSpell",
			"GET",
			"/spell/{spellID}",
			a.getSpell,
		},
		{
			"ListSpell",
			"GET",
			"/spell",
			a.listSpell,
		},
		//Npc
		{
			"GetNpc",
			"GET",
			"/npc/{npcID}",
			a.getNpc,
		},
		{
			"ListNpc",
			"GET",
			"/npc",
			a.listNpc,
		},
		{
			"ListNpcByZone",
			"GET",
			"/npc/byzone",
			a.listNpcByZone,
		},
		{
			"GetNpcByZone",
			"GET",
			"/npc/byzone/{zoneID}",
			a.getNpcByZone,
		},
		{
			"ListNpcByFaction",
			"GET",
			"/npc/byfaction",
			a.listNpcByFaction,
		},
		{
			"GetNpcByFaction",
			"GET",
			"/npc/byfaction/{factionID}",
			a.getNpcByFaction,
		},
		//Topic
		{
			"ListTopic",
			"GET",
			"/forum/{forumID}",
			a.listTopic,
		},
		{
			"GetTopic",
			"GET",
			"/topic/{topicID}/details",
			a.getTopic,
		},
		//Activity
		{
			"ListActivity",
			"GET",
			"/task/{taskID}",
			a.listActivity,
		},
		{
			"GetActivity",
			"GET",
			"/task/{taskID}/activity/{activityID}",
			a.getActivity,
		},
		//LootTable
		{
			"GetLootTable",
			"GET",
			"/loottable/{lootTableID}",
			a.getLootTable,
		},
		{
			"ListLootTable",
			"GET",
			"/loottable",
			a.listLootTable,
		},
		//LootDropEntry
		{
			"GetLootDropEntry",
			"GET",
			"/lootdrop/{lootDropID}/{itemID}",
			a.getLootDropEntry,
		},
		{
			"ListLootDropEntry",
			"GET",
			"/lootdrop/{lootDropID}",
			a.listLootDropEntry,
		},
		//Merchant
		{
			"GetMerchant",
			"GET",
			"/merchant/{merchantID}/details",
			a.getMerchant,
		},
		{
			"ListMerchant",
			"GET",
			"/merchant",
			a.listMerchant,
		},
		//MerchantEntry
		{
			"GetMerchantEntry",
			"GET",
			"/merchant/{merchantID}/{itemID}",
			a.getMerchantEntry,
		},
		{
			"ListMerchantEntry",
			"GET",
			"/merchant/{merchantID}",
			a.listMerchantEntry,
		},

		//Post
		{
			"ListPost",
			"GET",
			"/topic/{topicID}",
			a.listPost,
		},
		{
			"GetPost",
			"GET",
			"/post/{postID}",
			a.getPost,
		},
		//Recipe
		{
			"ListRecipe",
			"GET",
			"/recipe",
			a.listRecipe,
		},
		{
			"GetRecipe",
			"GET",
			"/recipe/{recipeID}/details",
			a.getRecipe,
		},
		//RecipeEntry
		{
			"ListRecipeEntry",
			"GET",
			"/recipe/{recipeID}",
			a.listRecipeEntry,
		},
		{
			"GetRecipeEntry",
			"GET",
			"/spawn/{recipeID}/{recipeEntryID}",
			a.getRecipeEntry,
		},
		//Spawn
		{
			"ListSpawn",
			"GET",
			"/spawn",
			a.listSpawn,
		},
		{
			"GetSpawn",
			"GET",
			"/spawn/{spawnID}/details",
			a.getSpawn,
		},
		//SpawnEntry
		{
			"ListSpawnEntry",
			"GET",
			"/spawn/{spawnGroupID}",
			a.listSpawnEntry,
		},
		{
			"GetSpawnEntry",
			"GET",
			"/spawn/{spawnGroupID}/{npcID}",
			a.getSpawnEntry,
		},
		//Forum
		{
			"ListForum",
			"GET",
			"/forum",
			a.listForum,
		},
		{
			"GetForum",
			"GET",
			"/forum/{forumID}/details",
			a.getForum,
		},
		{
			"CreateForum",
			"GET",
			"/forum/create",
			a.createForum,
		},
		//Task
		{
			"GetTask",
			"GET",
			"/task/{taskID}/details",
			a.getTask,
		},
		{
			"ListTask",
			"GET",
			"/task",
			a.listTask,
		},
		//Zone
		{
			"GetZone",
			"GET",
			"/zone/{zoneID}",
			a.getZone,
		},
		{
			"ListZone",
			"GET",
			"/zone",
			a.listZone,
		},
		{
			"ListZoneByLevels",
			"GET",
			"/zone/bylevels",
			a.listZoneByLevels,
		},
		{
			"ListZoneByHotzone",
			"GET",
			"/zone/byhotzone",
			a.listZoneByHotzone,
		},
		//Item
		{
			"GetItem",
			"GET",
			"/item/{itemID}",
			a.getItem,
		},
		{
			"SearchItem",
			"GET",
			"/item/search",
			a.searchItem,
		},
		{
			"SearchItemByAccount",
			"GET",
			"/item/search/byaccount",
			a.searchItemByAccount,
		},
		{
			"ListItemBySlot",
			"GET",
			"/item/byslot",
			a.listItemBySlot,
		},
		{
			"ListItemByZone",
			"GET",
			"/item/byzone",
			a.listItemByZone,
		},
		{
			"GetItemByZone",
			"GET",
			"/item/byzone/{zoneID}",
			a.getItemByZone,
		},
		{
			"GetItemBySlot",
			"GET",
			"/item/byslot/{slotID}",
			a.getItemBySlot,
		},
		{
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
