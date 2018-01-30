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
	alog "log"
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
var log *alog.Logger
var logErr *alog.Logger

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.PersistentFlags().StringVarP(&connection, "connection", "", "eqemuconfig", "connection settings to connect to database, leave default for eqemu_connection")
	serverCmd.PersistentFlags().StringVarP(&dbtype, "dbtype", "", "mysql", "type of database to connect to")
}

func runServer(cmd *cobra.Command, args []string) {
	var err error
	var sw storage.Writer
	var sr storage.Reader
	var si storage.Initializer
	w := os.Stdout
	wErr := os.Stderr

	log = alog.New(w, "Main: ", 0)
	logErr = alog.New(wErr, "MainErr: ", 0)

	if connection == "eqemuconfig" {
		connection = ""
	}
	if dbtype == "mysql" {
		var db *mariadb.Storage
		db, err = mariadb.New(connection, nil, nil)
		if err != nil {
			log.Fatal("Failed to create mariadb:", err.Error())
		}
		sw = db
		sr = db
		si = db
	} else {
		log.Fatal("unsupported db type:", dbtype)
	}

	err = si.VerifyTables()
	if err != nil {
		log.Fatal("Failed to verify tables: ", err.Error())
	}

	listen := os.Getenv("API_LISTEN")
	if len(listen) == 0 {
		listen = ":8080"
	}

	router := mux.NewRouter().StrictSlash(true)

	if err = bot.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize botServer:", err.Error())
	}
	bot.ApplyRoutes(router)

	if err = api.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize apiServer:", err.Error())
	}
	api.ApplyRoutes(router)

	if err = web.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize webServer:", err.Error())
	}
	web.ApplyRoutes(router)

	log.Println("Listening on", listen)
	err = http.ListenAndServe(listen, router)
	log.Println(err)
}
