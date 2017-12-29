package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/bot"
	"github.com/xackery/xegony/storage/mariadb"
	"github.com/xackery/xegony/web"
)

func main() {
	var err error

	config := ""
	stor := &mariadb.Storage{}
	stor.Initialize("")
	listen := os.Getenv("API_LISTEN")
	if len(listen) == 0 {
		listen = ":8080"
	}

	router := mux.NewRouter().StrictSlash(true)

	botServer := bot.Bot{}
	if err = botServer.Initialize(stor, config); err != nil {
		log.Fatal("Failed to initialize botServer:", err.Error())
	}
	botServer.ApplyRoutes(router)

	apiServer := api.Api{}
	if err = apiServer.Initialize(stor, config); err != nil {
		log.Fatal("Failed to initialize apiServer:", err.Error())
	}
	apiServer.ApplyRoutes(router)

	webServer := web.Web{}
	if err = webServer.Initialize(stor, config); err != nil {
		log.Fatal("Failed to initialize webServer:", err.Error())
	}
	webServer.ApplyRoutes(router)

	//go runBot(botServer)

	err = http.ListenAndServe(listen, router)
	log.Println(err)
}

func runBot(botServer bot.Bot) {
	err := botServer.CreateNpcDropsCache()
	if err != nil {
		log.Fatal("Failed botserver:", err.Error())
	}
}
