// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/bot"
	"github.com/xackery/xegony/storage"
	"github.com/xackery/xegony/storage/mariadb"
	"github.com/xackery/xegony/web"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Long:  `Start the server based on provided parameters. By default, it will run API, BOT, and WEB servers.`,
	Run:   runServer,
}

var connection string
var dbtype string

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringVarP(&connection, "connection", "", "eqemuconfig", "connection settings to connect to database, leave default for eqemu_connection")
	serverCmd.PersistentFlags().StringVarP(&dbtype, "dbtype", "", "mysql", "type of database to connect to")
}

func runServer(cmd *cobra.Command, args []string) {
	var err error
	var stor storage.Storage
	if dbtype == "mysql" {
		stor = &mariadb.Storage{}
	}

	if connection == "eqemuconfig" {
		connection = ""
	}
	stor.Initialize(connection, nil)

	err = stor.VerifyTables()
	if err != nil {
		log.Fatal("Failed to verify tables: ", err.Error())
	}

	listen := os.Getenv("API_LISTEN")
	if len(listen) == 0 {
		listen = ":8080"
	}

	router := mux.NewRouter().StrictSlash(true)

	botServer := bot.Bot{}
	if err = botServer.Initialize(stor, connection, nil); err != nil {
		log.Fatal("Failed to initialize botServer:", err.Error())
	}
	botServer.ApplyRoutes(router)

	apiServer := api.API{}
	if err = apiServer.Initialize(stor, connection, nil); err != nil {
		log.Fatal("Failed to initialize apiServer:", err.Error())
	}
	apiServer.ApplyRoutes(router)

	webServer := web.Web{}
	if err = webServer.Initialize(stor, connection, nil); err != nil {
		log.Fatal("Failed to initialize webServer:", err.Error())
	}
	webServer.ApplyRoutes(router)

	router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		url, _ := route.URL()
		path, _ := route.GetPathRegexp()
		log.Println(path, url, route.GetName())
		return nil
	})
	//go runBot(botServer)
	log.Println("Listening on", listen)
	err = http.ListenAndServe(listen, router)
	log.Println(err)
}

func runBot(botServer bot.Bot) {
	//err := botServer.CreateZoneMapCache()
	//if err != nil {
	///log.Fatal("Failed botserver:", err.Error())
	//}
	/*err := botServer.CreateZoneLevelCache()
	if err != nil {
		log.Fatal("Failed botserver:", err.Error())
	}*/
	/*err := botServer.CreateNpcDropsCache()
	if err != nil {
		log.Fatal("Failed botserver:", err.Error())
	}*/
}
