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
	"fmt"
	alog "log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/xackery/xegony/api"
	"github.com/xackery/xegony/cases"
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

	serverCmd.PersistentFlags().StringVarP(&dbtype, "dbtype", "", "mysql", "type of database to connect to")
}

func runServer(cmd *cobra.Command, args []string) {

	err := startServer(cmd, args)
	if err != nil {
		fmt.Println("fail during startserver:", err)
	}
}

func startServer(cmd *cobra.Command, args []string) (err error) {

	sw, sr, si, w, wErr, err := initializeSystem()
	if err != nil {
		return
	}
	router := mux.NewRouter().StrictSlash(true)
	if err = api.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize api:", err.Error())
	}
	api.ApplyRoutes(router)

	if err = web.Initialize(sr, sw, si, connection, w, wErr); err != nil {
		log.Fatal("Failed to initialize web:", err.Error())
	}
	web.ApplyRoutes(router)

	log.Println("Listening on", cases.GetConfigForHTTP())
	err = http.ListenAndServe(cases.GetConfigForHTTP(), router)
	return
}
