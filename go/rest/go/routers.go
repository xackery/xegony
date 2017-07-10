package 

import (
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"//",
		Index,
	},

	Route{
		"AddCharacter",
		"POST",
		"//character",
		AddCharacter,
	},

	Route{
		"DeleteCharacter",
		"DELETE",
		"//character",
		DeleteCharacter,
	},

	Route{
		"FindCharactersByAccountId",
		"GET",
		"//character/findByAccountId",
		FindCharactersByAccountId,
	},

	Route{
		"FindCharactersByName",
		"GET",
		"//character/findByName",
		FindCharactersByName,
	},

	Route{
		"GetCharacter",
		"GET",
		"//character/{characterId}",
		GetCharacter,
	},

	Route{
		"UpdateCharacter",
		"PUT",
		"//character",
		UpdateCharacter,
	},

	Route{
		"GetItem",
		"GET",
		"//item/{itemId}",
		GetItem,
	},

	Route{
		"ListItems",
		"GET",
		"//items",
		ListItems,
	},

	Route{
		"AddNpc",
		"POST",
		"//npc",
		AddNpc,
	},

	Route{
		"FindNpcsByName",
		"GET",
		"//npc/findByName",
		FindNpcsByName,
	},

	Route{
		"FindNpcsByZoneId",
		"GET",
		"//npc/findByZoneId",
		FindNpcsByZoneId,
	},

	Route{
		"GetNpc",
		"GET",
		"//npc/{npcId}",
		GetNpc,
	},

	Route{
		"UpdateNpc",
		"PUT",
		"//npc",
		UpdateNpc,
	},

	Route{
		"GetQuest",
		"GET",
		"//quest/{questId}",
		GetQuest,
	},

	Route{
		"ListQuests",
		"GET",
		"//quests",
		ListQuests,
	},

	Route{
		"GetRecipe",
		"GET",
		"//recipe/{recipeId}",
		GetRecipe,
	},

	Route{
		"ListRecipes",
		"GET",
		"//recipes",
		ListRecipes,
	},

	Route{
		"GetZone",
		"GET",
		"//zone/{zoneId}",
		GetZone,
	},

	Route{
		"ListZones",
		"GET",
		"//zones",
		ListZones,
	},

}